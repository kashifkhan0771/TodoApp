package models

import (
	"time"

	"github.com/fatih/structs"
)

type Task struct {
	ID      string                 `json:"id" structs:"id" bson:"_id" db:"id"`
	Title   string                 `json:"title" structs:"title" bson:"title" db:"title"`
	Desc    map[string]interface{} `json:"desc" structs:"desc" bson:"desc" db:"desc"`
	AddedOn time.Time              `json:"added_on" structs:"added_on" bson:"added_on" db:"added_on"`
	TodoOn  time.Time              `json:"todo_on" structs:"todo_on" bson:"todo_on" db:"todo_on"`
	Status  string                 `json:"status" structs:"status" bson:"status" db:"status"`
}

func (t *Task) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of Agent model
func (t *Task) Names() []string {
	fields := structs.Fields(t)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
