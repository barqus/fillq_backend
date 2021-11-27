CREATE TABLE IF NOT EXISTS participants(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    birth_day DATE,
    description TEXT,
    nickname VARCHAR(100) NOT NULL,
    summoner_name VARCHAR(100) NOT NULL,
    twitch_channel VARCHAR(100) NOT NULL,
    instagram VARCHAR(100) NOT NULL,
    twitter VARCHAR(100) NOT NULL,
    youtube VARCHAR(100) NOT NULL
);