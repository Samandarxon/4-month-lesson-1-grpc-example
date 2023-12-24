package main

import (
	"dict/dict_client/handler"
	"dict/dictpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DictinoaryPort = "localhost:8090"
	HttpPort       = ":8080"
)

func main() {

	r := gin.Default()

	conn, err := grpc.Dial(DictinoaryPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc dial connection error:" + err.Error())
	}
	defer conn.Close()

	stub := dictpb.NewDictionaryServiceClient(conn)
	handler := handler.NewHandler(stub)

	r.POST("/translate", handler.Translate)
	r.POST("/add-message", handler.AddMessage)
	r.GET("/get-message/:key", handler.GetMessage)

	if err := r.Run(HttpPort); err != nil {
		panic("Gin run error:" + err.Error())
	}
}
