package repository

import (
	"github.io/river0825/2023_coscup/module/backpack/application/query"
	"github.io/river0825/2023_coscup/module/backpack/domain/entity"
	domainError "github.io/river0825/2023_coscup/module/backpack/domain/errors"
	"github.io/river0825/2023_coscup/module/backpack/domain/repository"
)

type InMemRepository struct {
	storage map[string]*entity.Backpack
}

var _ repository.IBackpackRepo = (*InMemRepository)(nil)
var _ query.IBackpackQuery = (*InMemRepository)(nil)

func NewInMemRepository(defaultBackpackId string) *InMemRepository {
	repo := &InMemRepository{
		storage: make(map[string]*entity.Backpack),
	}

	repo.storage[defaultBackpackId] = entity.NewBackpack(defaultBackpackId)

	return repo
}
func (i *InMemRepository) Save(p *entity.Backpack) error {
	i.storage[p.Id] = p
	return nil
}

func (i *InMemRepository) Get(id string) (*entity.Backpack, error) {
	if _, ok := i.storage[id]; !ok {
		return nil, domainError.ErrBackpackNotFound
	}

	return i.storage[id], nil
}

/// GetItem Queries

func (i *InMemRepository) GetItems(backpackId string) ([]query.GetItemResponse, error) {
	b := i.storage[backpackId]
	if b == nil {
		return nil, domainError.ErrBackpackNotFound
	}

	// get item detail, we could query from database
	// itemInfo := db.Item.findOne({item_id: itemId})
	// itemName := itemInfo.Name
	itemName := "Master Sword"

	var items []query.GetItemResponse
	for _, s := range b.GetItems() {
		items = append(items, query.GetItemResponse{
			Id:    s.ItemId,
			Name:  itemName,
			Count: s.Count,
		})
	}

	return items, nil
}

func (i *InMemRepository) GetItem(backPackId string, itemId string) (*entity.Item, error) {
	b := i.storage[backPackId]
	if b == nil {
		return nil, domainError.ErrBackpackNotFound
	}

	s := b.GetSlotByItemId(itemId)

	// get item detail, we could query from database
	// itemInfo := db.Item.findOne({item_id: itemId})
	// itemName := itemInfo.Name
	itemName := "Master Sword"

	return &entity.Item{
		Id:    itemId,
		Name:  itemName,
		Count: s.Count,
	}, nil
}
