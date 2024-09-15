package main

import (
	"fmt"
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
	db, closeDB := deploy.OpenDB(config.App.Postgres.DSN)
	defer closeDB()

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
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

	listener, server := deploy.StartGRPCServer(fmt.Sprintf(":%d", config.App.Server.Port), "discussions")
	defer deploy.CloseGRPCServer(listener, server)

	discussions_pb.RegisterCreateMessageServer(server, createMessageHandler)
	discussions_pb.RegisterDeleteMessageServer(server, deleteMessageHandler)
	discussions_pb.RegisterGetDiscussionReadStatusServer(server, getDiscussionReadStatusHandler)
	discussions_pb.RegisterGetMessageServer(server, getMessageHandler)
	discussions_pb.RegisterListDiscussionMessagesServer(server, listDiscussionMessagesHandler)
	discussions_pb.RegisterListDiscussionsByTeamServer(server, listDiscussionsByTeamHandler)
	discussions_pb.RegisterUpdateDiscussionReadStatusServer(server, updateDiscussionReadStatusHandler)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
