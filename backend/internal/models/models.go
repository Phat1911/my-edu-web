package models

import "time"

type Video struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DriveURL    string    `json:"drive_url" db:"drive_url"`
	EmbedURL    string    `json:"embed_url" db:"embed_url"`
	Thumbnail   string    `json:"thumbnail" db:"thumbnail"`
	Category    string    `json:"category" db:"category"`
	Order       int       `json:"order" db:"order_num"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Audio struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DriveURL    string    `json:"drive_url" db:"drive_url"`
	EmbedURL    string    `json:"embed_url" db:"embed_url"`
	Category    string    `json:"category" db:"category"`
	Duration    string    `json:"duration" db:"duration"`
	Order       int       `json:"order" db:"order_num"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type QRCode struct {
	ID        int       `json:"id" db:"id"`
	Label     string    `json:"label" db:"label"`
	TargetURL string    `json:"target_url" db:"target_url"`
	Type      string    `json:"type" db:"type"`
	QRData    string    `json:"qr_data" db:"qr_data"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ChatSession struct {
	ID        int       `json:"id" db:"id"`
	SessionID string    `json:"session_id" db:"session_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ChatMessage struct {
	ID        int       `json:"id" db:"id"`
	SessionID string    `json:"session_id" db:"session_id"`
	Role      string    `json:"role" db:"role"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PsychScenario struct {
	ID       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category"`
	Trigger  string `json:"trigger" db:"trigger"`
	Response string `json:"response" db:"response"`
	Tips     string `json:"tips" db:"tips"`
}
