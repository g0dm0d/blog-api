DROP TABLE IF EXISTS articles;
DROP PROCEDURE IF EXISTS create_article;
DROP FUNCTION IF EXISTS search_article_by_tags;
DROP FUNCTION IF EXISTS get_article_by_id;

DROP TABLE IF EXISTS sessions;
DROP PROCEDURE IF EXISTS create_session;
DROP FUNCTION IF EXISTS update_session;
DROP PROCEDURE IF EXISTS delete_expired_sessions;

DROP TABLE IF EXISTS users;
DROP PROCEDURE IF EXISTS insert_user;
DROP FUNCTION IF EXISTS get_user_by_email_or_username;
DROP FUNCTION IF EXISTS get_user_by_id;
