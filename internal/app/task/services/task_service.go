package services

import (
	models "task-management/internal/app/task/models"
	repo "task-management/internal/app/task/repository"
)

type Service struct {
	repo *repo.Repository
}

func NewTaskService(r *repo.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *Service) GetTask(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListTasks(status string, offset, limit int) ([]models.Task, error) {
	return s.repo.GetAll(status, offset, limit)
}

func (s *Service) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
