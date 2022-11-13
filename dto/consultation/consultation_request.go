package consultationdto

type ConsultationRequest struct {
	Name       string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Phone      string `gorm:"type: varchar(255)" json:"phone" validate:"required"`
	BornDate   string `json:"bornDate"`
	Age        string `json:"age" gorm:"type: varchar(255)" validate:"required"`
	Height     string `json:"height" gorm:"type: varchar(255)" validate:"required"`
	Weight     string `json:"weight" gorm:"type: varchar(255)" validate:"required"`
	Gender     string `json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Subject    string `json:"subject" gorm:"type: varchar(255)" validate:"required"`
	LiveConsul string `json:"liveConsul" `
	Desc       string `json:"desc" gorm:"type:text" form:"desc" validate:"required"`
	UserID     int    `json:"user_id" form:"user_id"`
	Status     string `json:"status" gorm:"type: varchar(255)"`
	Reply      string `json:"reply" gorm:"type:text" form:"reply" `
	LinkLive   string `json:"linkLive" gorm:"type: varchar(255)"`
}

type UpdateConsultationRequest struct {
	Name       string `gorm:"type: varchar(255)" json:"name"`
	Phone      string `gorm:"type: varchar(255)" json:"phone"`
	BornDate   string `json:"bornDate"`
	Age        string `json:"age" gorm:"type: string"`
	Height     string `json:"height" gorm:"type: string"`
	Weight     string `json:"weight" gorm:"type: string"`
	Gender     string `json:"gender" gorm:"type: varchar(255)"`
	Subject    string `json:"subject" gorm:"type: varchar(255)"`
	LiveConsul string `json:"liveConsul"`
	Desc       string `json:"desc" gorm:"type:text" form:"desc"`
	Status     string `json:"status" gorm:"type: varchar(255)"`
	Reply      string `json:"reply" gorm:"type:text" form:"reply" `
	LinkLive   string `json:"linkLive" gorm:"type: varchar(255)"`
}
