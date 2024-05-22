package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepo struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) ClassRepo {
	return ClassRepo{db}
}

func (c ClassRepo) Init(data []model.Class) error {
	class := model.Class{Name: "IPA - 1"}
	err := c.db.Create(&class)
	if err != nil {
		return nil
	}
	return nil // TODO: replace this
}
