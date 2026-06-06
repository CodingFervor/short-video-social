# Short Video Social Platform

## English | [中文](#中文)

A short video social platform backend built with Go, Gin, PostgreSQL, and Redis.

### Features
- User registration, login, profile management
- Short video upload, publishing, and management
- Video feed with recommendation algorithm support
- Like, comment, share, and bookmark
- User follow/unfollow system
- Real-time messaging (chat)
- Live streaming support
- Content moderation and reporting
- Hashtags and trending topics
- Notification system

### Tech Stack
- Go 1.22 + Gin
- PostgreSQL 16
- Redis 7
- JWT Authentication
- Docker Compose

### Quick Start

```bash
# Start dependencies
docker-compose up -d

# Run the server
go run cmd/api/main.go
```

### API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | /api/v1/auth/register | Register |
| POST | /api/v1/auth/login | Login |
| GET | /api/v1/feed | Get video feed |
| POST | /api/v1/videos | Upload video |
| GET | /api/v1/videos/:id | Get video detail |
| POST | /api/v1/videos/:id/like | Like video |
| POST | /api/v1/videos/:id/comments | Add comment |
| POST | /api/v1/users/:id/follow | Follow user |
| GET | /api/v1/messages | Get messages |
| POST | /api/v1/live/start | Start live stream |

---

<a id="中文"></a>
# 短视频社交平台

基于 Go + Gin + PostgreSQL + Redis 构建的短视频社交平台后端。

### 功能特性
- 用户注册、登录、个人资料管理
- 短视频上传、发布和管理
- 视频流推荐算法支持
- 点赞、评论、分享、收藏
- 用户关注/取关系统
- 实时消息（聊天）
- 直播支持
- 内容审核和举报
- 话题标签和热门趋势
- 通知系统

### 技术栈
- Go 1.22 + Gin
- PostgreSQL 16
- Redis 7
- JWT 认证
- Docker Compose

### 快速开始

```bash
# 启动依赖服务
docker-compose up -d

# 运行服务
go run cmd/api/main.go
```

### API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/v1/auth/register | 用户注册 |
| POST | /api/v1/auth/login | 用户登录 |
| GET | /api/v1/feed | 获取视频流 |
| POST | /api/v1/videos | 上传视频 |
| GET | /api/v1/videos/:id | 获取视频详情 |
| POST | /api/v1/videos/:id/like | 点赞视频 |
| POST | /api/v1/videos/:id/comments | 添加评论 |
| POST | /api/v1/users/:id/follow | 关注用户 |
| GET | /api/v1/messages | 获取消息列表 |
| POST | /api/v1/live/start | 开始直播 |
