package models

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username" binding:"required, min=4,max=32"`
	Name            string `json:"name" binding:"required, min=1,max=70"`
	Bio             string `json:"bio" binding:"required, max=70"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required, min=8, max=16"`
	ConfirmPassword string `json:"cpassword" binding:"required"`
}

type Login struct {
	Credential string `json:"credential"`
	Password   string
}

type Follow struct {
	FollowID int `json:"follow-id"`
	FollowBy int `json:"follow-by"`
	FolloTo  int `json:"follow-to"`
}

type Post struct {
	PostID  int    `json:"post-id"`
	UserID  int    `json:"user-id"`
	Content string `json:"content"`
}
