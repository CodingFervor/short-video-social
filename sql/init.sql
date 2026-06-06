-- Short Video Social Platform - Database Schema
-- PostgreSQL

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    nickname VARCHAR(100) NOT NULL,
    avatar VARCHAR(500),
    bio TEXT,
    gender VARCHAR(10) DEFAULT 'unknown',
    birthday DATE,
    phone VARCHAR(20) UNIQUE,
    email VARCHAR(100) UNIQUE,
    region VARCHAR(50),
    follower_count INT DEFAULT 0,
    following_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    video_count INT DEFAULT 0,
    is_verified BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active','banned','deactivated')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Videos
CREATE TABLE videos (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    title VARCHAR(200),
    description TEXT,
    video_url VARCHAR(500) NOT NULL,
    cover_url VARCHAR(500),
    duration INT DEFAULT 0,
    width INT DEFAULT 0,
    height INT DEFAULT 0,
    play_count BIGINT DEFAULT 0,
    like_count INT DEFAULT 0,
    comment_count INT DEFAULT 0,
    share_count INT DEFAULT 0,
    bookmark_count INT DEFAULT 0,
    visibility VARCHAR(20) DEFAULT 'public' CHECK (visibility IN ('public','friends','private')),
    hashtags TEXT,
    music_id BIGINT,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active','hidden','deleted','under_review')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_videos_user ON videos(user_id);
CREATE INDEX idx_videos_status ON videos(status);
CREATE INDEX idx_videos_created ON videos(created_at DESC);

-- Comments
CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    video_id BIGINT NOT NULL REFERENCES videos(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    parent_id BIGINT REFERENCES comments(id),
    content TEXT NOT NULL,
    like_count INT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_comments_video ON comments(video_id);

-- Likes
CREATE TABLE likes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    target_id BIGINT NOT NULL,
    target_type VARCHAR(20) NOT NULL CHECK (target_type IN ('video','comment')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, target_id, target_type)
);

CREATE INDEX idx_likes_user ON likes(user_id);
CREATE INDEX idx_likes_target ON likes(target_id, target_type);

-- Follows
CREATE TABLE follows (
    id BIGSERIAL PRIMARY KEY,
    follower_id BIGINT NOT NULL REFERENCES users(id),
    following_id BIGINT NOT NULL REFERENCES users(id),
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(follower_id, following_id)
);

CREATE INDEX idx_follows_follower ON follows(follower_id);
CREATE INDEX idx_follows_following ON follows(following_id);

-- Bookmarks
CREATE TABLE bookmarks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    video_id BIGINT NOT NULL REFERENCES videos(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, video_id)
);

-- Shares
CREATE TABLE shares (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    video_id BIGINT NOT NULL REFERENCES videos(id),
    platform VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Conversations
CREATE TABLE conversations (
    id BIGSERIAL PRIMARY KEY,
    user1_id BIGINT NOT NULL REFERENCES users(id),
    user2_id BIGINT NOT NULL REFERENCES users(id),
    last_message TEXT,
    last_msg_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user1_id, user2_id)
);

-- Messages
CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    conversation_id BIGINT NOT NULL REFERENCES conversations(id),
    sender_id BIGINT NOT NULL REFERENCES users(id),
    receiver_id BIGINT NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    type VARCHAR(20) DEFAULT 'text' CHECK (type IN ('text','image','video','gift')),
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_messages_conversation ON messages(conversation_id);

-- Hashtags
CREATE TABLE hashtags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    use_count INT DEFAULT 0,
    hot_score DECIMAL(10,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Video hashtags junction
CREATE TABLE video_hashtags (
    video_id BIGINT NOT NULL REFERENCES videos(id) ON DELETE CASCADE,
    hashtag_id BIGINT NOT NULL REFERENCES hashtags(id),
    PRIMARY KEY (video_id, hashtag_id)
);

-- Notifications
CREATE TABLE notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    actor_id BIGINT REFERENCES users(id),
    type VARCHAR(20) NOT NULL CHECK (type IN ('like','comment','follow','mention','system')),
    content TEXT,
    target_id BIGINT,
    target_type VARCHAR(20),
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notifications_user ON notifications(user_id, is_read);

-- Music
CREATE TABLE music (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    artist VARCHAR(200),
    audio_url VARCHAR(500) NOT NULL,
    cover_url VARCHAR(500),
    duration INT DEFAULT 0,
    use_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Live rooms
CREATE TABLE live_rooms (
    id BIGSERIAL PRIMARY KEY,
    host_id BIGINT NOT NULL REFERENCES users(id),
    title VARCHAR(200),
    stream_key VARCHAR(100) UNIQUE,
    play_url VARCHAR(500),
    cover_url VARCHAR(500),
    viewer_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    gift_amount DECIMAL(10,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'live' CHECK (status IN ('live','ended')),
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ended_at TIMESTAMP
);

CREATE INDEX idx_live_status ON live_rooms(status);

-- Gifts
CREATE TABLE gifts (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(500),
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Gift records
CREATE TABLE gift_records (
    id BIGSERIAL PRIMARY KEY,
    sender_id BIGINT NOT NULL REFERENCES users(id),
    receiver_id BIGINT NOT NULL REFERENCES users(id),
    gift_id BIGINT NOT NULL REFERENCES gifts(id),
    live_id BIGINT REFERENCES live_rooms(id),
    quantity INT DEFAULT 1,
    amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reports
CREATE TABLE reports (
    id BIGSERIAL PRIMARY KEY,
    reporter_id BIGINT NOT NULL REFERENCES users(id),
    target_id BIGINT NOT NULL,
    target_type VARCHAR(20) NOT NULL CHECK (target_type IN ('video','comment','user')),
    reason VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending','reviewed','dismissed','actioned')),
    reviewed_by BIGINT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default admin user
INSERT INTO users (username, password, nickname, email, role) VALUES
('admin', '$2a$10$dummyhash', 'Admin', 'admin@video.com', 'admin');

-- Insert sample gifts
INSERT INTO gifts (name, icon, price) VALUES
('Heart', 'heart.png', 1.00),
('Rose', 'rose.png', 5.00),
('Crown', 'crown.png', 50.00),
('Rocket', 'rocket.png', 100.00);
