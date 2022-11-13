package articledto

type ArticleRequest struct {
	Title  string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	UserID int    `json:"user_id" form:"user_id"`
	Image  string `json:"image" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text" form:"desc" validate:"required"`
}

type UpdateArticleRequest struct {
	Title string `json:"title" gorm:"type: varchar(255)"`
	Image string `json:"image" gorm:"type: varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text" form:"desc"`
}
