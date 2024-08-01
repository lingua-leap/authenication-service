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
    deleted_at      BIGINT                   DEFAULT 0
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

CREATE TABLE IF NOT EXISTS lessons
(
    id            UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    language_code VARCHAR(5)   NOT NULL,
    title         VARCHAR(100) NOT NULL,
    level         VARCHAR(20)  NOT NULL,
    content       JSONB        NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at    BIGINT                   DEFAULT 0
);

CREATE TABLE IF NOT EXISTS user_lessons
(
    id           UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id      UUID REFERENCES users (id),
    lesson_id    UUID REFERENCES lessons (id),
    completed_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (user_id, lesson_id)
);

CREATE TABLE vocabulary
(
    id               UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    language_code    VARCHAR(5)   NOT NULL,
    lesson_id        uuid references lessons (id),
    word             VARCHAR(100) NOT NULL,
    translation      VARCHAR(100) NOT NULL,
    example_sentence TEXT,
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at       BIGINT                   DEFAULT 0
);

CREATE TABLE user_vocabulary
(
    id               UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id          UUID REFERENCES users (id),
    vocabulary_id    UUID REFERENCES vocabulary (id),
    learned_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_reviewed_at TIMESTAMP WITH TIME ZONE,
    mastery_level    INTEGER                  DEFAULT 0,
    UNIQUE (user_id, vocabulary_id)
);