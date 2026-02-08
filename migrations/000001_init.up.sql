CREATE TABLE players (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    platform        VARCHAR(10) NOT NULL,
    gamertag        VARCHAR(100) NOT NULL,
    activision_id   VARCHAR(100),
    last_fetched_at TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(platform, gamertag)
);

CREATE TABLE squads (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE squad_members (
    squad_id    UUID NOT NULL REFERENCES squads(id) ON DELETE CASCADE,
    player_id   UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    added_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (squad_id, player_id)
);

CREATE TABLE matches (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id        VARCHAR(100) NOT NULL UNIQUE,
    player_id       UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    mode            VARCHAR(50),
    map_name        VARCHAR(100),
    placement       INT,
    kills           INT,
    deaths          INT,
    damage_dealt    INT,
    damage_taken    INT,
    gulag_result    VARCHAR(10),
    match_time      TIMESTAMPTZ,
    raw_data        JSONB,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE player_stats (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id       UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    mode            VARCHAR(50) NOT NULL,
    stats_data      JSONB NOT NULL,
    fetched_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_matches_player_id ON matches(player_id);
CREATE INDEX idx_matches_match_time ON matches(match_time);
CREATE INDEX idx_player_stats_player_mode ON player_stats(player_id, mode);
