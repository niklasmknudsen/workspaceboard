package repositories

import (
	"errors"
	"example/WorkspaceBoard/configurations"
	"example/WorkspaceBoard/models"
	"fmt"
	"log"
)

type TaskRepository interface {
	GetTasks() []models.Task
	CreateTask(task models.Task) int64
	DeleteTask(id int64) (bool, error)
	GetById(id int64) models.Task
	GetAllByBoard(id int64) []models.Task
}

func GetTasks() []models.Task {
	connection, err := configurations.EstablishConnectionToDatabase()
	if err != nil {
		log.Fatal(err)
		return []models.Task{}
	}

	query := "SELECT * FROM tasks"
	rows, err := connection.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Description)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func CreateTask(task models.Task) int64 {
	connection, err := configurations.EstablishConnectionToDatabase()

	sql := "INSERT INTO workspaces(description) VALUES (" + task.Description + ")"

	res, err := connection.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return lastId
}

func DeleteTask(id int64) (bool, error) {
	connection, err := configurations.EstablishConnectionToDatabase()

	deleteStatement := "DELETE FROM tasks WHERE ID=?"

	stmt, err := connection.Prepare(deleteStatement)
	if err != nil {
		log.Fatal(err)
		return false, errors.New(err.Error())
	}

	response, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return false, errors.New(err.Error())
	}

	_, error := response.RowsAffected()
	if error != nil {
		log.Fatal(err)
		return false, errors.New(err.Error())
	}

	fmt.Println("succesfully deleted tasks")
	return true, nil
}

func GetById(id int64) models.Task {
	connection, err := configurations.EstablishConnectionToDatabase()
	if err != nil {
		log.Fatal(err)
		return models.Task{}
	}
	var task models.Task
	query := "SELECT task_id, description, board_id FROM tasks WHERE ID= ?"

	connection.QueryRow(query, id).Scan(&task.ID, &task.Description, &task.Board)

	return task
}

func GetAllByBoard(id int64) []models.Task {
	connection, err := configurations.EstablishConnectionToDatabase()
	if err != nil {
		log.Fatal(err)
		return []models.Task{}
	}
	var tasks []models.Task

	query := "SELECT * FROM tasks WHERE board_id= ?"

	rows, err := connection.Query(query)
	if err != nil {
		return []models.Task{}
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Description, &task.Board)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}
