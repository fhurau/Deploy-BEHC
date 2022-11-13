package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	ListAs   string `gorm:"type: varchar(255)" json:"listAs"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	ListAs   string `gorm:"type: varchar(255)" json:"listAs"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	ListAs   string `gorm:"type: varchar(255)" json:"listAs"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}
