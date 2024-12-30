package models


type TaskType string

const (
	MapTask TaskType = "map"
	ReduceTask TaskType = "reduce"
)

type Task struct {
	ID string `json:"id"`
	Type TaskType `json:"type"`
	Data []byte `json:"data"`
	Status string `json:"status"`
	Key string `json:"key,omitempty"`
	Values []string `json:"value,omitempty"`
}
