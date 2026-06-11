package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CodingFervor/short-video-social/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().Format(time.RFC3339)})
	})

	api := r.Group("/api/v1")
	{
		// Auth
		api.POST("/auth/register", Register)
		api.POST("/auth/login", Login)

		auth := api.Group("/")
		auth.Use(AuthMiddleware())
		{
			// User profile
			auth.GET("/users/profile", GetProfile)
			auth.PUT("/users/profile", UpdateProfile)
			auth.GET("/users/:id", GetUserProfile)

			// Video feed
			auth.GET("/feed", GetFeed)
			auth.GET("/feed/following", GetFollowingFeed)
			auth.GET("/feed/trending", GetTrendingFeed)

			// Videos
			auth.POST("/videos", UploadVideo)
			auth.GET("/videos/:id", GetVideo)
			auth.PUT("/videos/:id", UpdateVideo)
			auth.DELETE("/videos/:id", DeleteVideo)
			auth.GET("/videos/:id/comments", ListComments)
			auth.POST("/videos/:id/comments", AddComment)
			auth.DELETE("/videos/:id/comments/:cid", DeleteComment)

			// Interactions
			auth.POST("/videos/:id/like", LikeVideo)
			auth.DELETE("/videos/:id/like", UnlikeVideo)
			auth.POST("/videos/:id/share", ShareVideo)
			auth.POST("/videos/:id/bookmark", BookmarkVideo)
			auth.DELETE("/videos/:id/bookmark", RemoveBookmark)

			// Follow system
			auth.POST("/users/:id/follow", FollowUser)
			auth.DELETE("/users/:id/follow", UnfollowUser)
			auth.GET("/users/:id/followers", GetFollowers)
			auth.GET("/users/:id/following", GetFollowing)

			// Messages / Chat
			auth.GET("/messages", GetMessages)
			auth.POST("/messages", SendMessage)
			auth.GET("/conversations", GetConversations)

			// Hashtags
			auth.GET("/hashtags/trending", GetTrendingHashtags)
			auth.GET("/hashtags/:tag/videos", GetHashtagVideos)

			// Notifications
			auth.GET("/notifications", GetNotifications)
			auth.PUT("/notifications/:id/read", MarkNotificationRead)

			// Live streaming
			auth.POST("/live/start", StartLive)
			auth.POST("/live/:id/end", EndLive)
			auth.GET("/live/rooms", ListLiveRooms)
			auth.POST("/live/:id/join", JoinLive)
			auth.POST("/live/:id/gift", SendGift)

			// Search
			auth.GET("/search/videos", SearchVideos)
			auth.GET("/search/users", SearchUsers)
			auth.GET("/search/hashtags", SearchHashtags)

			// Reports / Moderation
			auth.POST("/reports", CreateReport)

			// Admin
			admin := auth.Group("/admin")
			admin.Use(AdminMiddleware())
			{
				admin.GET("/reports", ListReports)
				admin.PUT("/reports/:id", HandleReport)
				admin.PUT("/videos/:id/status", UpdateVideoStatus)
			}
		}
	}

	log.Println("Short Video Social server starting on :8080")
	if err := addr := ":" + strconv.Itoa(8080)
	srv := &http.Server{Addr: addr, Handler: r}
	go func() {
		logger.Info("server listening", "port", 8080)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("forced shutdown", "error", err)
	}
	logger.Info("server exited"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// --- Handler stubs ---

func Login(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "login"}) }
func Register(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "registered"}) }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func GetProfile(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func UpdateProfile(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"message": "profile updated"}) }
func GetUserProfile(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }

func GetFeed(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func GetFollowingFeed(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func GetTrendingFeed(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }

func UploadVideo(c *gin.Context)     { c.JSON(http.StatusCreated, gin.H{"message": "video uploaded"}) }
func GetVideo(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func UpdateVideo(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "video updated"}) }
func DeleteVideo(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "video deleted"}) }

func ListComments(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func AddComment(c *gin.Context)      { c.JSON(http.StatusCreated, gin.H{"message": "comment added"}) }
func DeleteComment(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"message": "comment deleted"}) }

func LikeVideo(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"message": "liked"}) }
func UnlikeVideo(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "unliked"}) }
func ShareVideo(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"message": "shared"}) }
func BookmarkVideo(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"message": "bookmarked"}) }
func RemoveBookmark(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"message": "bookmark removed"}) }

func FollowUser(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"message": "followed"}) }
func UnfollowUser(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "unfollowed"}) }
func GetFollowers(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func GetFollowing(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }

func GetMessages(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func SendMessage(c *gin.Context)     { c.JSON(http.StatusCreated, gin.H{"message": "sent"}) }
func GetConversations(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }

func GetTrendingHashtags(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func GetHashtagVideos(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }

func GetNotifications(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func MarkNotificationRead(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "marked read"}) }

func StartLive(c *gin.Context)   { c.JSON(http.StatusCreated, gin.H{"message": "live started"}) }
func EndLive(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "live ended"}) }
func ListLiveRooms(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func JoinLive(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "joined"}) }
func SendGift(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "gift sent"}) }

func SearchVideos(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func SearchUsers(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func SearchHashtags(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }

func CreateReport(c *gin.Context)   { c.JSON(http.StatusCreated, gin.H{"message": "report created"}) }

func ListReports(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func HandleReport(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"message": "report handled"}) }
func UpdateVideoStatus(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"message": "status updated"}) }
