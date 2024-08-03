CREATE TABLE IF NOT EXISTS users
(
    id              UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    username        VARCHAR(50) UNIQUE  NOT NULL,
    email           VARCHAR(100) UNIQUE NOT NULL,
    password_hash   VARCHAR(255)        NOT NULL,
    full_name       VARCHAR(100),
    native_language VARCHAR(5),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at      BIGINT                   DEFAULT 0,
    UNIQUE (username, deleted_at)
);

CREATE TABLE IF NOT EXISTS user_languages
(
    id                UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id           UUID REFERENCES users (id),
    language_code     VARCHAR(5)  NOT NULL,
    proficiency_level VARCHAR(20) NOT NULL,
    started_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, language_code)
);