CREATE TABLE events
(
    subscription_id TEXT,
    type            TEXT,
    twitch_user_id  TEXT,
    slack_user_id   TEXT,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);