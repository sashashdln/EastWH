package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name  string `gorm:"column:name" json:"name" validate:"required"`
	Users []User `gorm:"many2many:user_projects;" json:"users"`
}

func (Project) TableName() string {
	return "projects"
}
