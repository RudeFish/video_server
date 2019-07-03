package defs

// request
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// 用户信息
type UserInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// 用户
type User struct {
	Id int
	LoginName string
	Pwd string
}

// Resqonse
type SignedUp struct {
	Sucess bool `json:"sucess"`
	SessionId string `json:"session_id"`
} 

// data model

// 视频
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string // 页面上显示的时间
}

// 评论
type Comment struct {
	Id string
	VideoId string
	AuthorName string
	Content string
}

// session
type SimpleSession struct {
	Username string // login name
	TTL int64 // 判断session是否过期
}

