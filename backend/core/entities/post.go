package entities

type Post struct {
	Userimg  string `json:"user_img"`
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	UserPic  string `json:"user_pic"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"created_at"`
	Likes    int    `json:"likes"`
	Privacy  string `json:"privacy"`
	Image    string `json:"image"`
	Comments int    `json:"comments"`
}
type Comment struct {
	Id      int    `json:"id"`
	Post_id int    `json:"post_id"`
	User_id int    `json:"user_id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Image   string `json:"image"`
}
