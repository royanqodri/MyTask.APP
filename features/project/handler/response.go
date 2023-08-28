package handler

import "time"

type ProjectResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	UserID        uint      `json:"user_id"`
	DetailProject string    `json:"detail_project"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}
