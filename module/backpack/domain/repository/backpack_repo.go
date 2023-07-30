package repository

import "github.io/river0825/2023_coscup/module/backpack/domain/entity"

type IBackpackRepo interface {
	Get(id string) (*entity.Backpack, error)
	Save(p *entity.Backpack) error
}
