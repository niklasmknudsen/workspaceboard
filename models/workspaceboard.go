package models

type WorkSpaceBoard struct {
	ID      int64
	Name    string
	Boards  []Board
	Created string
}
