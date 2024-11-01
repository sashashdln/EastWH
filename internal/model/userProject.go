package model

import "gorm.io/gorm"

type UserProject struct {
	gorm.Model
	ProjectID uint `gorm:"column:project_id" json:"project_id"`
	UserID    uint `gorm:"column:user_id" json:"user_id"`
}

func (UserProject) TableName() string {
	return "user_projects"
}
