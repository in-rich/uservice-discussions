package main

import (
	"database/sql"
	"fmt"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/config"
	"github.com/in-rich/uservice-discussions/migrations"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.App.Postgres.DSN)))
	db := bun.NewDB(sqldb, pgdialect.New())

	defer func() {
		_ = db.Close()
		_ = sqldb.Close()
	}()

	err := db.Ping()
	for i := 0; i < 10 && err != nil; i++ {
		time.Sleep(1 * time.Second)
		err = db.Ping()
	}

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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	defer func() {
		server.GracefulStop()
		_ = listener.Close()
	}()

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
