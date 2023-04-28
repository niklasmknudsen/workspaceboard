package repositories

import (
	"errors"
	"example/WorkspaceBoard/configurations"
	"example/WorkspaceBoard/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BoardRepository interface {
	GetBoards() []models.Board
	CreateBoard(board models.Board) int64
	DeleteBoard(id int64) (bool, error)
}

// method to get all workspaces from database
func GetBoards() []models.Board {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	query := "SELECT * FROM boards"
	rows, err := connection.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var boards []models.Board
	for rows.Next() {
		var board models.Board

		err := rows.Scan(&board.ID, &board.Name, &board.Workspace, &board.Created)
		if err != nil {
			log.Fatal(err)
		}

		boards = append(boards, board)
	}

	return boards
}

func CreateBoard(newBoard models.Board) int64 {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	sql := "INSERT INTO boards(name) VALUES (" + newBoard.Name + ")"

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

func DeleteBoard(id int64) (bool, error) {
	connection, err := configurations.EstablishConnectionToDatabase()
	fmt.Println("connection to database has etablished...")

	deleteStatement := "DELETE FROM boards WHERE ID=?"
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

	fmt.Println("succesfully deleted boards")
	return true, nil
}
