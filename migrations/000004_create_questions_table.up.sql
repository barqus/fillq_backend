CREATE TABLE IF NOT EXISTS questions
(
    id       SERIAL PRIMARY KEY,
    participant_id  INTEGER NOT NULL,
    question TEXT,
    answer   TEXT,
    CONSTRAINT fk_user
        FOREIGN KEY(participant_id)
            REFERENCES participants(id)
);