package storage

import (
	"encoding/json"
	"fmt"

	"github.com/RakanMyHusbando/shogun/types"
)

/* ------------------------------ team ------------------------------ */

func (s *SQLiteStorage) CreateTeam(team *types.Team) error {
	var values map[string]any
	bytes, err := json.Marshal(team)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("Team", values)
}

func (s *SQLiteStorage) GetTeam() ([]*map[string]any, error) {
	// TODO: add Team-Member relation
	return s.Select("Team", nil, nil)
}

func (s *SQLiteStorage) GetTeamById(id int) ([]*map[string]any, error) {
	// TODO: add Team-Member relation
	return s.SelectUnique("Team", nil, "id", id)
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

/* ------------------------------ team role ------------------------------ */

func (s *SQLiteStorage) CreateTeamRole(role *types.TeamRole) error {
	var values map[string]any
	bytes, err := json.Marshal(role)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("TeamRole", values)
}

func (s *SQLiteStorage) GetTeamRole() ([]*map[string]any, error) {
	return s.Select("TeamRole", nil, nil)
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
	var values map[string]any
	bytes, err := json.Marshal(member)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("TeamMember", values)
}

func (s *SQLiteStorage) GetTeamMemberByUserId(userId int) ([]*map[string]any, error) {
	return s.SelectUnique("TeamMember", nil, "user_id", userId)
}

func (s *SQLiteStorage) GetTeamMemberByTeamId(teamId int) ([]*map[string]any, error) {
	return s.SelectUnique("TeamMember", nil, "team_id", teamId)
}

func (s *SQLiteStorage) UpdateTeamMember(member *types.TeamMember, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(member)
	if err != nil {
		return fmt.Errorf("[storage.team] error: %v", err)
	}
	json.Unmarshal(bytes, &values)
	return s.Update("TeamMember", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteTeamMember(id int) error {
	return s.Delete("TeamMember", map[string]any{"id": id})
}
