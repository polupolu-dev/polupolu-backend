-- comments テーブルの作成
CREATE TABLE comments (
    id UUID PRIMARY KEY,
    reply_to_id UUID NOT NULL,
    replied_ids UUID[],
    user_id VARCHAR(16) NOT NULL,
    content VARCHAR(2047) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    empathy INTEGER NOT NULL CHECK (empathy >= 0),
    insight INTEGER NOT NULL CHECK (insight >= 0),
    mediocre INTEGER NOT NULL CHECK (mediocre >= 0)
);

-- news テーブルの作成
CREATE TABLE news (
    id UUID PRIMARY KEY,
    category VARCHAR(63),
    title VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    summary VARCHAR(1023) NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    empathy INTEGER NOT NULL CHECK (empathy >= 0),
    insight INTEGER NOT NULL CHECK (insight >= 0),
    mediocre INTEGER NOT NULL CHECK (mediocre >= 0),
    comment_ids UUID[]
);

-- users テーブルの作成
CREATE TABLE users (
    id VARCHAR(16) PRIMARY KEY,
    comment_ids UUID[],
    gender VARCHAR(63),
    age_group INTEGER,
    occupation VARCHAR(63),
    political_view VARCHAR(63),
    opinion_tone VARCHAR(63),
    speech_style VARCHAR(63),
    comment_length INTEGER CHECK (comment_length >= 0),
    background_knowledge VARCHAR(63),
    emotion VARCHAR(63)
);
