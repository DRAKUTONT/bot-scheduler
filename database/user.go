package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserID   int
	IsActive bool
	TaskList string
}

type Database struct {
	db *sql.DB
}

func New(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open db %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open db %w", err)
	}
	return &Database{db: db}, nil
}

func (db *Database) AddUser(userID int) error {
	query := `INSERT INTO users (userID, isActive, taskList) VALUES (?, ?, ?)`

	if _, err := db.db.Exec(query, userID, true, ""); err != nil {
		return fmt.Errorf("can't save user: %w", err)
	}

	return nil
}

func (db *Database) GetUsers() ([]User, error) {
	query := `SELECT userID, taskList FROM users WHERE isActive == true`

	rows, err := db.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("can't found users: %w", err)
	}

	defer rows.Close()
	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.UserID, &u.TaskList)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	return users, nil
}



func (db *Database) UpdateTask(userID int, task string) error {
	query := `UPDATE users SET taskList = ? WHERE userID = ?`

	if _, err := db.db.Exec(query, task, userID); err != nil {
		return fmt.Errorf("can't add task: %w", err)
	}

	return nil
}

func (db *Database) UpdateIsActive(userID int, isActive bool) error {
	query := `UPDATE users SET isActive = ? WHERE userID = ?`

	if _, err := db.db.Exec(query, isActive, userID); err != nil {
		return fmt.Errorf("can't add task: %w", err)
	}

	return nil
}

func (db *Database) GetTask(userID int) (string, error) {
	q := `SELECT taskList FROM users WHERE userID = ?`

	var task string

	if err := db.db.QueryRow(q, userID).Scan(&task); err != nil {
		return "", fmt.Errorf("can't get task: %w", err)
	}

	return task, nil
}

func (db *Database) IsUserExists(userID int) (bool, error) {
	q := `SELECT COUNT(*) FROM users WHERE userID = ?`

	var count int

	if err := db.db.QueryRow(q, userID).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check if user exists: %w", err)
	}

	return count > 0, nil
}

func (db *Database) Init() error {
	query := `CREATE TABLE IF NOT EXISTS users (userID INTEGER, isActive INTEGER, taskLits)`

	_, err := db.db.Exec(query)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
