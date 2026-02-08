-- Seed demo data for portfolio demonstration.
-- Run with: psql $DATABASE_URL -f migrations/seed.sql

BEGIN;

-- Clean previous seed data (idempotent)
DELETE FROM matches WHERE match_id LIKE 'demo-%';
DELETE FROM player_stats WHERE player_id IN (
    SELECT id FROM players WHERE id IN (
        'a0000001-0000-0000-0000-000000000001',
        'a0000001-0000-0000-0000-000000000002',
        'a0000001-0000-0000-0000-000000000003',
        'a0000001-0000-0000-0000-000000000004',
        'a0000001-0000-0000-0000-000000000005'
    )
);
DELETE FROM players WHERE id IN (
    'a0000001-0000-0000-0000-000000000001',
    'a0000001-0000-0000-0000-000000000002',
    'a0000001-0000-0000-0000-000000000003',
    'a0000001-0000-0000-0000-000000000004',
    'a0000001-0000-0000-0000-000000000005'
);

--------------------------------------------------------------------------------
-- 1. Demo players
--------------------------------------------------------------------------------
INSERT INTO players (id, platform, gamertag, last_fetched_at) VALUES
('a0000001-0000-0000-0000-000000000001', 'xbl',    'TacticalNuke99',    NOW()),
('a0000001-0000-0000-0000-000000000002', 'psn',    'ShadowSniper_TTV',  NOW()),
('a0000001-0000-0000-0000-000000000003', 'xbl',    'CasualCarl',        NOW()),
('a0000001-0000-0000-0000-000000000004', 'uno',    'GhostRecon42',      NOW()),
('a0000001-0000-0000-0000-000000000005', 'battle', 'NoobMaster69',      NOW());

--------------------------------------------------------------------------------
-- 2. Lifetime stats snapshots (one per player, mode = 'wz')
--------------------------------------------------------------------------------

-- TacticalNuke99 — Elite (K/D 2.51)
INSERT INTO player_stats (player_id, mode, stats_data) VALUES
('a0000001-0000-0000-0000-000000000001', 'wz', '{
  "platform": "xbl",
  "gamertag": "TacticalNuke99",
  "level": 155,
  "prestige": 3,
  "kills": 15847,
  "deaths": 6314,
  "kdRatio": 2.51,
  "wins": 312,
  "losses": 1188,
  "winPct": 0.208,
  "scorePerMin": 342.5,
  "headshots": 4123,
  "timePlayed": 864000,
  "matchesPlayed": 1500,
  "topFive": 487,
  "topTen": 612,
  "topTwentyFive": 823,
  "assists": 2156,
  "damageDone": 4567890,
  "modeBreakdown": {
    "br_brquads": {
      "kills": 6234, "deaths": 2487, "kdRatio": 2.51,
      "wins": 156, "losses": 444, "matchesPlayed": 600,
      "scorePerMin": 365.0, "timePlayed": 345600,
      "topFive": 234, "topTen": 298, "topTwentyFive": 378
    },
    "br_brtrios": {
      "kills": 4321, "deaths": 1721, "kdRatio": 2.51,
      "wins": 89, "losses": 311, "matchesPlayed": 400,
      "scorePerMin": 330.0, "timePlayed": 230400,
      "topFive": 134, "topTen": 167, "topTwentyFive": 234
    },
    "br_rebirth_rbrthquads": {
      "kills": 3456, "deaths": 1378, "kdRatio": 2.51,
      "wins": 67, "losses": 233, "matchesPlayed": 300,
      "scorePerMin": 380.0, "timePlayed": 172800,
      "topFive": 89, "topTen": 112, "topTwentyFive": 156
    },
    "br_brsolo": {
      "kills": 1836, "deaths": 728, "kdRatio": 2.52,
      "wins": 0, "losses": 200, "matchesPlayed": 200,
      "scorePerMin": 295.0, "timePlayed": 115200,
      "topFive": 30, "topTen": 35, "topTwentyFive": 55
    }
  }
}'::jsonb);

-- ShadowSniper_TTV — Good (K/D 1.68)
INSERT INTO player_stats (player_id, mode, stats_data) VALUES
('a0000001-0000-0000-0000-000000000002', 'wz', '{
  "platform": "psn",
  "gamertag": "ShadowSniper_TTV",
  "level": 142,
  "prestige": 2,
  "kills": 9234,
  "deaths": 5497,
  "kdRatio": 1.68,
  "wins": 156,
  "losses": 844,
  "winPct": 0.156,
  "scorePerMin": 278.3,
  "headshots": 2987,
  "timePlayed": 612000,
  "matchesPlayed": 1000,
  "topFive": 298,
  "topTen": 412,
  "topTwentyFive": 567,
  "assists": 1543,
  "damageDone": 2876543,
  "modeBreakdown": {
    "br_brquads": {
      "kills": 3890, "deaths": 2315, "kdRatio": 1.68,
      "wins": 78, "losses": 322, "matchesPlayed": 400,
      "scorePerMin": 290.0, "timePlayed": 230400,
      "topFive": 134, "topTen": 189, "topTwentyFive": 256
    },
    "br_brtrios": {
      "kills": 2567, "deaths": 1528, "kdRatio": 1.68,
      "wins": 45, "losses": 255, "matchesPlayed": 300,
      "scorePerMin": 265.0, "timePlayed": 172800,
      "topFive": 89, "topTen": 123, "topTwentyFive": 178
    },
    "br_rebirth_rbrthtrios": {
      "kills": 2777, "deaths": 1654, "kdRatio": 1.68,
      "wins": 33, "losses": 267, "matchesPlayed": 300,
      "scorePerMin": 310.0, "timePlayed": 208800,
      "topFive": 75, "topTen": 100, "topTwentyFive": 133
    }
  }
}'::jsonb);

