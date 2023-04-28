package repositories

import (
	"errors"
	"example/WorkspaceBoard/configurations"
	"example/WorkspaceBoard/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type WorkSpaceRepository interface {
	GetWorkSpaces() []models.WorkSpaceBoard
	CreateWorkSpace(workspace models.WorkSpaceBoard) int64
	DeleteWorkSpace(id int64) (bool, error)
}

// method to get all workspaces from database
func GetWorkSpaces() []models.WorkSpaceBoard {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	query := "SELECT * FROM workspaces"
	rows, err := connection.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var workspaces []models.WorkSpaceBoard
	for rows.Next() {
		var workspaceBoard models.WorkSpaceBoard

		err := rows.Scan(&workspaceBoard.ID, &workspaceBoard.Name, &workspaceBoard.Created)
		if err != nil {
			log.Fatal(err)
		}

		workspaces = append(workspaces, workspaceBoard)
	}

	return workspaces
}

func CreateWorkSpace(newWorkspaceBoard models.WorkSpaceBoard) int64 {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	sql := "INSERT INTO workspaces(name) VALUES (" + newWorkspaceBoard.Name + ")"

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

func DeleteWorkSpace(id int64) (bool, error) {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	deleteStatement := "DELETE FROM workspaces WHERE ID=?"

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

	fmt.Println("succesfully deleted workspace")
	return true, nil
}
