CREATE TABLE slack_auths
(
    user_id      TEXT      PRIMARY KEY ,
    team_id      TEXT      NOT NULL,
    view_id      TEXT,
    access_token TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);