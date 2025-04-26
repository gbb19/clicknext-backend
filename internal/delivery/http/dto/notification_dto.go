package dto

import (
	"clicknext-backend/internal/domain"
	"time"
)

type NotificationResponse struct {
	NotifyID  uint      `json:"notify_id"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	Type      string    `json:"type"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotificationCreateRequest struct {
	Message string `json:"message" validate:"required"`
	Type    string `json:"type" validate:"required,oneof=task board"`
	UserID  uint   `json:"user_id" validate:"required"`
}

func (r *NotificationCreateRequest) ToNotificationDomain() *domain.Notification {
	return &domain.Notification{
		Message: r.Message,
		Type:    domain.NotifyType(r.Type),
		UserID:  r.UserID,
		IsRead:  false, // Default value for is_read is false
	}
}

func FromNotificationDomain(n *domain.Notification) *NotificationResponse {
	return &NotificationResponse{
		NotifyID:  n.NotifyID,
		Message:   n.Message,
		IsRead:    n.IsRead,
		Type:      string(n.Type),
		UserID:    n.UserID,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}
