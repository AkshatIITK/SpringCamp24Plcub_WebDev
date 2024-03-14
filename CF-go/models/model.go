package model

type BlogEntry struct {
	Rating       int    `json:"rating"`
	AuthorHandle string `json:"authorHandle"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
}

type Comment struct {
	ID                int    `json:"id"`
	CommentatorHandle string `json:"commentatorHandle"`
	Text              string `json:"text"`
	ParentCommentID   int    `json:"parentCommentId"`
	Rating            int    `json:"rating"`
}

type Result struct {
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment"`
	TimeSeconds int64     `json:"timeSeconds"`
}
type RecentActionApiResponse struct {
	Result []Result `json:"result"`
}

type TimeStamp struct {
	TimeSeconds int64 `json:"timeseconds"`
}

type User struct {
	Username          string `json:"username"`
	Email             string `json:"email"`
	Codeforces_Handle string `json:"codeforces_handle"`
	Password          string `json:"password"`
	SubscribedBlogsId []int  `json:"subscribedblogid"`
}

type SubscribedBlogs struct {
	Blogs    BlogEntry
	Comments []Comment
}

type UserSubscribe struct {
	Email  string `json:"email"`
	BlogId int    `json:"blogid"`
}

// AddComment appends a comment to the SubscribedBlogs' Comments field
func (sb *SubscribedBlogs) AddComment(comment Comment) {
	sb.Comments = append(sb.Comments, comment)
}
