package main

import (
	"github.com/in-rich/lib-go/deploy"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/config"
	"github.com/in-rich/uservice-discussions/migrations"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"log"
)

func main() {
	log.Println("Starting server")
	db, closeDB, err := deploy.OpenDB(config.App.Postgres.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer closeDB()

	log.Println("Running migrations")
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	depCheck := func() map[string]bool {
		errDB := db.Ping()

		return map[string]bool{
			"CreateMessage":              errDB == nil,
			"DeleteMessage":              errDB == nil,
			"GetDiscussionReadStatus":    errDB == nil,
			"GetMessage":                 errDB == nil,
			"ListDiscussionMessages":     errDB == nil,
			"ListDiscussionsByTeam":      errDB == nil,
			"UpdateDiscussionReadStatus": errDB == nil,
			"":                           errDB == nil,
		}
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

	createMessageHandler := handlers.NewCreateMessageHandler(createMessageService)
	deleteMessageHandler := handlers.NewDeleteMessageHandler(deleteMessageService)
	getDiscussionReadStatusHandler := handlers.NewGetDiscussionReadStatusHandler(getDiscussionReadStatusService)
	getMessageHandler := handlers.NewGetMessageHandler(getMessageService)
	listDiscussionMessagesHandler := handlers.NewListDiscussionMessagesHandler(listDiscussionMessagesService)
	listDiscussionsByTeamHandler := handlers.NewListDiscussionsByTeamHandler(listDiscussionsByTeamService)
	updateDiscussionReadStatusHandler := handlers.NewUpdateDiscussionReadStatusHandler(updateDiscussionReadStatusService)

	log.Println("Starting to listen on port", config.App.Server.Port)
	listener, server, health := deploy.StartGRPCServer(config.App.Server.Port, depCheck)
	defer deploy.CloseGRPCServer(listener, server)
	go health()

	discussions_pb.RegisterCreateMessageServer(server, createMessageHandler)
	discussions_pb.RegisterDeleteMessageServer(server, deleteMessageHandler)
	discussions_pb.RegisterGetDiscussionReadStatusServer(server, getDiscussionReadStatusHandler)
	discussions_pb.RegisterGetMessageServer(server, getMessageHandler)
	discussions_pb.RegisterListDiscussionMessagesServer(server, listDiscussionMessagesHandler)
	discussions_pb.RegisterListDiscussionsByTeamServer(server, listDiscussionsByTeamHandler)
	discussions_pb.RegisterUpdateDiscussionReadStatusServer(server, updateDiscussionReadStatusHandler)

	log.Println("Server started")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
