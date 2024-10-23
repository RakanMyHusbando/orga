CREATE TABLE IF NOT EXISTS User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    discord_id TEXT
);
CREATE TABLE IF NOT EXISTS AccountLeagueOfLegends (
    user_id INTEGER NOT NULL,
    name TEXT,
    FOREIGN KEY (user_id) REFERENCES User(id)
);
CREATE TABLE IF NOT EXISTS UserLeagueOfLegends (
    user_id INTEGER NOT NULL UNIQUE,
    main_role TEXT,
    second_role TEXT,
    champ_0 TEXT,
    champ_1 TEXT,
    champ_2 TEXT,
    FOREIGN KEY (user_id) REFERENCES User(id)
);