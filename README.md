# ORGA
"ORGA" is a Go-based application that integrates with the Riot API to manage user data, game accounts, teams, and guilds (and guild discord server with roles). It uses SQLite for data storage and provides an API server for interaction.

## Status 

- in progress
- not all endpoints are tested 

## Features

- User management (game-accounts, game-stats)
- Team management
- Guild management

## Requirements

- Go 1.16+
- SQLite3
- Riot API Key

## Setup

1. Clone the repository:
```sh
git clone https://github.com/RakanMyHusbando/orga.git
cd orga
```

2. Create a `.env` file in the root directory with the following content:
```env
API_KEY=your_riot_api_key
DB_FILE=path_to_your_sqlite_db_file
PORT=your_server_port
```

3. Install dependencies:
```sh
go mod tidy
```

## Usage

Run the application:
```sh
go run github.com/RakanMyHusbando/orga
```

## API Endpoints

### User

- `POST /user` - Create a new user.
- `GET /user` - Retrieve all users.
- `GET /user/{userId}` - Retrieve a user by ID.
- `PATCH /user/{userId}` - Update a user by ID.
- `DELETE /user/{userId}` - Delete a user by ID.

### League of Legends

- `POST /user/{userId}/league_of_legends` - Create League of Legends data for a user.
- `PATCH /user/{userId}/league_of_legends` - Partially update League of Legends data for a user.
- `DELETE /user/{userId}/league_of_legends` - Delete League of Legends data for a user.

### Game Account

- `POST /user/{userId}/game_account` - Create a game account for a user.
- `DELETE /user/{userId}/game_account/{name}` - Delete a game account for a user.

### Guild

- `POST /guild` - Create a new guild.
- `GET /guild` - Retrieve all guilds.
- `GET /guild/{guildId}` - Retrieve a guild by ID.
- `PATCH /guild/{guildId}` - Partailly update a guild by ID.
- `DELETE /guild/{guildId}` - Delete a guild by ID.

### Guild Role

- `POST /guild_role` - Add a new guild role.
- `GET /guild_role` - Retrieve all guild roles.
- `PATCH /guild_role/{guildRoleId}` - Partially update guild role data.
- `DELETE /guild_role/{guildRoleId}` - Delete a guild role by ID.

### Guild Members

- `POST /guild_member` - Add a member to a guild.
- `DELETE /guild_member/{userId}` - Remove a member from a guild by user ID.

### Team

- `POST /team` - Create a new team.
- `GET /team` - Retrieve all teams.
- `GET /team/{teamId}` - Retrieve a team by ID.
- `PATCH /team/{teamId}` - Partially update a team by ID.
- `DELETE /team/{teamId}` - Delete a team by ID.

### Team Role

- `POST /team_role` - Add a new team role.
- `GET /team_role` - Retrieve all team roles.
- `PATCH /team_role/{teamRoleId}` - Partially update team role data.
- `DELETE /team_role/{teamRoleId}` - Delete a guild team by ID.

### Team Members

- `POST /team_member` - Add a member to a team.
- `DELETE /team_member/{userId}` - Remove a member from a team by user ID.

### Discord Server

- `POST /discord` - Add a new discord server.
- `GET /discord` - Retrieve all discord server.
- `GET /discord/{discordId}` - Retrieve a disocrd server by ID.
- `PATCH /discord/{discordId}` - Partially update a discord sever by ID.
- `DELETE /discord/{discordId}` - Delete a discord server by ID.

### Discord Role

- `POST /discord_role` - Add a new discord role.
- `GET /discord_role` - Retrieve all discord roles.
- `PATCH /discord_role/{discordRoleId}` - Partially update discord role data.
- `DELETE /discord_role/{discordRoleId}` - Delete a discord role by ID.

### Discord Members

- `POST /team_member` - Add a member to a discord server.
- `DELETE /team_member/{userId}` - Remove a member from a discord server by user ID.
