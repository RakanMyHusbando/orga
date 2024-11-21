package storage

import (
	"encoding/json"
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

/* ------------------------------ team ------------------------------ */

// POST
func (s *SQLiteStorage) CreateTeam(team *types.ReqTeam) error {
	prep, err := s.db.Prepare(`INSERT INTO Team (guild_id, name, abbreviation) VALUES (?,?,?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		team.GuildId,
		team.Name,
		team.Abbreviation,
	); err != nil {
		return err
	}

	prep.Close()
	log.Println("Storage: successfully created team")

	return nil
}

// GET
func (s *SQLiteStorage) GetTeam() ([]*types.ResTeam, error) {
	rows, err := s.db.Query(`SELECT * FROM Team`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []*types.ResTeam{}
	for rows.Next() {
		team := new(types.ResTeam)
		if err := rows.Scan(
			&team.Id,
			&team.GuildId,
			&team.Name,
			&team.Abbreviation,
		); err != nil {
			return nil, err
		}
		// TODO: add Team-Member relation
		teams = append(teams, team)
	}

	log.Println("[Storage.team]: successfully get teams")

	return teams, nil
}

// GET
func (s *SQLiteStorage) GetTeamById(id int) (*types.ResTeam, error) {
	row := s.db.QueryRow(`SELECT * FROM Team WHERE id = ?`, id)

	team := new(types.ResTeam)
	if err := row.Scan(
		&team.Id,
		&team.GuildId,
		&team.Name,
		&team.Abbreviation,
	); err != nil {
		return nil, err
	}
	// TODO: add Team-Member relation

	return team, nil
}

// PATCH
func (s *SQLiteStorage) UpdateTeam(team *types.ResTeam) error {
	prep, err := s.db.Prepare(`UPDATE Team SET guild_id=?, name=?, abbreviation=? WHERE id=?`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(
		team.GuildId,
		team.Name,
		team.Abbreviation,
		team.Id,
	); err != nil {
		return err
	}
	prep.Close()
	log.Println("[Storage.team]: successfully updated team")
	return nil
}

/* ------------------------------ team role ------------------------------ */

// POST
func (s *SQLiteStorage) CreateTeamRole(teamRole *types.ReqTeamRole) error {
	var values map[string]any
	bytes, err := json.Marshal(teamRole)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	if err = s.Insert("TeamRole", values); err != nil {
		return err
	}
	return nil
}

// GET
func (s *SQLiteStorage) GetTeamRole() ([]*map[string]any, error) {
	data, err := s.Select("TeamRole", nil, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DELET
func (s *SQLiteStorage) DeletTeamRole(id int) error {
	return nil
}

/* ------------------------------ team member ------------------------------ */
