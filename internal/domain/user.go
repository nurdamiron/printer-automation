package domain

import (
    "time"

    "github.com/google/uuid"
)

// User представляет сущность пользователя в системе
type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Username  string    `gorm:"type:varchar(255);not null"`
    CreatedAt time.Time `gorm:"not null;default:now()"`
}
