package data

import (
	"time"
)

type Task struct {
	ID        int64     `json:"id"`
	ProjectID int64     `json:"project_id"`
	OwnerID   int64     `json:"owner_id"`
	AssignID  int64     `json:"assign_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Priority  string    `json:"priority"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int32     `json:"version"`
}
