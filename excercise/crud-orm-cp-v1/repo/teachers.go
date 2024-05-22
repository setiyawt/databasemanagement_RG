package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	err := t.db.Create(&data)
	if err != nil {
		return nil
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	query, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listTeacher []model.Teacher
	for query.Next() {
		t.db.ScanRows(query, &listTeacher)
	}
	return listTeacher, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	teacher := model.Teacher{}
	if err := t.db.First(&teacher, id).Error; err != nil {
		return err
	}
	teacher.Name = name
	if err := t.db.Save(&teacher).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if err := t.db.Where("id = ?", id).Delete(&teacher); err.Error != nil {
		return nil // TODO: replace
	}
	return nil // TODO: replace this
}
