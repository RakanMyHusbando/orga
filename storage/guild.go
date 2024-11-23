package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateGuild(guild *types.Guild) error {
	var values map[string]any
	bytes, err := json.Marshal(guild)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("Guild", values)
}

func (s *SQLiteStorage) GetGuild() ([]*map[string]any, error) {
	return s.Select("Guild", nil, nil)
}

func (s *SQLiteStorage) GetGuildById(id int) (*map[string]any, error) {
	return s.SelectUnique("Guild", nil, "id", id)
}

func (s *SQLiteStorage) UpdateGuild(guild *types.Guild, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(guild)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Patch("Guild", values, map[string]any{"id": id})
}

/* ------------------------------ Role ------------------------------ */

func (s *SQLiteStorage) CreateGuildRole(guildRole *types.GuildRole) error {
	var values map[string]any
	bytes, err := json.Marshal(guildRole)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("GuildRole", values)
}

func (s *SQLiteStorage) GetGuildRole() ([]*map[string]any, error) {
	return s.Select("GuildRole", nil, nil)
}

func (s *SQLiteStorage) GetGuildRoleById(id int) (*map[string]any, error) {
	return s.SelectUnique("GuildRole", nil, "id", id)
}

func (s *SQLiteStorage) UpdateGuildRole(guildRole *types.GuildRole, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(guildRole)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Patch("GuildRole", values, map[string]any{"id": id})
}

/* ------------------------------ Member ------------------------------ */

func (s *SQLiteStorage) CreateGuildMember(guildUser *types.GuildMember) error {
	var values map[string]any
	bytes, err := json.Marshal(guildUser)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("GuildUser", values)
}

func (s *SQLiteStorage) GetGuildMemberByGuildId(guildId int) ([]*map[string]any, error) {
	return s.Select("GuildUser", nil, map[string]any{"guild_id": guildId})
}

func (s *SQLiteStorage) UpdateGuildMember(guildUser *types.GuildMember, userId int) error {
	var values map[string]any
	bytes, err := json.Marshal(guildUser)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Patch("GuildUser", values, map[string]any{"user_id": userId})
}

func (s *SQLiteStorage) DeleteGuildMember(userId int) error {
	return s.Delete("GuildUser", map[string]any{"user_id": userId})
}
