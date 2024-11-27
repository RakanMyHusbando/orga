CREATE TABLE IF NOT EXISTS User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    discord_id TEXT UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS GameAccount (
    user_id INTEGER NOT NULL,
    name TEXT UNIQUE NOT NULL,
    game TEXT NOT NULL,
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
CREATE TABLE IF NOT EXISTS Guild (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    abbreviation TEXT NOT NULL,
    description TEXT
);
CREATE TABLE IF NOT EXISTS GuildRole (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);
CREATE TABLE IF NOT EXISTS GuildUser (
    user_id INTEGER NOT NULL,
    guild_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (guild_id) REFERENCES Guild(id),
    FOREIGN KEY (role_id) REFERENCES GuildRole(id)
);
CREATE TABLE IF NOT EXISTS Team (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    guild_id INTEGER,
    name TEXT NOT NULL,
    abbreviation TEXT NOT NULL,
    FOREIGN KEY (guild_id) REFERENCES Guild(id)
);
CREATE TABLE IF NOT EXISTS TeamRole (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);
CREATE TABLE IF NOT EXISTS TeamUser (
    user_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (team_id) REFERENCES Team(id),
    FOREIGN KEY (role_id) REFERENCES TeamRole(id)
);
CREATE TABLE IF NOT EXISTS DiscordServer (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    discord_id TEXT UNIQUE,
    name TEXT,
    description TEXT,
    guild_id INTEGER,
    team_id INTEGER,
    FOREIGN KEY (guild_id) REFERENCES Guild(id)
    FOREIGN KEY (team_id) REFERENCES Team(id)
);
CREATE TABLE IF NOT EXISTS DiscordRole (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  
    name TEXT UNIQUE NOT NULL,
    description TEXT
    -- TODO: add some role privileges (type: boolean)
);
CREATE TABLE IF NOT EXISTS DiscordUser (
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    server_id INTEGER NOT NULL,
    FOREIGN KEY (role_id) REFERENCES DiscordRole(id),
    FOREIGN KEY (server_id) REFERENCES DiscordServer(id)
);