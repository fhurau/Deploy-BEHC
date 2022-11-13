package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	ListAs   string `gorm:"type: varchar(100)" json:"listAs"`
	Gender   string `gorm:"type: varchar(100)" json:"gender"`
	Phone    string `gorm:"type: varchar(100)" json:"phone"`
	Address  string `gorm:"type: text" json:"address"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}
