package model

type NotificationType string

const (
	Status    NotificationType = "status"
	News      NotificationType = "news"
	Marketing NotificationType = "marketing"
)

type Notification struct {
	Type    NotificationType `json:"type"`
	User    int              `json:"user"`
	Message string           `json:"message"`
}
