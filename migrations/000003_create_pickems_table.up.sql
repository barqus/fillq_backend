CREATE TABLE IF NOT EXISTS pickems
(
    user_id        INTEGER NOT NULL,
    participant_id INTEGER NOT NULL,
    CHECK (participant_id > -1 AND participant_id < 13),
    position       INTEGER NOT NULL,
    CHECK (participant_id > -1 AND participant_id < 13),
    PRIMARY KEY (user_id, participant_id, position),
    UNIQUE (user_id, position),
    CONSTRAINT fk_participant
        FOREIGN KEY(participant_id)
            REFERENCES participants(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);