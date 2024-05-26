package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Status int

const (
	Created    Status = iota //Status = 0
	InProgress               //Status = 1
	Completed                //Status = 2
)

// Status (ENUM) to string
func (e Status) String() string {
	return [...]string{"Created", "InProgress", "Completed"}[e]
}

// Strin to Status (ENUM)
func ParseStatus(s string) (Status, error) {
	switch s {
	case "Created":
		return Created, nil
	case "InProgress":
		return InProgress, nil
	case "Completed":
		return Completed, nil
	}

	return 0, errors.New("invalid status")
}

// Allows to store Status values in a db
func (e Status) Value() (driver.Value, error) {
	return e.String(), nil
}

// Allows us to read the Status type from a bd
func (e *Status) Scan(value interface{}) error {
	estadoStr, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for Status")
	}
	estado, err := ParseStatus(string(estadoStr))
	if err != nil {
		return err
	}
	*e = estado
	return nil
}

// Deselarised a Status value from a Json
func (e *Status) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}
	status, err := ParseStatus(statusStr)
	if err != nil {
		return err
	}
	*e = status
	return nil
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Date        string `json:"date"`
}

func GetTaskList(db *sql.DB) ([]Task, error) {
	query := `SELECT name, description, status, DATE_FORMAT(date, '%Y-%m-%d') FROM tasks`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Name, &task.Description, &task.Status, &task.Date)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTask(db *sql.DB, task Task) error {
	query := `INSERT INTO tasks (name, description, status, date) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, task.Name, task.Description, task.Status, task.Date)
	return err
}

func UpdateTask(db *sql.DB, task Task, id int) error {
	present, err := TaskPresent(db, id)
	if err != nil {
		return err
	}
	if !present {
		return errors.New("ID was not found in the database")
	}

	query := `UPDATE tasks SET name = ?, description = ?, status = ?, date = ? WHERE id = ?`
	_, err = db.Exec(query, task.Name, task.Description, task.Status, task.Date, id)
	return err
}

func UpdateTaskStatus(db *sql.DB, status Status, id int) error {
	present, err := TaskPresent(db, id)
	if err != nil {
		return err
	}
	if !present {
		return errors.New("ID was not found in the database")
	}

	query := `UPDATE tasks SET status = ? WHERE id = ?`
	_, err = db.Exec(query, status, id)
	return err
}

func DeleteTask(db *sql.DB, id int) error {
	present, err := TaskPresent(db, id)
	if err != nil {
		return err
	}
	if !present {
		return errors.New("ID was not found in the database")
	}

	query := `DELETE FROM tasks WHERE id = ?`
	_, err = db.Exec(query, id)
	return err
}

func TaskPresent(db *sql.DB, id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
