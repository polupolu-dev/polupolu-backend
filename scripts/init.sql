-- comments テーブルの作成
CREATE TABLE comments (
    id UUID PRIMARY KEY,
    reply_to_id UUID NOT NULL,
    replied_ids UUID[],
    user_id UUID NOT NULL,
    content VARCHAR(2047) NOT NULL,
    created_at VARCHAR(19) NOT NULL,
    empathy INTEGER NOT NULL,
    insight INTEGER NOT NULL,
    mediocre INTEGER NOT NULL
);

-- news テーブルの作成
CREATE TABLE news (
    id UUID PRIMARY KEY,
    category VARCHAR(63),
    title VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    summary VARCHAR(1023) NOT NULL,
    published_at DATE NOT NULL,
    empathy INTEGER NOT NULL,
    insight INTEGER NOT NULL,
    mediocre INTEGER NOT NULL,
    comment_ids UUID[]
);

-- users テーブルの作成
CREATE TABLE users (
    id UUID PRIMARY KEY,
    comment_ids UUID[],
    gender VARCHAR(63),
    age_group INTEGER,
    occupation VARCHAR(63),
    political_view VARCHAR(63),
    opinion_tone VARCHAR(63),
    speech_style VARCHAR(63),
    comment_length INTEGER,
    background_knowledge VARCHAR(63),
    emotion VARCHAR(63)
);

