CREATE TABLE IF NOT EXISTS twitch
(
    twitch_id      VARCHAR(100) PRIMARY KEY,
    display_name   VARCHAR(100),
    game_name      VARCHAR(100),
    is_live        VARCHAR(100),
    title          VARCHAR(100),
    started_at     VARCHAR(100),
    participant_id INTEGER NOT NULL,
    CONSTRAINT fk_participant
        FOREIGN KEY (participant_id)
            REFERENCES participants (id)
);