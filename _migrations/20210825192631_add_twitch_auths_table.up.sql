CREATE TABLE twitch_auths
(
    user_id       TEXT PRIMARY KEY,
    access_token  TEXT      NOT NULL,
    refresh_token TEXT      NOT NULL,
    slack_user_id TEXT      NOT NULL REFERENCES slack_auths (user_id),
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);