-- CasualCarl — Average (K/D 0.94)
INSERT INTO player_stats (player_id, mode, stats_data) VALUES
('a0000001-0000-0000-0000-000000000003', 'wz', '{
  "platform": "xbl",
  "gamertag": "CasualCarl",
  "level": 87,
  "prestige": 0,
  "kills": 4567,
  "deaths": 4858,
  "kdRatio": 0.94,
  "wins": 43,
  "losses": 757,
  "winPct": 0.054,
  "scorePerMin": 156.2,
  "headshots": 1023,
  "timePlayed": 432000,
  "matchesPlayed": 800,
  "topFive": 112,
  "topTen": 198,
  "topTwentyFive": 345,
  "assists": 876,
  "damageDone": 1234567,
  "modeBreakdown": {
    "br_brquads": {
      "kills": 2345, "deaths": 2494, "kdRatio": 0.94,
      "wins": 32, "losses": 368, "matchesPlayed": 400,
      "scorePerMin": 165.0, "timePlayed": 230400,
      "topFive": 67, "topTen": 112, "topTwentyFive": 189
    },
    "br_rebirth_rbrthquads": {
      "kills": 1567, "deaths": 1667, "kdRatio": 0.94,
      "wins": 11, "losses": 289, "matchesPlayed": 300,
      "scorePerMin": 178.0, "timePlayed": 144000,
      "topFive": 34, "topTen": 56, "topTwentyFive": 112
    },
    "br_dmz": {
      "kills": 655, "deaths": 697, "kdRatio": 0.94,
      "wins": 0, "losses": 100, "matchesPlayed": 100,
      "scorePerMin": 120.0, "timePlayed": 57600,
      "topFive": 11, "topTen": 30, "topTwentyFive": 44
    }
  }
}'::jsonb);

-- GhostRecon42 — Above Average (K/D 1.31)
INSERT INTO player_stats (player_id, mode, stats_data) VALUES
('a0000001-0000-0000-0000-000000000004', 'wz', '{
  "platform": "uno",
  "gamertag": "GhostRecon42",
  "level": 121,
  "prestige": 1,
  "kills": 7123,
  "deaths": 5437,
  "kdRatio": 1.31,
  "wins": 98,
  "losses": 802,
  "winPct": 0.109,
  "scorePerMin": 223.7,
  "headshots": 1876,
  "timePlayed": 540000,
  "matchesPlayed": 900,
  "topFive": 198,
  "topTen": 312,
  "topTwentyFive": 467,
  "assists": 1234,
  "damageDone": 2345678,
  "modeBreakdown": {
    "br_brquads": {
      "kills": 2890, "deaths": 2206, "kdRatio": 1.31,
      "wins": 45, "losses": 305, "matchesPlayed": 350,
      "scorePerMin": 235.0, "timePlayed": 201600,
      "topFive": 87, "topTen": 134, "topTwentyFive": 198
    },
    "br_brtrios": {
      "kills": 2134, "deaths": 1629, "kdRatio": 1.31,
      "wins": 33, "losses": 217, "matchesPlayed": 250,
      "scorePerMin": 215.0, "timePlayed": 144000,
      "topFive": 56, "topTen": 89, "topTwentyFive": 134
    },
    "br_brduos": {
      "kills": 1345, "deaths": 1027, "kdRatio": 1.31,
      "wins": 15, "losses": 185, "matchesPlayed": 200,
      "scorePerMin": 210.0, "timePlayed": 115200,
      "topFive": 34, "topTen": 56, "topTwentyFive": 89
    },
    "br_rebirth_rbrthtrios": {
      "kills": 754, "deaths": 575, "kdRatio": 1.31,
      "wins": 5, "losses": 95, "matchesPlayed": 100,
      "scorePerMin": 250.0, "timePlayed": 79200,
      "topFive": 21, "topTen": 33, "topTwentyFive": 46
    }
  }
}'::jsonb);

