package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateGuild(guild *types.Guild) error {
	return s.Insert("Guild", map[string]any{
		"name":         guild.Name,
		"abbreviation": guild.Abbreviation,
		"description":  guild.Description,
	})
}

func (s *SQLiteStorage) GetGuild() ([]*types.Guild, error) {
	rows, err := s.db.Query("SELECT * FROM Guild")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var guilds []*types.Guild
	for rows.Next() {
		guild := new(types.Guild)
		if err := rows.Scan(&guild.Id, &guild.Name, &guild.Abbreviation, &guild.Description); err != nil {
			return nil, err
		}
		guilds = append(guilds, guild)
	}
	return guilds, nil
}

func (s *SQLiteStorage) GetGuildById(id int) ([]*types.Guild, error) {
	row := s.db.QueryRow("SELECT * FROM Guild WHERE id = ?", id)
	guild := new(types.Guild)
	if err := row.Scan(&guild.Id, &guild.Name, &guild.Abbreviation, &guild.Description); err != nil {
		return nil, err
	}
	return []*types.Guild{guild}, nil
}

func (s *SQLiteStorage) UpdateGuild(guild *types.Guild, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(guild)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Update("Guild", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteGuild(id int) error {
	return s.Delete("Guild", map[string]any{"id": id})
}

/* ------------------------------ Role ------------------------------ */

func (s *SQLiteStorage) CreateGuildRole(guildRole *types.GuildRole) error {
	values := map[string]any{
		"name":        guildRole.Name,
		"description": guildRole.Description,
	}
	return s.Insert("GuildRole", values)
}

func (s *SQLiteStorage) GetGuildRole() ([]*types.GuildRole, error) {
	rows, err := s.db.Query("SELECT * FROM GuildRole")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var guildRoles []*types.GuildRole
	for rows.Next() {
		guildRole := new(types.GuildRole)
		if err := rows.Scan(&guildRole.Id, &guildRole.Name, &guildRole.Description); err != nil {
			return nil, err
		}
		guildRoles = append(guildRoles, guildRole)
	}
	return guildRoles, nil
}

func (s *SQLiteStorage) GetGuildRoleById(id int) ([]*types.GuildRole, error) {
	row := s.db.QueryRow("SELECT * FROM GuildRole WHERE id = ?", id)
	guildRole := new(types.GuildRole)
	if err := row.Scan(&guildRole.Id, &guildRole.Name, &guildRole.Description); err != nil {
		return nil, err
	}
	return []*types.GuildRole{guildRole}, nil
}

func (s *SQLiteStorage) UpdateGuildRole(guildRole *types.GuildRole, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(guildRole)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Update("GuildRole", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteGuildRole(id int) error {
	return s.Delete("GuildRole", map[string]any{"id": id})
}

/* ------------------------------ Member ------------------------------ */

func (s *SQLiteStorage) CreateGuildMember(guildUser *types.GuildMember) error {
	return s.Insert("GuildUser", map[string]any{
		"user_id":  guildUser.UserId,
		"guild_id": guildUser.GuildId,
		"role_id":  guildUser.RoleId,
	})
}

func (s *SQLiteStorage) GetGuildMemberByGuildId(guildId int) ([]*types.GuildMember, error) {
	rows, err := s.db.Query("SELECT * FROM GuildUser WHERE guild_id = ?", guildId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var guildUsers []*types.GuildMember
	for rows.Next() {
		guildUser := new(types.GuildMember)
		if err := rows.Scan(&guildUser.UserId, &guildUser.GuildId, &guildUser.RoleId); err != nil {
			return nil, err
		}
		guildUsers = append(guildUsers, guildUser)
	}
	return guildUsers, nil
}

func (s *SQLiteStorage) GetGuildMemberByUserId(userId int) ([]*types.GuildMember, error) {
	rows, err := s.db.Query("SELECT * FROM GuildUser WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var guildUsers []*types.GuildMember
	for rows.Next() {
		guildUser := new(types.GuildMember)
		if err := rows.Scan(&guildUser.UserId, &guildUser.GuildId, &guildUser.RoleId); err != nil {
			return nil, err
		}
		guildUsers = append(guildUsers, guildUser)
	}
	return guildUsers, nil
}

func (s *SQLiteStorage) DeleteGuildMember(userId int) error {
	return s.Delete("GuildUser", map[string]any{"user_id": userId})
}
