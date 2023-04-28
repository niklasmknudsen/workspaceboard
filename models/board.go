package models

type Board struct {
	ID        int64
	Name      string
	Workspace int64
	Tasks     []Task
	Created   string
}
