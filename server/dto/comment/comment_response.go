package commentdto

type CommentResponse struct {
	ID      int    `json:"id"`
	Comment string `json:"comment" gorm:"type: varchar(255)"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
