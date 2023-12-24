package handler

import "dict/dictpb"

type Handler struct {
	Stub dictpb.DictionaryServiceClient
}

func NewHandler(stub dictpb.DictionaryServiceClient) *Handler {
	return &Handler{
		Stub: stub,
	}
}
