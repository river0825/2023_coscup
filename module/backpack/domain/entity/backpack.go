package entity

import (
	"errors"
	"fmt"
)

func NewBackpack(id string) *Backpack {
	return &Backpack{
		Id:    id,
		slots: make([]*Slot, 0),
	}

}

type Backpack struct {
	slots []*Slot
	Id    string
}

func (b *Backpack) PutItem(itemId string, qty int) (*Slot, error) {
	slot := b.GetSlotByItemId(itemId)

	if slot == nil {
		slot = &Slot{
			ItemId: itemId,
			Count:  qty,
		}
		b.slots = append(b.slots, slot)
	} else {
		if slot.Count+qty > 999 {
			return slot, errors.New("total count is over 999")
		} else {
			slot.Count += qty
		}
	}

	return slot, nil
}

func (b *Backpack) GetSlotByItemId(id string) *Slot {
	for _, slot := range b.slots {
		if slot.ItemId == id {
			return slot
		}
	}

	return nil
}

func (b *Backpack) TakeOut(id string, qty int) (*Slot, error) {
	slot := b.GetSlotByItemId(id)
	if slot == nil {
		return nil, errors.New("item not found")
	}
	if slot.Count-qty < 0 {
		return slot, errors.New(fmt.Sprintf("item is not enough, total %d, requsted %d", slot.Count, qty))
	}

	slot.Count -= qty
	return slot, nil
}

func (b *Backpack) GetItem(id string) *Slot {
	return b.GetSlotByItemId(id)
}

func (b *Backpack) GetItems() []*Slot {
	return b.slots
}
