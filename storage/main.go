package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/RakanMyHusbando/shogun/types"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{
		db: db,
	}, nil
}

func RunSQLiteStorage(storage *SQLiteStorage, schemaFile string) error {
	// read sql file with create table statements
	byteContent, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	// execute sql create table statements
	queries := strings.Split(string(byteContent), ";")
	for i := range queries {
		query := strings.TrimSpace(queries[i])
		_, err := storage.db.Exec(query)
		if err != nil {
			return fmt.Errorf("%v [Table: %v]", err.Error(), strings.Split(query, " ")[5])
		}
	}
	return nil
}

type Storage interface {
	// User

	// LeagueOfLegends

	// GameAccount

	// Guild

	// Guild.Role

	// Guild.Member

	// Team
	InsertTeam(team map[string]any) error
	GetTeam() ([]*map[string]any, error)
	GetTeamById(id int) (*map[string]any, error)
	UpdateTeam(team *types.Team, id int) error
	DeleteTeam(id int) error
	// Team.Role
	CreateTeamRole(role *types.TeamRole) error
	GetTeamRole() ([]*map[string]any, error)
	UpdateTeamRole(role *types.TeamRole, id int) error
	DeletTeamRole(id int) error
	// Team.Member
	CreateTeamMember(member *types.TeamMember) error
	GetTeamMemberByUserId(userId int) ([]*map[string]any, error)
	GetTeamMemberByTeamId(teamId int) ([]*map[string]any, error)
	UpdateTeamMember(member *types.TeamMember, id int) error
	DeleteTeamMember(id int) error
}

type SQLiteStorage struct {
	db *sql.DB
}

/* ------------------------------ SQLite reqs ------------------------------ */

// CREATE
func (s *SQLiteStorage) Insert(table string, insertValues map[string]any) error {
	var keys, values string
	first := true
	for key, value := range insertValues {
		if value != nil {
			if !first {
				keys += ", "
				values += ", "
			} else {
				first = false
			}
			if reflect.TypeOf(value).String() == "string" {
				value = fmt.Sprintf("'%v'", value)
			}
			keys += key
			values = fmt.Sprintln(values, value)
		}
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, keys, values)
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[Srotage.main] inserted element in table ", table)
	return nil
}

// GET
func (s *SQLiteStorage) Select(table string, selectKeys []string, where map[string]any) ([]*map[string]any, error) {
	var selectKeysStr string
	if len(selectKeys) == 0 || selectKeys == nil {
		selectKeysStr = "*"
	} else {
		selectKeysStr = strings.Join(selectKeys, ",")
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE ", selectKeysStr, table)
	first := true
	for key, value := range where {
		if value != nil {
			if !first {
				query += " AND "
			} else {
				first = false
			}
			query += fmt.Sprintf("%s = %v", key, value)
		}
	}
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var elems []*map[string]any
	for rows.Next() {
		elem := make(map[string]any)
		if err := rows.Scan(&elem); err != nil {
			return nil, err
		}
		elems = append(elems, &elem)
	}
	log.Println("[Storage.main] get elements from table: ", table)
	return elems, nil
}

// GET
func (s *SQLiteStorage) SelectUnique(table string, selectKeys []string, whereKey string, whereValeu any) (*map[string]any, error) {
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = %v",
		strings.Join(selectKeys, ","),
		table,
		whereKey,
		whereValeu,
	)
	row := s.db.QueryRow(query)
	elem := make(map[string]any)
	if err := row.Scan(&elem); err != nil {
		return nil, err
	}
	return &elem, nil
}

// PATCH
func (s *SQLiteStorage) Patch(table string, set map[string]any, where map[string]any) error {
	query := fmt.Sprintf(`UPDATE %s SET `, table)
	first := true
	for key, value := range set {
		if value != nil {

			if !first {
				query += ", "
			} else {
				first = false
			}
			query += fmt.Sprintf(`%s = %v`, key, value)
		}
	}
	query += " WHERE "
	first = true
	for key, value := range where {
		if !first {
			query += " AND "
		} else {
			first = false
		}
		query += fmt.Sprintf(`%s = %v`, key, value)
	}
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[Storage.main] element updated from table: ", table)
	return nil
}

// DELETE
func (s *SQLiteStorage) Delete(table string, where map[string]any) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE `, table)
	first := true
	for key, value := range where {
		if value == nil {
			if !first {
				query += " AND "
			} else {
				first = false
			}
			query += fmt.Sprintf(`%s = %v`, key, value)
		}
	}
	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}
	prep.Close()
	log.Println("[Storage.main] delete element from ", table)
	return nil
}

/* ------------------------------ helper functions ------------------------------ */
