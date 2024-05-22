package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SchoolRepo struct {
	db *gorm.DB
}

func NewSchoolRepo(db *gorm.DB) SchoolRepo {
	return SchoolRepo{db}
}

func (s SchoolRepo) Init(data []model.School) error {
	school := model.School{Name: "SMAN 1 Jakarta", Phone: "(021) 3865001", Address: "Jl. Budi Utomo No.7, Ps. Baru, Kecamatan Sawah Besar, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10710", Province: "Jakarta"}
	err := s.db.Create(&school)
	if err != nil {
		return nil
	}
	return nil // TODO: replace this
}
