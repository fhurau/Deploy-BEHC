package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	CreateConsultation(Consultation models.Consultation) (models.Consultation, error)
	FindConsultation() ([]models.Consultation, error)
	GetConsultation(ID int) (models.Consultation, error)
	UpdateConsultation(Consultation models.Consultation, ID int) (models.Consultation, error)
	DeleteConsultation(Consultation models.Consultation, ID int) (models.Consultation, error)
}

func RepositoryConsultation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateConsultation(Consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Create(&Consultation).Error

	return Consultation, err
}

func (r *repository) FindConsultation() ([]models.Consultation, error) {
	var Consultation []models.Consultation
	err := r.db.Preload("User").Find(&Consultation).Error

	return Consultation, err
}

func (r *repository) GetConsultation(ID int) (models.Consultation, error) {
	var Consultation models.Consultation
	err := r.db.Preload("User").First(&Consultation, ID).Error

	return Consultation, err
}
func (r *repository) UpdateConsultation(Consultation models.Consultation, ID int) (models.Consultation, error) {
	err := r.db.Save(&Consultation).Error

	return Consultation, err
}
func (r *repository) DeleteConsultation(Consultation models.Consultation, ID int) (models.Consultation, error) {
	err := r.db.Delete(&Consultation).Error

	return Consultation, err
}
