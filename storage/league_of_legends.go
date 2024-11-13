package storage

import (
	"fmt"
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

// POST
func (s *SQLiteStorage) CreateLeagueOfLegends(lol *types.ReqLeagueOfLegends) error {
	var insertKeys, insertValues string
	if lol.MainChamps != nil {
		for i := range lol.MainChamps {
			insertKeys += fmt.Sprintf(", champ_%d", i)
			insertValues += fmt.Sprintf(", '%s'", lol.MainChamps[i])
		}
	}
	query := fmt.Sprintf(
		"INSERT INTO UserLeagueOfLegends (user_id, main_role, second_role %s) VALUES (%d, '%s', '%s' %s)",
		insertKeys, lol.UserId, lol.MainRole, lol.SecondRole, insertValues,
	)

	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()

	log.Printf("Storage: successfully added league_of_legends to user with id %v", lol.UserId)

	return nil
}

// GET
func (s *SQLiteStorage) GetLeagueOfLegendsById(userId int) (*types.ResLeagueOfLegends, error) {
	row := s.db.QueryRow(`
		SELECT 
			main_role, 
			second_role, 
			IFNULL(champ_0, ''), 
			IFNULL(champ_1, ''), 
			IFNULL(champ_2, '') 
		FROM 
			UserLeagueOfLegends 
		WHERE 
			user_id = ?`,
		userId,
	)

	userLol := new(types.ResLeagueOfLegends)
	mainChamps := []string{"", "", ""}
	if err := row.Scan(
		&userLol.MainRole,
		&userLol.SecondRole,
		&mainChamps[0],
		&mainChamps[1],
		&mainChamps[2],
	); err != nil {
		return nil, err
	}

	userLol.MainChamps = []string{}
	for i := range mainChamps {
		if mainChamps[i] != "" {
			userLol.MainChamps = append(userLol.MainChamps, mainChamps[i])
		}
	}

	log.Printf("Storage: successfully get league_of_legends from user with id %v", userId)

	return userLol, nil
}

// GET
func (s *SQLiteStorage) GetLeagueOfLegendsWithAccountsById(userId int) (*types.ResLeagueOfLegends, error) {
	userLol, err := s.GetLeagueOfLegendsById(userId)
	if err != nil {
		return nil, err
	}

	accounts, err := s.GetGameAccountByUserId(userId, "league_of_legends")
	if err != nil {
		log.Println(err)
		userLol.Accounts = []string{}
	} else {
		userLol.Accounts = accounts
	}

	return userLol, nil
}

// DELETE
func (s *SQLiteStorage) DeleteLeagueOfLegends(userId int) error {
	log.Println(userId)
	prep, err := s.db.Prepare(`DELETE FROM UserLeagueOfLegends WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(userId); err != nil {
		return err
	}

	prep.Close()

	log.Printf("Storage: successfully delete league_of_legends from user with id %v", userId)

	return nil
}

// PATCH
func (s *SQLiteStorage) UpdateLeagueOfLegends(lol *types.ReqLeagueOfLegends) error {
	var champs string
	if lol.MainChamps != nil {
		for i := range lol.MainChamps {
			champs += fmt.Sprintf(", champ_%d = '%s'", i, lol.MainChamps[i])
		}
	}
	query := fmt.Sprintf(
		`UPDATE UserLeagueOfLegends SET main_role = '%s', second_role = '%s' %s WHERE user_id = %d`,
		lol.MainRole, lol.SecondRole, champs, lol.UserId,
	)

	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}

	prep.Close()

	log.Printf("Storage: successfully update league of legends user with id %v", lol.UserId)

	return nil
}
