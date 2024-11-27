package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/RakanMyHusbando/shogun/types"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{
		db: db,
	}, nil
}

func RunSQLiteStorage(storage *SQLiteStorage, schemaFile string) error {
	// read sql file with create table statements
	byteContent, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	// execute sql create table statements
	queries := strings.Split(string(byteContent), ";")
	for i := range queries {
		query := strings.TrimSpace(queries[i])
		_, err := storage.db.Exec(query)
		if err != nil {
			return fmt.Errorf("%v [Table: %v]", err.Error(), strings.Split(query, " ")[5])
		}
	}
	return nil
}

type Storage interface {
	// User

	CreateUser(user *types.User) error
	GetUser() ([]*types.User, error)
	GetUserById(id int) ([]*types.User, error)
	UpdateUser(user *types.User, id int) error
	DeleteUser(id int) error

	// LeagueOfLegends

	CreateLeagueOfLeagends(lol *types.LeagueOfLegends, userId int) error
	GetLeagueOfLegendsByUserId(userId int) (*types.LeagueOfLegends, error)
	UpdateLeagueOfLegends(lol *types.LeagueOfLegends, userId int) error
	DeleteLeagueOfLegends(userId int) error

	// GameAccount

	CreateGameAccount(account *types.GameAccount) error
	GetGameAccountBy(userId int, game string) ([]string, error)
	UpdateGameAccount(userId int, oldName string, newName string) error
	DeleteGameAccount(userId int, name string) error

	// Guild

	CreateGuild(guild *types.Guild) error
	GetGuild() ([]*types.Guild, error)
	GetGuildById(id int) ([]*types.Guild, error)
	UpdateGuild(guild *types.Guild, id int) error
	DeleteGuild(id int) error

	// Guild.Role

	CreateGuildRole(guildRole *types.GuildRole) error
	GetGuildRole() ([]*types.GuildRole, error)
	GetGuildRoleById(id int) ([]*types.GuildRole, error)
	UpdateGuildRole(guildRole *types.GuildRole, id int) error
	DeleteGuildRole(id int) error

	// Guild.Member

	CreateGuildMember(guildUser *types.GuildMember) error
	GetGuildMemberByGuildId(guildId int) ([]*types.GuildMember, error)
	GetGuildMemberByUserId(userId int) ([]*types.GuildMember, error)
	DeleteGuildMember(userId int) error

	// Team

	CreateTeam(team *types.Team) error
	GetTeam() ([]*types.Team, error)
	GetTeamById(id int) ([]*types.Team, error)
	UpdateTeam(team *types.Team, id int) error
	DeleteTeam(id int) error

	// Team.Role

	CreateTeamRole(role *types.TeamRole) error
	GetTeamRole() ([]*types.TeamRole, error)
	GetTeamRoleByUserId(id int) ([]*types.TeamRole, error)
	UpdateTeamRole(role *types.TeamRole, id int) error
	DeletTeamRole(id int) error

	// Team.Member

	CreateTeamMember(member *types.TeamMember) error
	GetTeamMemberByUserId(userId int) ([]*types.TeamMember, error)
	GetTeamMemberByTeamId(teamId int) ([]*types.TeamMember, error)
	DeleteTeamMember(id int) error

	// Discord

	CreateDiscord(discord *types.Discord) error
	GetDiscord() ([]*types.Discord, error)
	GetDiscordById(id int) ([]*types.Discord, error)
	UpdateDiscord(discord *types.Discord, id int) error
	DeleteDiscord(id int) error

	// Discord.Role

	CreateDiscordRole(role *types.DiscordRole) error
	GetDiscordRole() ([]*types.DiscordRole, error)
	GetDiscordRoleById(id int) ([]*types.DiscordRole, error)
	UpdateDiscordRole(role *types.DiscordRole, id int) error
	DeleteDiscordRole(id int) error

	// Discord.Member

	CreateDiscordMember(member *types.DiscordMember) error
	GetDiscordMemberByServerId(serverId int) ([]*types.DiscordMember, error)
	GetDiscordMemberByUserId(userId int) ([]*types.DiscordMember, error)
	DeleteDiscordMember(userId int) error
}

type SQLiteStorage struct {
	db *sql.DB
}

/* ------------------------------ SQLite reqs ------------------------------ */

func (s *SQLiteStorage) Insert(table string, insertValues map[string]any) error {
	var keys, values string
	first := true
	for key, value := range insertValues {
		if value != nil && value != "" && value != 0 {
			if !first {
				keys += ", "
				values += ", "
			} else {
				first = false
			}
			if reflect.TypeOf(value).String() == "string" {
				value = fmt.Sprintf("'%s'", value)
			}
			keys += key
			values = fmt.Sprintln(values, value)
		}
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, keys, values)
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[storage.main] inserted element in table ", table)
	return nil
}

func (s *SQLiteStorage) Update(table string, set map[string]any, where map[string]any) error {
	query := fmt.Sprintf(`UPDATE %s SET `, table)
	first := true
	for key, value := range set {
		if value != nil && value != "" && value != 0 {
			if !first {
				query += ", "
			} else {
				first = false
			}
			query += fmt.Sprintf(`%s = `, key)
			if reflect.TypeOf(value).String() == "string" {
				query += fmt.Sprintf("'%s'", value)
			} else {
				query = fmt.Sprintln(query, value)
			}
		}
	}
	query += " WHERE "
	first = true
	for key, value := range where {
		if !first {
			query += " AND "
		} else {
			first = false
		}
		query += fmt.Sprintf(`%s = `, key)
		if reflect.TypeOf(value).String() == "string" {
			query += fmt.Sprintf("'%s'", value)
		} else {
			query = fmt.Sprintln(query, value)
		}
	}
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[storage.main] element updated from table: ", table)
	return nil
}

func (s *SQLiteStorage) Delete(table string, where map[string]any) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE `, table)
	first := true
	for key, value := range where {
		if value != nil {
			if !first {
				query += " AND "
			} else {
				first = false
			}
			query += fmt.Sprintf(`%s = `, key)
			if reflect.TypeOf(value).String() == "string" {
				query += fmt.Sprintf("'%s'", value)
			} else {
				query = fmt.Sprintln(query, value)
			}
		}
	}
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[storage.main] delete element from ", table)
	return nil
}
