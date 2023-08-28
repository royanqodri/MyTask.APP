package handler

import "time"

type TaskResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	ProjectID  uint      `json:"project_id"`
	Detailtask string    `json:"detail_task"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
