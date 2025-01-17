-- comments テーブルの作成
CREATE TABLE comments (
    id UUID PRIMARY KEY,
    reply_to_id UUID NOT NULL,
    replied_ids UUID[],
    user_id VARCHAR(16) NOT NULL,
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
    id VARCHAR(16) PRIMARY KEY,
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


-- 制約の追加 (必要に応じて)
-- comments テーブルの user_id が users テーブルに存在することを保証する外部キー制約を追加
ALTER TABLE comments ADD CONSTRAINT fk_comments_user_id FOREIGN KEY (user_id) REFERENCES users(id);

-- comments テーブルの reply_to_id が news テーブルもしくは comments テーブルに存在することを保証する制約を追加 (複雑なためトリガーで実装)
-- newsテーブルの comment_ids が comments テーブルに存在することを保証する外部キー制約を追加
-- ALTER TABLE news ADD CONSTRAINT fk_news_comment_ids FOREIGN KEY (comment_ids) REFERENCES comments(id);
-- ユーザーの制約は後で定義

-- UUIDの検証用関数 (comments, news で共通)
CREATE OR REPLACE FUNCTION is_valid_uuid(input TEXT)
RETURNS BOOLEAN
AS $$
BEGIN
    RETURN input ~ '^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$';
END;
$$ LANGUAGE plpgsql;


-- トリガー関数の定義: commentsテーブルのreply_to_idがnewsテーブルかcommentsテーブルに存在するか検証
CREATE OR REPLACE FUNCTION check_reply_to_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.reply_to_id IS NOT NULL THEN
        IF NOT EXISTS (SELECT 1 FROM news WHERE id = NEW.reply_to_id::UUID) AND NOT EXISTS (SELECT 1 FROM comments WHERE id = NEW.reply_to_id::UUID) THEN
            RAISE EXCEPTION 'reply_to_id must reference a record in news or comments';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- トリガーの作成: commentsテーブルでINSERTまたはUPDATEが発生したときに、check_reply_to_id関数を呼び出す
CREATE TRIGGER check_reply_to_id_trigger
BEFORE INSERT OR UPDATE ON comments
FOR EACH ROW
EXECUTE FUNCTION check_reply_to_id();


-- トリガー関数の定義: newsテーブルのcomment_idsがcommentsテーブルに存在するか検証
CREATE OR REPLACE FUNCTION check_news_comment_ids()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.comment_ids IS NOT NULL THEN
        FOREACH id_text IN ARRAY NEW.comment_ids
        LOOP
            IF NOT EXISTS (SELECT 1 FROM comments WHERE id = id_text::UUID) THEN
                RAISE EXCEPTION 'comment_ids must reference records in comments';
            END IF;
        END LOOP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- トリガーの作成: newsテーブルでINSERTまたはUPDATEが発生したときに、check_news_comment_ids関数を呼び出す
CREATE TRIGGER check_news_comment_ids_trigger
BEFORE INSERT OR UPDATE ON news
FOR EACH ROW
EXECUTE FUNCTION check_news_comment_ids();


-- トリガー関数の定義: usersテーブルのcomment_idsがcommentsテーブルに存在するか検証
CREATE OR REPLACE FUNCTION check_user_comment_ids()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.comment_ids IS NOT NULL THEN
        FOREACH id_text IN ARRAY NEW.comment_ids
        LOOP
            IF NOT EXISTS (SELECT 1 FROM comments WHERE id = id_text::UUID) THEN
                RAISE EXCEPTION 'comment_ids must reference records in comments';
            END IF;
        END LOOP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- トリガーの作成: usersテーブルでINSERTまたはUPDATEが発生したときに、check_user_comment_ids関数を呼び出す
CREATE TRIGGER check_user_comment_ids_trigger
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION check_user_comment_ids();


-- 制約の追加 (JSONスキーマに基づく)
-- commentsテーブルのidカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_id_format CHECK (is_valid_uuid(id::TEXT));
-- ALTER TABLE comments ADD CONSTRAINT check_comments_id_length CHECK (LENGTH(id::TEXT) = 36);
-- reply_to_idカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_reply_to_id_format CHECK (is_valid_uuid(reply_to_id::TEXT));
-- ALTER TABLE comments ADD CONSTRAINT check_comments_reply_to_id_length CHECK (LENGTH(reply_to_id::TEXT) = 36);
-- replied_idsカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_replied_ids_length CHECK (array_length(replied_ids, 1) IS NULL OR array_length(replied_ids, 1) <= 1023);
-- user_idカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_user_id_format CHECK (user_id ~ '^[A-Z0-9]{1}-[A-Z0-9]{2}-[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}$');
ALTER TABLE comments ADD CONSTRAINT check_comments_user_id_length CHECK (LENGTH(user_id) = 16);
-- contentカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_content_length CHECK (LENGTH(content) BETWEEN 1 AND 2047);
-- created_atカラムの制約
ALTER TABLE comments ADD CONSTRAINT check_comments_created_at_format CHECK (created_at ~ '^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$');
ALTER TABLE comments ADD CONSTRAINT check_comments_created_at_length CHECK (LENGTH(created_at) = 19);
-- feedback_scores の制約 (empathy, insight, mediocre)
ALTER TABLE comments ADD CONSTRAINT check_comments_empathy_range CHECK (empathy >= 0 AND empathy <= 4294967295);
ALTER TABLE comments ADD CONSTRAINT check_comments_insight_range CHECK (insight >= 0 AND insight <= 4294967295);
ALTER TABLE comments ADD CONSTRAINT check_comments_mediocre_range CHECK (mediocre >= 0 AND mediocre <= 4294967295);

-- newsテーブルの制約
ALTER TABLE news ADD CONSTRAINT check_news_id_format CHECK (is_valid_uuid(id::TEXT));
-- ALTER TABLE news ADD CONSTRAINT check_news_id_length CHECK (LENGTH(id::TEXT) = 36);
ALTER TABLE news ADD CONSTRAINT check_news_category_length CHECK (LENGTH(category) BETWEEN 1 AND 63);
ALTER TABLE news ADD CONSTRAINT check_news_title_length CHECK (LENGTH(title) BETWEEN 1 AND 255);
ALTER TABLE news ADD CONSTRAINT check_news_source_length CHECK (LENGTH(source) BETWEEN 1 AND 255);
ALTER TABLE news ADD CONSTRAINT check_news_url_length CHECK (LENGTH(url) BETWEEN 8 AND 255);
ALTER TABLE news ADD CONSTRAINT check_news_summary_length CHECK (LENGTH(summary) BETWEEN 1 AND 1023);
ALTER TABLE news ADD CONSTRAINT check_news_published_at_format CHECK (published_at::TEXT ~ '^\d{4}-\d{2}-\d{2}$');
-- feedback_scores の制約 (empathy, insight, mediocre)
ALTER TABLE news ADD CONSTRAINT check_news_empathy_range CHECK (empathy >= 0 AND empathy <= 4294967295);
ALTER TABLE news ADD CONSTRAINT check_news_insight_range CHECK (insight >= 0 AND insight <= 4294967295);
ALTER TABLE news ADD CONSTRAINT check_news_mediocre_range CHECK (mediocre >= 0 AND mediocre <= 4294967295);
ALTER TABLE news ADD CONSTRAINT check_news_comment_ids_length CHECK (array_length(comment_ids, 1) IS NULL OR array_length(comment_ids, 1) <= 1023);

-- usersテーブルの制約
ALTER TABLE users ADD CONSTRAINT check_users_id_format CHECK (id ~ '^[A-Z0-9]{1}-[A-Z0-9]{2}-[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}$');
ALTER TABLE users ADD CONSTRAINT check_users_id_length CHECK (LENGTH(id) = 16);
ALTER TABLE users ADD CONSTRAINT check_users_comment_ids_length CHECK (array_length(comment_ids, 1) IS NULL OR array_length(comment_ids, 1) <= 1023);
ALTER TABLE users ADD CONSTRAINT check_users_gender_length CHECK (LENGTH(gender) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_age_group_range CHECK (age_group BETWEEN 00 AND 90 AND age_group % 10 = 0);
ALTER TABLE users ADD CONSTRAINT check_users_occupation_length CHECK (LENGTH(occupation) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_political_view_length CHECK (LENGTH(political_view) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_opinion_tone_length CHECK (LENGTH(opinion_tone) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_speech_style_length CHECK (LENGTH(speech_style) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_comment_length_range CHECK (comment_length BETWEEN 1 AND 2047);
ALTER TABLE users ADD CONSTRAINT check_users_background_knowledge_length CHECK (LENGTH(background_knowledge) BETWEEN 1 AND 63);
ALTER TABLE users ADD CONSTRAINT check_users_emotion_length CHECK (LENGTH(emotion) BETWEEN 1 AND 63);
