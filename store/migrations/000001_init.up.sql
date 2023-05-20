-- Create the table 'users'
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(50)     NOT NULL UNIQUE
    CONSTRAINT CH_user_name CHECK (LENGTH(username) >= 3),
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
    INSERT INTO Users (username, email, password)
    VALUES (p_username, p_email, p_password);
    EXCEPTION
        WHEN unique_violation THEN
            RAISE NOTICE 'The email or username already exists.';
        WHEN check_violation THEN
            RAISE NOTICE 'One or more input parameters violate the constraints.';
END;
$$;

-- Create get user func
CREATE OR REPLACE FUNCTION get_user_by_email_or_username(
    IN email_or_username VARCHAR
)
RETURNS TABLE (
    p_id        INT,
    p_username  VARCHAR,
    p_email     VARCHAR,
    p_password  VARCHAR,
    p_role      SMALLINT,
    p_created_at TIMESTAMP
) LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, username, email, password, role, created_at
    FROM users
    WHERE email = email_or_username OR username = email_or_username;
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

-- Create update session procedure
CREATE OR REPLACE PROCEDURE update_session(
    p_user_id       INT,
    p_refresh_token VARCHAR(40)
) LANGUAGE plpgsql AS $$
BEGIN
    UPDATE sessions
    SET refresh_token = p_refresh_token,
        expires_at = (current_date + interval '7 days')
    WHERE user_id = p_user_id;
END;
$$;

-- Create procedure to delete expired sessions
CREATE OR REPLACE PROCEDURE delete_expired_sessions()
LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM sessions
    WHERE expires_at < NOW();
END;
$$;
