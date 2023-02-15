package repository

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type LoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type RegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}

type FeedResponse struct {
	Response
	NextTime  int64   `json:"next_time,omitempty"`
	VideoList []Video `json:"video_list,omitempty"`
}

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

type MessageListResponse struct {
	Response
	MessageList []Message `json:"message_list,omitempty"`
}
