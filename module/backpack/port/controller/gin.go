package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.io/river0825/2023_coscup/module/backpack/application/command"
	"github.io/river0825/2023_coscup/module/backpack/domain/repository"
	"net/http"
)

type GinController struct {
	putItemUsecase *command.PutItemUsecase
	takeOutUsecase *command.TakeOutUsecase
}

func NewGinController(repo repository.IBackpackRepo) *GinController {
	return &GinController{
		putItemUsecase: command.NewPutItemUsecase(repo),
		takeOutUsecase: command.NewTakeOutUsecase(repo),
	}
}

func (r *GinController) PutItem(ctx *gin.Context) {
	type PutItemRequest struct {
		BackpackId string `json:"backpack_id"`
		ItemId     string `json:"item_id"`
		Count      int    `json:"count"`
	}
	var request PutItemRequest
	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "request body must be json",
		})
	}

	cmd := &command.PutItemCommand{
		BackpackId: request.BackpackId,
		ItemId:     request.ItemId,
		Count:      request.Count,
	}

	rsp, err := r.putItemUsecase.Handle(ctx, cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    rsp,
	})
}

func (r *GinController) TakeOut(ctx *gin.Context) {
	type TakeOutRequest struct {
		BackpackId string `json:"backpack_id"`
		ItemId     string `json:"item_id"`
		Count      int    `json:"count"`
	}
	var request TakeOutRequest
	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	cmd := &command.TakeOutCommand{
		BackpackId: request.BackpackId,
		ItemId:     request.ItemId,
		Count:      request.Count,
	}

	rsp, err := r.takeOutUsecase.Handle(ctx, cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    rsp,
	})
}
