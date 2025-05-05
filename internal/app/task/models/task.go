package models

import (
	"strings"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "Pending"
	StatusInProgress TaskStatus = "In Progress"
	StatusCompleted  TaskStatus = "Completed"
)

var ValidStatuses = []TaskStatus{
	StatusPending,
	StatusInProgress,
	StatusCompleted,
}

func IsValidStatus(status string) bool {
	for _, s := range ValidStatuses {
		if strings.EqualFold(string(s), status) {
			return true
		}
	}
	return false
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