-- NoobMaster69 — Below Average (K/D 0.62)
INSERT INTO player_stats (player_id, mode, stats_data) VALUES
('a0000001-0000-0000-0000-000000000005', 'wz', '{
  "platform": "battle",
  "gamertag": "NoobMaster69",
  "level": 45,
  "prestige": 0,
  "kills": 2345,
  "deaths": 3782,
  "kdRatio": 0.62,
  "wins": 12,
  "losses": 488,
  "winPct": 0.024,
  "scorePerMin": 98.4,
  "headshots": 456,
  "timePlayed": 288000,
  "matchesPlayed": 500,
  "topFive": 45,
  "topTen": 89,
  "topTwentyFive": 178,
  "assists": 543,
  "damageDone": 678901,
  "modeBreakdown": {
    "br_brquads": {
      "kills": 1234, "deaths": 1990, "kdRatio": 0.62,
      "wins": 10, "losses": 290, "matchesPlayed": 300,
      "scorePerMin": 105.0, "timePlayed": 172800,
      "topFive": 28, "topTen": 56, "topTwentyFive": 112
    },
    "br_rebirth_rbrthquads": {
      "kills": 789, "deaths": 1273, "kdRatio": 0.62,
      "wins": 2, "losses": 148, "matchesPlayed": 150,
      "scorePerMin": 115.0, "timePlayed": 86400,
      "topFive": 12, "topTen": 23, "topTwentyFive": 45
    },
    "br_dmz": {
      "kills": 322, "deaths": 519, "kdRatio": 0.62,
      "wins": 0, "losses": 50, "matchesPlayed": 50,
      "scorePerMin": 78.0, "timePlayed": 28800,
      "topFive": 5, "topTen": 10, "topTwentyFive": 21
    }
  }
}'::jsonb);

--------------------------------------------------------------------------------
-- 3. Match history (25 matches per player, generated via PL/pgSQL)
--------------------------------------------------------------------------------
DO $$
DECLARE
    p RECORD;
    i INT;
    skill FLOAT;
    k INT;
    d INT;
    pl INT;
    dmg_dealt INT;
    dmg_taken INT;
    modes TEXT[] := ARRAY[
        'br_brquads', 'br_brtrios', 'br_brduos',
        'br_brsolo', 'br_rebirth_rbrthquads', 'br_rebirth_rbrthtrios'
    ];
    maps TEXT[] := ARRAY[
        'Verdansk', 'Caldera', 'Rebirth Island',
        'Fortune''s Keep', 'Ashika Island', 'Al Mazrah'
    ];
    gulag TEXT;
BEGIN
    FOR p IN
        SELECT id, gamertag,
            CASE gamertag
                WHEN 'TacticalNuke99'   THEN 2.5
                WHEN 'ShadowSniper_TTV' THEN 1.7
                WHEN 'CasualCarl'       THEN 0.95
                WHEN 'GhostRecon42'     THEN 1.3
                WHEN 'NoobMaster69'     THEN 0.6
                ELSE 1.0
            END AS skill
        FROM players
        WHERE id IN (
            'a0000001-0000-0000-0000-000000000001',
            'a0000001-0000-0000-0000-000000000002',
            'a0000001-0000-0000-0000-000000000003',
            'a0000001-0000-0000-0000-000000000004',
            'a0000001-0000-0000-0000-000000000005'
        )
    LOOP
        skill := p.skill;

        FOR i IN 1..25 LOOP
            -- Kills: varies with skill and match index
            k := greatest(0, round(skill * 3.5 + 2.0 * sin(i * 1.7))::INT);
            d := greatest(1, round(k / skill + cos(i * 2.3))::INT);

            -- Placement: better players get more wins/top placements
            IF mod(i, greatest(1, round(10.0 / skill)::INT)) = 0 THEN
                pl := 1;
            ELSIF mod(i, 3) = 0 THEN
                pl := 2 + mod(i, 8);
            ELSE
                pl := 10 + mod(i * 7, 140);
            END IF;

            dmg_dealt := k * (700 + mod(i, 5) * 150);
            dmg_taken := d * (500 + mod(i, 4) * 100);

            -- Gulag: cycle win/loss/none
            IF mod(i, 3) = 0 THEN gulag := 'win';
            ELSIF mod(i, 3) = 1 THEN gulag := 'loss';
            ELSE gulag := '';
            END IF;

            INSERT INTO matches (
                match_id, player_id, mode, map_name, placement,
                kills, deaths, damage_dealt, damage_taken,
                gulag_result, match_time, raw_data
            ) VALUES (
                'demo-' || left(p.gamertag, 12) || '-' || i,
                p.id,
                modes[1 + mod(i - 1, 6)],
                maps[1 + mod(i + ascii(left(p.gamertag, 1)), 6)],
                pl,
                k,
                d,
                dmg_dealt,
                dmg_taken,
                gulag,
                NOW() - make_interval(hours => i * 6 + mod(ascii(left(p.gamertag, 1)), 12)),
                '{}'::jsonb
            )
            ON CONFLICT (match_id, player_id) DO NOTHING;
        END LOOP;
    END LOOP;
END $$;

COMMIT;
