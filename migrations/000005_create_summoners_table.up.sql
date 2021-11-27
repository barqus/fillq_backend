CREATE TABLE IF NOT EXISTS summoners
(
    puuid          VARCHAR(100) PRIMARY KEY,
    summoner_name  VARCHAR(100),
    tier           VARCHAR(100),
    rank           VARCHAR(100),
    points         VARCHAR(100),
    wins           INT,
    losses         INT,
    participant_id INTEGER NOT NULL,
    CONSTRAINT fk_participant
        FOREIGN KEY(participant_id)
            REFERENCES participants(id)
);
--
--
-- SummonerName: leagueItem.SummonerName,
-- Tier:         &leagueItem.Tier,
-- Rank:         &leagueItem.Rank,
-- Points:       &leagueItem.LeaguePoints,
-- Wins:         leagueItem.Wins,
-- Losses:       leagueItem.Losses,
-- ParticipantID:       item.Id,