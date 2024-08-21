CREATE TABLE IF NOT EXISTS users
(
    id              UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    username        VARCHAR(50) UNIQUE  NOT NULL,
    email           VARCHAR(100) UNIQUE NOT NULL,
    password_hash   VARCHAR(255)        NOT NULL,
    full_name       VARCHAR(100),
    native_language VARCHAR(5),
    role            VARCHAR(10) DEFAULT 'user',
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at      BIGINT                   DEFAULT 0,
    UNIQUE (username, deleted_at)
);
