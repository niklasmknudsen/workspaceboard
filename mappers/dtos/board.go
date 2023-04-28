package dtos

type Board struct {
	ID        int64     `json:id`
	Name      string    `json:name`
	workspace WorkSpace `json:workspaceboard_id`
	Tasks     []Task    `json:tasks`
}
