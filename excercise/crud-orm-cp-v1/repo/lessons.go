package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type LessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) LessonRepo {
	return LessonRepo{db}
}

func (l LessonRepo) Init(data []model.Lesson) error {
	lesson := model.Lesson{Name: "Matematika"}
	err := l.db.Create(&lesson)
	if err != nil {
		return nil
	}
	return nil // TODO: replace this
}
