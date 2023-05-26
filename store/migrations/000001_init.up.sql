-- Create the table 'users'
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(50)     NOT NULL UNIQUE
    CONSTRAINT CH_user_name CHECK (LENGTH(username) >= 3),
    name        VARCHAR(100)    NOT NULL,
    avatar      VARCHAR(200)    DEFAULT NULL,
    bio         VARCHAR(200)    DEFAULT NULL,
    email       VARCHAR(100)    NOT NULL UNIQUE
    CONSTRAINT CH_user_email CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password    VARCHAR(100)    NOT NULL,
    role        SMALLINT        NOT NULL DEFAULT 1,
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

-- Create insert user procedure
CREATE OR REPLACE PROCEDURE create_user(
    p_username  VARCHAR(50),
    p_email     VARCHAR(100),
    p_password  VARCHAR(50)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO Users (username, name, email, password)
    VALUES (p_username, p_username, p_email, p_password);
    EXCEPTION
        WHEN unique_violation THEN
            RAISE EXCEPTION 'The email or username already exists.';
        WHEN check_violation THEN
            RAISE EXCEPTION 'One or more input parameters violate the constraints.';
END;
$$;

-- Create get user by email or username func
CREATE OR REPLACE FUNCTION get_user_by_email_or_username(
    IN email_or_username VARCHAR
)
RETURNS TABLE (
    p_id        INT,
    p_username  VARCHAR,
    p_name      VARCHAR,
    p_avatar    VARCHAR,
    p_bio       VARCHAR,
    p_email     VARCHAR,
    p_password  VARCHAR,
    p_role      SMALLINT,
    p_created_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE email = email_or_username OR username = email_or_username;
END;
$$;

-- Create get user by id func
CREATE OR REPLACE FUNCTION get_user_by_id(
    IN search_id INT
)
RETURNS TABLE (
    p_id        INT,
    p_username  VARCHAR,
    p_name      VARCHAR,
    p_avatar    VARCHAR,
    p_bio       VARCHAR,
    p_email     VARCHAR,
    p_password  VARCHAR,
    p_role      SMALLINT,
    p_created_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE id = search_id;
END;
$$;

-- Create get user by username func
CREATE OR REPLACE FUNCTION get_user_by_username(
    IN v_username VARCHAR
)
RETURNS TABLE (
    p_id        INT,
    p_username  VARCHAR,
    p_name      VARCHAR,
    p_bio       VARCHAR,
    p_email     VARCHAR,
    p_password  VARCHAR,
    p_role      SMALLINT,
    p_created_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, username, name, avatar, bio, email, password, role, created_at
    FROM users
    WHERE username = username;
END;
$$;



-- Create the table 'sessions'
CREATE TABLE IF NOT EXISTS sessions (
	id				UUID		PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         INT         NOT NULL REFERENCES users(id),
    refresh_token   VARCHAR(40) NOT NULL,
    created_at      TIMESTAMP   NOT NULL DEFAULT NOW(),
    expires_at      TIMESTAMP   NOT NULL DEFAULT (current_date + interval '7 days')
);

-- Create insert session procedure
CREATE OR REPLACE PROCEDURE create_session(
    p_user_id       INT,
    p_refresh_token VARCHAR(40)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO sessions (user_id, refresh_token)
    VALUES (p_user_id, p_refresh_token);
END;
$$;

-- Update refresh token procedure
CREATE OR REPLACE FUNCTION update_session(
    p_new_token VARCHAR(40),
    p_old_token VARCHAR(40)
) RETURNS INT AS $$
DECLARE
    p_user_id INT;
BEGIN
    IF EXISTS (SELECT 1 FROM sessions WHERE refresh_token = p_old_token) THEN
        UPDATE sessions
        SET refresh_token = p_new_token,
            expires_at = (current_date + interval '7 days')
        WHERE refresh_token = p_old_token
        RETURNING user_id INTO p_user_id;
        RETURN p_user_id;
    ELSE
        RAISE EXCEPTION 'Session not found with token: %', p_old_token;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Create procedure to delete expired sessions
CREATE OR REPLACE PROCEDURE delete_expired_sessions()
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM sessions
    WHERE expires_at < NOW();
END;
$$;



-- Create the table 'articles'
CREATE TABLE IF NOT EXISTS articles (
    id          SERIAL          PRIMARY KEY,
    title       VARCHAR(100)    NOT NULL,
    path        VARCHAR(105)    NOT NULL,
    markdown    TEXT            NOT NULL,
    tags        TEXT[],
    preview     VARCHAR(200),
    author_id   INT             NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

-- Create insert article procedure
CREATE OR REPLACE PROCEDURE create_article(
    p_title     VARCHAR(100),
    p_path      VARCHAR(105),
    p_markdown  TEXT,
    p_tags      TEXT[],
    p_preview   VARCHAR(100),
    p_author    INT
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO articles (title, path, markdown, tags, preview, author_id)
    VALUES (p_title, p_path, p_markdown, p_tags, p_preview, p_author);
END;
$$;

-- Create search article by tags func
CREATE OR REPLACE FUNCTION search_article_by_tags(
    IN v_tags TEXT[]
)
RETURNS TABLE (
    p_id        INT,
    p_path      VARCHAR,
    p_title     VARCHAR,
    p_tags      TEXT[],
    p_preview   VARCHAR,
    p_author    INT
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, title, path, tags, preview, author_id
    FROM articles
    WHERE tags @> v_tags;
END;
$$;

-- Create get article by id func
CREATE OR REPLACE FUNCTION get_article_by_id(
    IN v_id INT
)
RETURNS TABLE (
    p_title         VARCHAR,
    p_path          VARCHAR,
    p_markdown      TEXT,
    p_tags          TEXT[],
    p_preview       VARCHAR,
    p_author        INT,
    p_created_at    TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT title, path, markdown, tags, preview, author_id, created_at
    FROM articles
    WHERE id = v_id;
END;
$$;

-- Create get article by path func
CREATE OR REPLACE FUNCTION get_article_by_path(
    IN v_path VARCHAR
)
RETURNS TABLE (
    p_title         VARCHAR,
    p_markdown      TEXT,
    p_tags          TEXT[],
    p_preview       VARCHAR,
    p_author        INT,
    p_created_at    TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT title, markdown, tags, preview, author_id, created_at
    FROM articles
    WHERE path = v_path;
END;
$$;

-- Create get article for feed
CREATE OR REPLACE FUNCTION get_article_feed(
    IN v_last INT
)
RETURNS TABLE (
    p_title         VARCHAR,
    p_path          VARCHAR,
    p_markdown      TEXT,
    p_tags          TEXT[],
    p_preview       VARCHAR,
    p_author        INT,
    p_created_at    TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT title, path, markdown, tags, preview, author_id, created_at
    FROM articles
    WHERE id <= (SELECT (id - 15 * v_last) FROM articles ORDER BY 1 DESC LIMIT 1)
    ORDER BY id desc limit 15;
END;
$$;

