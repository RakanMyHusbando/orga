package storage

import (
	"encoding/json"
	"fmt"

	"github.com/RakanMyHusbando/orga/types"
)

/* ------------------------------ Team ------------------------------ */

func (s *SQLiteStorage) CreateTeam(team *types.Team) error {
	return s.Insert("Team", map[string]any{
		"guild_id":     team.GuildId,
		"name":         team.Name,
		"abbreviation": team.Abbreviation,
	})
}

func (s *SQLiteStorage) GetTeam() ([]*types.Team, error) {
	rows, err := s.db.Query("SELECT * FROM Team")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var teams []*types.Team
	for rows.Next() {
		team := new(types.Team)
		if err := rows.Scan(&team.Id, &team.GuildId, &team.Name, &team.Abbreviation); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (s *SQLiteStorage) GetTeamById(id int) ([]*types.Team, error) {
	row := s.db.QueryRow("SELECT * FROM Team WHERE id = ?", id)
	team := new(types.Team)
	if err := row.Scan(&team.Id, &team.GuildId, &team.Name, &team.Abbreviation); err != nil {
		return nil, err
	}
	return []*types.Team{team}, nil
}

func (s *SQLiteStorage) UpdateTeam(team *types.Team, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(team)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Update("Team", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteTeam(id int) error {
	return s.Delete("Team", map[string]any{"id": id})
}

/* ------------------------------ Role ------------------------------ */

func (s *SQLiteStorage) CreateTeamRole(role *types.TeamRole) error {
	return s.Insert("TeamRole", map[string]any{
		"name":        role.Name,
		"description": role.Description,
	})
}

func (s *SQLiteStorage) GetTeamRole() ([]*types.TeamRole, error) {
	rows, err := s.db.Query("SELECT * FROM TeamRole")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*types.TeamRole
	for rows.Next() {
		role := new(types.TeamRole)
		if err := rows.Scan(&role.Id, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (s *SQLiteStorage) GetTeamRoleByUserId(userId int) ([]*types.TeamRole, error) {
	row := s.db.QueryRow("SELECT * FROM TeamRole WHERE user_id = ?", userId)
	role := new(types.TeamRole)
	if err := row.Scan(&role.Id, &role.Name, &role.Description); err != nil {
		return nil, err
	}
	return []*types.TeamRole{role}, nil
}

func (s *SQLiteStorage) UpdateTeamRole(role *types.TeamRole, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(role)
	if err != nil {
		return fmt.Errorf("[storage.team] %v", err)
	}
	json.Unmarshal(bytes, &values)
	return s.Update("TeamRole", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeletTeamRole(id int) error {
	return s.Delete("TeamRole", map[string]any{"id": id})
}

/* ------------------------------ Member ------------------------------ */

func (s *SQLiteStorage) CreateTeamMember(member *types.TeamMember) error {
	return s.Insert("TeamMember", map[string]any{
		"user_id": member.UserId,
		"team_id": member.TeamId,
		"role_id": member.RoleId,
	})
}

func (s *SQLiteStorage) GetTeamMemberByUserId(userId int) ([]*types.TeamMember, error) {
	rows, err := s.db.Query("SELECT * FROM TeamMember WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*types.TeamMember
	for rows.Next() {
		member := new(types.TeamMember)
		if err := rows.Scan(&member.UserId, &member.TeamId, &member.RoleId); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *SQLiteStorage) GetTeamMemberByTeamId(teamId int) ([]*types.TeamMember, error) {
	rows, err := s.db.Query("SELECT * FROM TeamMember WHERE team_id = ?", teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*types.TeamMember
	for rows.Next() {
		member := new(types.TeamMember)
		if err := rows.Scan(&member.UserId, &member.TeamId, &member.RoleId); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *SQLiteStorage) DeleteTeamMember(id int) error {
	return s.Delete("TeamMember", map[string]any{"id": id})
}
