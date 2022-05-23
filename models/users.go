package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Status   string `json:"status" gorm:"default:active"`
	Role     string `json:"role" gorm:"default:user"`
}

type Users struct {
	gorm.Model
	No            int    `json:"no"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	GiroAccountNo string `json:"giro_account_no"`
	RoleID        int    `json:"role_id"`
	RoleName      string `json:"role_name"`
	Status        string `json:"status"`
	WorkUnit      string `sql:"type:JSON" json:"work_unit"`
	MenuFeature   string `sql:"type:JSON" json:"menu_feature"`
}

func (s Users) TableName() string {
	return "tb_users"
}

type UserRequest struct {
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}
