package storage

import (
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

/* ------------------------------ guild ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuild(Guild *types.ReqGuild) error {
	prep, err := s.db.Prepare(`INSERT INTO Guild (name, abbreviation, description) VALUES (?,?,?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		Guild.Name,
		Guild.Abbreviation,
		Guild.Description,
	); err != nil {
		return err
	}

	prep.Close()
	log.Println("Storage: successfully created guild")

	return nil
}

// GET
func (s *SQLiteStorage) GetGuild() ([]*types.ResGuild, error) {
	rows, err := s.db.Query(`SELECT * FROM Guild`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guilds := []*types.ResGuild{}
	for rows.Next() {
		guild := new(types.ResGuild)
		if err := rows.Scan(
			&guild.Id,
			&guild.Name,
			&guild.Abbreviation,
			&guild.Description,
		); err != nil {
			return nil, err
		}
		roleUserList, err := s.GetGuildMemberMapByGuildId(guild.Id)
		if err != nil {
			return nil, err
		}
		guild.Member = roleUserList

		guilds = append(guilds, guild)
	}

	log.Println("Storage: successfully get guilds")

	return guilds, nil
}

// GET
func (s *SQLiteStorage) GetGuildById(id int) (*types.ResGuild, error) {
	row := s.db.QueryRow(`SELECT * FROM Guild WHERE id = ?`, id)

	guild := new(types.ResGuild)
	if err := row.Scan(
		&guild.Id,
		&guild.Name,
		&guild.Abbreviation,
		&guild.Description,
	); err != nil {
		return nil, err
	}

	roleUserList, err := s.GetGuildMemberMapByGuildId(guild.Id)
	if err != nil {
		return nil, err
	}
	guild.Member = roleUserList
	log.Println("Storage: successfully got guild by id")
	return guild, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuild(id int) error {
	prep, err := s.db.Prepare(`DELETE FROM Guild WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(id); err != nil {
		return err
	}

	prep.Close()
	log.Println("Storage: successfully deleted guild")
	return nil
}

// PATCH
func (s *SQLiteStorage) UpdateGuild(guild *types.ResGuild) error {
	prep, err := s.db.Prepare(`
		UPDATE 
			Guild 
		SET 
			name = ?, 
			abbreviation = ?, 
			description = ? 
		WHERE 
			id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		guild.Name,
		guild.Abbreviation,
		guild.Description,
		guild.Id,
	); err != nil {
		return err
	}

	prep.Close()
	log.Println("Storage: successfully updated guild")
	return nil
}

/* ------------------------------ guild role ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuildRole(guildRole *types.ReqGuildRole) error {
	prep, err := s.db.Prepare(`INSERT INTO GuildRole (name, description) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(guildRole.Name, guildRole.Description); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully created guild role")
	return nil
}

// GET
func (s *SQLiteStorage) GetGuildRole() ([]*types.ReqGuildRole, error) {
	rows, err := s.db.Query(`SELECT id, name, description FROM GuildRole`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	guildRoles := []*types.ReqGuildRole{}
	for rows.Next() {
		guildRole := new(types.ReqGuildRole)
		if err := rows.Scan(&guildRole.Id, &guildRole.Name, &guildRole.Description); err != nil {
			return nil, err
		}
		guildRoles = append(guildRoles, guildRole)
	}
	log.Println("Storage: successfully get guild roles")
	return guildRoles, nil
}

// GET
func (s *SQLiteStorage) GetGuildRoleById(id int) (*types.ReqGuildRole, error) {
	row := s.db.QueryRow(`SELECT name, description FROM GuildRole WHERE id = ?`, id)
	guildRole := new(types.ReqGuildRole)
	if err := row.Scan(&guildRole.Name, &guildRole.Description); err != nil {
		return nil, err
	}
	log.Println("Storage: successfully get guild role by id")
	return guildRole, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuildRole(id int) error {
	prep, err := s.db.Prepare(`DELETE FROM GuildRole WHERE id = ?`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(id); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully deleted guild role")
	return nil
}

// PATCH
func (s *SQLiteStorage) UpdateGuildRole(guildRole *types.ReqGuildRole) error {
	prep, err := s.db.Prepare(`
		UPDATE GuildRole SET name = ?, description = ? WHERE id = ?
	`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(
		guildRole.Name, guildRole.Description, guildRole.Id,
	); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully updated guild role")
	return nil
}

/* ------------------------------ guild member ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuildMember(guildUser *types.ReqGuildMember) error {
	prep, err := s.db.Prepare(`
		INSERT INTO GuildUser (user_id, guild_id, role_id) VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(
		guildUser.UserId, guildUser.GuildId, guildUser.RoleId,
	); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully created guild member")
	return nil
}

// GET
func (s *SQLiteStorage) GetGuildMemberByGuildId(guildId int) ([]*types.ReqGuildMember, error) {
	rows, err := s.db.Query(
		`SELECT user_id, role_id FROM GuildUser WHERE guild_id = ?`, guildId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	guildUsers := []*types.ReqGuildMember{}
	for rows.Next() {
		guildUser := new(types.ReqGuildMember)
		if err := rows.Scan(&guildUser.UserId, &guildUser.RoleId); err != nil {
			return nil, err
		}
		guildUsers = append(guildUsers, guildUser)
	}
	log.Println("Storage: successfully get guild members by guild id")
	return guildUsers, nil
}

// GET
func (s *SQLiteStorage) GetGuildMemberMapByGuildId(guildId int) (map[string][]string, error) {
	guildUser, err := s.GetGuildMemberByGuildId(guildId)
	if err != nil {
		return nil, err
	}
	var roleUserMap map[string][]string = make(map[string][]string)
	for _, guildUser := range guildUser {
		role, err := s.GetGuildRoleById(guildUser.RoleId)
		if err != nil {
			return nil, err
		}
		user, err := s.GetUserById(guildUser.UserId)
		if err != nil {
			return nil, err
		}

		if roleUserMap[role.Name] == nil {
			roleUserMap[role.Name] = []string{user.Name}
		} else {
			roleUserMap[role.Name] = append(roleUserMap[role.Name], user.Name)
		}
	}
	log.Println("Storage: successfully get guild member map by guild id")
	return roleUserMap, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuildMember(user_id int) error {
	prep, err := s.db.Prepare(`DELETE FROM GuildUser WHERE user_id = ? `)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(user_id); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully deleted guild member")
	return nil
}

// PATCH
func (s *SQLiteStorage) UpdateGuildMember(guildUser *types.ReqGuildMember) error {
	prep, err := s.db.Prepare(`
		UPDATE GuildUser SET role_id = ?, guild_id = ? WHERE user_id = ?
	`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(
		guildUser.RoleId, guildUser.GuildId, guildUser.UserId,
	); err != nil {
		return err
	}
	prep.Close()
	log.Println("Storage: successfully updated guild member")
	return nil
}
