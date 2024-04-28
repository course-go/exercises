package todos

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"autoUpdateTime:false" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:false" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Todo struct {
	Model
	Description string     `json:"description,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func (*Todo) TableName() (name string) {
	return "todos"
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	var defaultUUID uuid.UUID
	if t.ID == defaultUUID {
		t.ID = uuid.New()
	}

	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
	fmt.Println("Create:", *t)
	return nil
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	fmt.Println("Updated:", *t)
	return nil
}
