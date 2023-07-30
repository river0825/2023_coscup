package command

import (
	"context"
	"errors"
	"fmt"
	"github.io/river0825/2023_coscup/module/backpack/domain/repository"
)

type PutItemUsecase struct {
	repo repository.IBackpackRepo
}

type PutItemCommand struct {
	BackpackId string `json:"backpack_id"`
	ItemId     string `json:"item_id"`
	Count      int    `json:"count"`
}

type PutItemResponse struct {
	BackpackId string `json:"backpack_id"`
	ItemId     string `json:"item_id"`
	Count      int    `json:"count"`
}

func validatePutItem(command *PutItemCommand) error {
	err := make([]error, 0)
	if command.BackpackId == "" {
		err = append(err, errors.New("backpack id is required"))
	}

	if command.ItemId == "" {
		err = append(err, errors.New("item id is required"))
	}

	if command.Count <= 0 {
		err = append(err, errors.New("count must be greater than 0"))
	}

	if len(err) != 0 {
		return errors.Join(err...)
	}

	return nil
}

func NewPutItemUsecase(repo repository.IBackpackRepo) *PutItemUsecase {
	return &PutItemUsecase{
		repo: repo,
	}
}

func (u *PutItemUsecase) Handle(ctx context.Context, command *PutItemCommand) (*PutItemResponse, error) {
	err := validatePutItem(command)
	if err != nil {
		return nil, err
	}

	p, err := u.repo.Get(command.BackpackId)
	if p == nil {
		return nil, errors.New(fmt.Sprintf("Backpack %s not found", command.BackpackId))
	}

	if err != nil {
		return nil, err
	}

	slot, err := p.PutItem(command.ItemId, command.Count)
	if err != nil {
		return nil, err
	}

	err = u.repo.Save(p)
	if err != nil {
		return nil, err
	}
	return &PutItemResponse{
		ItemId:     slot.ItemId,
		Count:      slot.Count,
		BackpackId: command.BackpackId,
	}, nil
}
