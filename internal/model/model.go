package model

import "time"

// User represents a platform user
type User struct {
	ID            int64     `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	Password      string    `json:"-" db:"password"`
	Nickname      string    `json:"nickname" db:"nickname"`
	Avatar        string    `json:"avatar" db:"avatar"`
	Bio           string    `json:"bio" db:"bio"`
	Gender        string    `json:"gender" db:"gender"`
	Birthday      *time.Time `json:"birthday" db:"birthday"`
	Phone         string    `json:"phone" db:"phone"`
	Email         string    `json:"email" db:"email"`
	Region        string    `json:"region" db:"region"`
	FollowerCount int       `json:"follower_count" db:"follower_count"`
	FollowingCount int      `json:"following_count" db:"following_count"`
	LikeCount     int       `json:"like_count" db:"like_count"`
	VideoCount    int       `json:"video_count" db:"video_count"`
	IsVerified    bool      `json:"is_verified" db:"is_verified"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Video represents a short video
type Video struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	VideoURL     string    `json:"video_url" db:"video_url"`
	CoverURL     string    `json:"cover_url" db:"cover_url"`
	Duration     int       `json:"duration" db:"duration"`      // seconds
	Width        int       `json:"width" db:"width"`
	Height       int       `json:"height" db:"height"`
	PlayCount    int64     `json:"play_count" db:"play_count"`
	LikeCount    int       `json:"like_count" db:"like_count"`
	CommentCount int       `json:"comment_count" db:"comment_count"`
	ShareCount   int       `json:"share_count" db:"share_count"`
	BookmarkCount int      `json:"bookmark_count" db:"bookmark_count"`
	Visibility   string    `json:"visibility" db:"visibility"`   // public, friends, private
	Hashtags     string    `json:"hashtags" db:"hashtags"`
	MusicID      *int64    `json:"music_id" db:"music_id"`
	Status       string    `json:"status" db:"status"`           // active, hidden, deleted, under_review
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Comment represents a video comment
type Comment struct {
	ID        int64     `json:"id" db:"id"`
	VideoID   int64     `json:"video_id" db:"video_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	ParentID  *int64    `json:"parent_id" db:"parent_id"`    // for replies
	Content   string    `json:"content" db:"content"`
	LikeCount int       `json:"like_count" db:"like_count"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Like represents a like on video or comment
type Like struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	TargetID  int64     `json:"target_id" db:"target_id"`
	TargetType string   `json:"target_type" db:"target_type"` // video, comment
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Follow represents a follow relationship
type Follow struct {
	ID          int64     `json:"id" db:"id"`
	FollowerID  int64     `json:"follower_id" db:"follower_id"`
	FollowingID int64     `json:"following_id" db:"following_id"`
	Status      string    `json:"status" db:"status"`    // active, blocked
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Bookmark represents a saved/bookmarked video
type Bookmark struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	VideoID   int64     `json:"video_id" db:"video_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Share represents a video share record
type Share struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	VideoID   int64     `json:"video_id" db:"video_id"`
	Platform  string    `json:"platform" db:"platform"`  // wechat, weibo, qq, link
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Message represents a direct message
type Message struct {
	ID             int64     `json:"id" db:"id"`
	ConversationID int64     `json:"conversation_id" db:"conversation_id"`
	SenderID       int64     `json:"sender_id" db:"sender_id"`
	ReceiverID     int64     `json:"receiver_id" db:"receiver_id"`
	Content        string    `json:"content" db:"content"`
	Type           string    `json:"type" db:"type"`   // text, image, video, gift
	IsRead         bool      `json:"is_read" db:"is_read"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Conversation represents a chat conversation
type Conversation struct {
	ID           int64     `json:"id" db:"id"`
	User1ID      int64     `json:"user1_id" db:"user1_id"`
	User2ID      int64     `json:"user2_id" db:"user2_id"`
	LastMessage  string    `json:"last_message" db:"last_message"`
	LastMsgTime  time.Time `json:"last_msg_time" db:"last_msg_time"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Hashtag represents a trending hashtag
type Hashtag struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	UseCount   int       `json:"use_count" db:"use_count"`
	HotScore   float64   `json:"hot_score" db:"hot_score"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Notification represents a user notification
type Notification struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	ActorID    int64     `json:"actor_id" db:"actor_id"`
	Type       string    `json:"type" db:"type"`        // like, comment, follow, mention, system
	Content    string    `json:"content" db:"content"`
	TargetID   int64     `json:"target_id" db:"target_id"`
	TargetType string    `json:"target_type" db:"target_type"`
	IsRead     bool      `json:"is_read" db:"is_read"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// LiveRoom represents a live streaming room
type LiveRoom struct {
	ID          int64     `json:"id" db:"id"`
	HostID      int64     `json:"host_id" db:"host_id"`
	Title       string    `json:"title" db:"title"`
	StreamKey   string    `json:"stream_key" db:"stream_key"`
	PlayURL     string    `json:"play_url" db:"play_url"`
	CoverURL    string    `json:"cover_url" db:"cover_url"`
	ViewerCount int       `json:"viewer_count" db:"viewer_count"`
	LikeCount   int       `json:"like_count" db:"like_count"`
	GiftAmount  float64   `json:"gift_amount" db:"gift_amount"`
	Status      string    `json:"status" db:"status"`    // live, ended
	StartedAt   time.Time `json:"started_at" db:"started_at"`
	EndedAt     *time.Time `json:"ended_at" db:"ended_at"`
}

// Gift represents a virtual gift
type Gift struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Icon      string    `json:"icon" db:"icon"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// GiftRecord represents a gift sent in live
type GiftRecord struct {
	ID        int64     `json:"id" db:"id"`
	SenderID  int64     `json:"sender_id" db:"sender_id"`
	ReceiverID int64    `json:"receiver_id" db:"receiver_id"`
	GiftID    int64     `json:"gift_id" db:"gift_id"`
	LiveID    int64     `json:"live_id" db:"live_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Amount    float64   `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Report represents a content report
type Report struct {
	ID          int64     `json:"id" db:"id"`
	ReporterID  int64     `json:"reporter_id" db:"reporter_id"`
	TargetID    int64     `json:"target_id" db:"target_id"`
	TargetType  string    `json:"target_type" db:"target_type"`  // video, comment, user
	Reason      string    `json:"reason" db:"reason"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`    // pending, reviewed, dismissed, actioned
	ReviewedBy  *int64    `json:"reviewed_by" db:"reviewed_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Music represents background music for videos
type Music struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Artist    string    `json:"artist" db:"artist"`
	AudioURL  string    `json:"audio_url" db:"audio_url"`
	CoverURL  string    `json:"cover_url" db:"cover_url"`
	Duration  int       `json:"duration" db:"duration"`
	UseCount  int       `json:"use_count" db:"use_count"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
