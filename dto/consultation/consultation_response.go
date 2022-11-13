package consultationdto

import (
	"backend/models"
	"time"
)

type ConsultationResponse struct {
	ID         int                        `json:"id"`
	Name       string                     `gorm:"type: varchar(255)" json:"name"`
	Phone      string                     `gorm:"type: varchar(255)" json:"phone"`
	BornDate   string                     `json:"bornDate"`
	Age        string                     `json:"age" gorm:"type: varchar(255)"`
	Height     string                     `json:"height" gorm:"type: varchar(255)"`
	Weight     string                     `json:"weight" gorm:"type: varchar(255)"`
	Gender     string                     `json:"gender" gorm:"type: varchar(255)"`
	Subject    string                     `json:"subject" gorm:"type: varchar(255)"`
	LiveConsul string                     `json:"liveConsul"`
	Desc       string                     `json:"desc" gorm:"type:text" form:"desc"`
	User       models.UserProfileResponse `json:"user" gorm:"foreignKey:UserID"`
	Status     string                     `json:"status" gorm:"type: varchar(255)"`
	Reply      string                     `json:"reply" gorm:"type:text" form:"reply" `
	LinkLive   string                     `json:"linkLive" gorm:"type: varchar(255)"`
	CreatedAt  time.Time                  `json:"createdAt"`
	UpdatedAt  time.Time                  `json:"updatedAt"`
}
