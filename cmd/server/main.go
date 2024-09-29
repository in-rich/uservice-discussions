package main

import (
	"fmt"
	"github.com/in-rich/lib-go/deploy"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/config"
	"github.com/in-rich/uservice-discussions/migrations"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/rs/zerolog"
	"os"
)

func getLogger() monitor.GRPCLogger {
	if deploy.IsReleaseEnv() {
		return monitor.NewGCPGRPCLogger(zerolog.New(os.Stdout), "uservice-discussions")
	}

	return monitor.NewConsoleGRPCLogger()
}

func main() {
	logger := getLogger()

	logger.Info("Starting server")
	db, closeDB, err := deploy.OpenDB(config.App.Postgres.DSN)
	if err != nil {
		logger.Fatal(err, "failed to connect to database")
	}
	defer closeDB()

	logger.Info("Running migrations")
	if err := migrations.Migrate(db); err != nil {
		logger.Fatal(err, "failed to migrate")
	}

	depCheck := deploy.DepsCheck{
		Dependencies: func() map[string]error {
			return map[string]error{
				"Postgres": db.Ping(),
			}
		},
		Services: deploy.DepCheckServices{
			"CreateMessage":              {"Postgres"},
			"DeleteMessage":              {"Postgres"},
			"GetDiscussionReadStatus":    {"Postgres"},
			"GetMessage":                 {"Postgres"},
			"ListDiscussionMessages":     {"Postgres"},
			"ListDiscussionsByTeam":      {"Postgres"},
			"UpdateDiscussionReadStatus": {"Postgres"},
		},
	}

	createMessageDAO := dao.NewCreateMessageRepository(db)
	deleteMessageDAO := dao.NewDeleteMessageRepository(db)
	getDiscussionReadStatusDAO := dao.NewGetDiscussionReadStatusRepository(db)
	getMessageDAO := dao.NewGetMessageRepository(db)
	listDiscussionMessagesDAO := dao.NewListDiscussionMessagesRepository(db)
	listDiscussionsByTeamDAO := dao.NewListDiscussionsByTeamRepository(db)
	upsertDiscussionReadStatusDAO := dao.NewUpsertDiscussionReadStatusRepository(db)

	createMessageService := services.NewCreateMessageService(createMessageDAO)
	deleteMessageService := services.NewDeleteMessageService(deleteMessageDAO)
	getDiscussionReadStatusService := services.NewGetDiscussionReadStatusService(getDiscussionReadStatusDAO)
	getMessageService := services.NewGetMessageService(getMessageDAO)
	listDiscussionMessagesService := services.NewListDiscussionMessagesService(listDiscussionMessagesDAO)
	listDiscussionsByTeamService := services.NewListDiscussionsByTeamService(listDiscussionsByTeamDAO)
	updateDiscussionReadStatusService := services.NewUpdateDiscussionReadStatusService(upsertDiscussionReadStatusDAO, getMessageDAO)

	createMessageHandler := handlers.NewCreateMessageHandler(createMessageService, logger)
	deleteMessageHandler := handlers.NewDeleteMessageHandler(deleteMessageService, logger)
	getDiscussionReadStatusHandler := handlers.NewGetDiscussionReadStatusHandler(getDiscussionReadStatusService, logger)
	getMessageHandler := handlers.NewGetMessageHandler(getMessageService, logger)
	listDiscussionMessagesHandler := handlers.NewListDiscussionMessagesHandler(listDiscussionMessagesService, logger)
	listDiscussionsByTeamHandler := handlers.NewListDiscussionsByTeamHandler(listDiscussionsByTeamService, logger)
	updateDiscussionReadStatusHandler := handlers.NewUpdateDiscussionReadStatusHandler(updateDiscussionReadStatusService, logger)

	logger.Info(fmt.Sprintf("Starting to listen on port %v", config.App.Server.Port))
	listener, server, health := deploy.StartGRPCServer(logger, config.App.Server.Port, depCheck)
	defer deploy.CloseGRPCServer(listener, server)
	go health()

	discussions_pb.RegisterCreateMessageServer(server, createMessageHandler)
	discussions_pb.RegisterDeleteMessageServer(server, deleteMessageHandler)
	discussions_pb.RegisterGetDiscussionReadStatusServer(server, getDiscussionReadStatusHandler)
	discussions_pb.RegisterGetMessageServer(server, getMessageHandler)
	discussions_pb.RegisterListDiscussionMessagesServer(server, listDiscussionMessagesHandler)
	discussions_pb.RegisterListDiscussionsByTeamServer(server, listDiscussionsByTeamHandler)
	discussions_pb.RegisterUpdateDiscussionReadStatusServer(server, updateDiscussionReadStatusHandler)

	logger.Info("Server started")
	if err := server.Serve(listener); err != nil {
		logger.Fatal(err, "failed to serve")
	}
}
