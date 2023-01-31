package base

import (
	"github.com/google/uuid"
	"time"
)

// EntityWithGuidKey is a base DB entity with uuid.UUID as a primary key.
type EntityWithGuidKey struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v1();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}

// EntityWithIntegerKey is a base DB entity with uint as a primary key.
type EntityWithIntegerKey struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}
