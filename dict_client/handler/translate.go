package handler

import (
	"context"
	"dict/dictpb"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Translate(c *gin.Context) {
	var req dictpb.TranslateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	resp, err := h.Stub.Tranlate(context.Background(), &req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, resp)
}

func (h *Handler) AddMessage(c *gin.Context) {
	var req dictpb.AddMessageRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	resp, err := h.Stub.AddMessage(context.Background(), &req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, resp)
}

func (h *Handler) GetMessage(c *gin.Context) {
	var req dictpb.GetMessageRequest
	req.Key = c.Param("key")

	fmt.Println(req.Key)

	resp, err := h.Stub.GetMessage(context.Background(), &req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, resp)
}
