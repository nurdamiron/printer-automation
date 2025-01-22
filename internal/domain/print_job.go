package domain

import (
    "time"

    "github.com/google/uuid"
)

// PrintJob представляет сущность задания печати
type PrintJob struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID    uuid.UUID `gorm:"type:uuid;not null"`
    Status    string    `gorm:"type:varchar(50);not null"`
    CreatedAt time.Time `gorm:"not null;default:now()"`
}
