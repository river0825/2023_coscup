package command

import (
	"context"
	"errors"
	"fmt"
	"github.io/river0825/2023_coscup/module/backpack/domain/repository"
)

type TakeOutUsecase struct {
	repo repository.IBackpackRepo
}

type TakeOutCommand struct {
	BackpackId string `json:"backpack_id"`
	ItemId     string `json:"item_id"`
	Count      int    `json:"count"`
}
type TakeOutResponse struct {
	ItemId   string `json:"item_id"`
	Count    int    `json:"count"`
	ItemName string `json:"item_name"`
}

func validateTakeOut(command *TakeOutCommand) error {
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

func NewTakeOutUsecase(repo repository.IBackpackRepo) *TakeOutUsecase {
	return &TakeOutUsecase{
		repo: repo,
	}
}

func (u *TakeOutUsecase) Handle(ctx context.Context, command *TakeOutCommand) (*TakeOutResponse, error) {
	err := validateTakeOut(command)
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

	slot, err := p.TakeOut(command.ItemId, command.Count)
	if err != nil {
		return nil, err
	}

	err = u.repo.Save(p)
	return &TakeOutResponse{
		ItemId:   slot.ItemId,
		ItemName: slot.ItemName,
		Count:    slot.Count,
	}, err
}
