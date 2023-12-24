package main

import (
	"context"
	"dict/dictpb"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8090"
)

type DictionaryService struct {
	dictpb.UnimplementedDictionaryServiceServer
}

func (d *DictionaryService) Tranlate(ctx context.Context, req *dictpb.TranslateRequest) (*dictpb.TranslateResponse, error) {

	fmt.Printf("Req ~~~~~~~~~~~~>>>>>>%+v\n", req)

	var resp dictpb.TranslateResponse
	switch req.Input {
	case "Apple":
		resp.Output = "Olma"
	case "Orange":
		resp.Output = "Apelsin"
	default:
		return &dictpb.TranslateResponse{}, errors.New("Not found output")
	}

	return &resp, nil
}

var AddMessage = make(map[string]string, 0)

func (d *DictionaryService) AddMessage(ctx context.Context, req *dictpb.AddMessageRequest) (*dictpb.AddMessageResponse, error) {

	var resp dictpb.AddMessageResponse
	fmt.Printf("Req ~~~~~~~~~~~~>>>>>>%+v\n", req)

	if len(req.Key) <= 0 && len(req.Value) <= 0 {
		return nil, errors.New("Key Or Value are empty")

	}

	AddMessage[req.Key] = req.Value

	fmt.Println(AddMessage)
	resp.Message = "Added Key Value"

	return &resp, nil
}

func (d *DictionaryService) GetMessage(ctx context.Context, req *dictpb.GetMessageRequest) (*dictpb.GetMessageResponse, error) {

	var resp dictpb.GetMessageResponse
	fmt.Printf("Req ~~~~~~~~~~~~>>>>>>%+v\n", req)

	if len(req.Key) <= 0 {
		return nil, errors.New("Key Or Value are empty")

	}

	resp.Value = AddMessage[req.Key]

	fmt.Println(AddMessage, resp.Value)

	return &resp, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("listent erorr:" + err.Error())
	}

	grpcServer := grpc.NewServer()

	dictpb.RegisterDictionaryServiceServer(grpcServer, &DictionaryService{})

	fmt.Println("Listening....", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic("grpc server error:" + err.Error())
	}

}
