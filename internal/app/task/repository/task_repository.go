package task

import (
	models "task-management/internal/app/task/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *Repository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.DB.First(&task, id).Error
	return &task, err
}

func (r *Repository) GetAll(status string, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	query := r.DB
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(limit).Find(&tasks).Error
	return tasks, err
}

func (r *Repository) Update(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
