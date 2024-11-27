package storage

import "github.com/RakanMyHusbando/shogun/types"

func (s *SQLiteStorage) CreateDiscord(discord *types.Discord) error {
	return s.Insert("Server", map[string]any{
		"name": discord.Name,
	})
}

func (s *SQLiteStorage) GetDiscord() ([]*types.Discord, error) {
	rows, err := s.db.Query("SELECT * FROM Server")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var servers []*types.Discord
	for rows.Next() {
		server := new(types.Discord)
		if err := rows.Scan(&server.Id, &server.Name); err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func (s *SQLiteStorage) GetServerById(id int) ([]*types.Discord, error) {
	row := s.db.QueryRow("SELECT * FROM Server WHERE id = ?", id)
	server := new(types.Discord)
	if err := row.Scan(&server.Id, &server.Name); err != nil {
		return nil, err
	}
	return []*types.Discord{server}, nil
}

func (s *SQLiteStorage) GetDiscordById(id int) ([]*types.DiscordRole, error) {
	row := s.db.QueryRow("SELECT * FROM DiscordRole WHERE id = ?", id)
	role := new(types.DiscordRole)
	if err := row.Scan(&role.Id, &role.Name, &role.Description); err != nil {
		return nil, err
	}
	return []*types.DiscordRole{role}, nil
}

func (s *SQLiteStorage) UpdateServer(server *types.Discord, id int) error {
	return s.Update("Server", map[string]any{
		"name": server.Name,
	}, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteServer(id int) error {
	return s.Delete("Server", map[string]any{"id": id})
}

/* ------------------------------ Role ------------------------------ */

func (s *SQLiteStorage) CreateDiscordRole(role *types.DiscordRole) error {
	return s.Insert("DiscordRole", map[string]any{
		"name":        role.Name,
		"description": role.Description,
	})
}

func (s *SQLiteStorage) GetDiscordRole() ([]*types.DiscordRole, error) {
	rows, err := s.db.Query("SELECT * FROM DiscordRole")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*types.DiscordRole
	for rows.Next() {
		role := new(types.DiscordRole)
		if err := rows.Scan(&role.Id, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (s *SQLiteStorage) UpdateDiscordRole(role *types.DiscordRole, id int) error {
	return s.Update("DiscordRole", map[string]any{
		"name":        role.Name,
		"description": role.Description,
	}, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteDiscordRole(id int) error {
	return s.Delete("DiscordRole", map[string]any{"id": id})
}

/* ------------------------------ Member ------------------------------ */

func (s *SQLiteStorage) CreateDiscordMember(member *types.DiscordMember) error {
	return s.Insert("DiscordMember", map[string]any{
		"user_id":   member.UserId,
		"role_id":   member.RoleId,
		"server_id": member.ServerId,
	})
}

func (s *SQLiteStorage) GetDiscordMemberByServerId(serverId int) ([]*types.DiscordMember, error) {
	rows, err := s.db.Query("SELECT * FROM DiscordMember WHERE server_id = ?", serverId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*types.DiscordMember
	for rows.Next() {
		member := new(types.DiscordMember)
		if err := rows.Scan(&member.UserId, &member.RoleId, &member.ServerId); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *SQLiteStorage) GetDiscordMemberByUserId(userId int) ([]*types.DiscordMember, error) {
	rows, err := s.db.Query("SELECT * FROM DiscordMember WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*types.DiscordMember
	for rows.Next() {
		member := new(types.DiscordMember)
		if err := rows.Scan(&member.UserId, &member.RoleId, &member.ServerId); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *SQLiteStorage) DeleteDiscordMember(userId int, serverId int) error {
	return s.Delete("DiscordMember", map[string]any{
		"user_id":   userId,
		"server_id": serverId,
	})
}
