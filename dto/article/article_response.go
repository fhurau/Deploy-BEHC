package articledto

import (
	"backend/models"
	"time"
)

type ArticleResponse struct {
	ID        int                        `json:"id" gorm:"primary_key:auto_increment"`
	Title     string                     `json:"title" gorm:"type:varchar(255)"`
	Image     string                     `json:"image" gorm:"type: varchar(255)"`
	User      models.UserProfileResponse `json:"user" gorm:"foreignKey:UserID"`
	Desc      string                     `json:"desc" gorm:"type:text" form:"desc"`
	CreatedAt time.Time                  `json:"createdAt"`
	UpdatedAt time.Time                  `json:"updatedAt"`
}
