package storage

import (
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

	log.Println("Storage: successfully get teams")

	return teams, nil
}

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
