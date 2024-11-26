package storage

import "github.com/RakanMyHusbando/shogun/types"

/* ------------------------------ Server ------------------------------ */

func (s *SQLiteStorage) CreateServer(server *types.DiscordServer) error {
	return s.Insert("Server", map[string]any{
		"name": server.Name,
	})
}

func (s *SQLiteStorage) GetServer() ([]*types.DiscordServer, error) {
	rows, err := s.db.Query("SELECT * FROM Server")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var servers []*types.DiscordServer
	for rows.Next() {
		server := new(types.DiscordServer)
		if err := rows.Scan(&server.Id, &server.Name); err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func (s *SQLiteStorage) GetServerById(id int) ([]*types.DiscordServer, error) {
	row := s.db.QueryRow("SELECT * FROM Server WHERE id = ?", id)
	server := new(types.DiscordServer)
	if err := row.Scan(&server.Id, &server.Name); err != nil {
		return nil, err
	}
	return []*types.DiscordServer{server}, nil
}

func (s *SQLiteStorage) UpdateServer(server *types.DiscordServer, id int) error {
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

func (s *SQLiteStorage) GetDiscordById(id int) ([]*types.DiscordRole, error) {
	row := s.db.QueryRow("SELECT * FROM DiscordRole WHERE id = ?", id)
	role := new(types.DiscordRole)
	if err := row.Scan(&role.Id, &role.Name, &role.Description); err != nil {
		return nil, err
	}
	return []*types.DiscordRole{role}, nil
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
