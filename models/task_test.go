package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_Map(t1 *testing.T) {
	t1.Parallel()
	type fields struct {
		ID      string
		Title   string
		Desc    map[string]interface{}
		AddedOn time.Time
		TodoOn  time.Time
		Status  string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - Convert struct to map",
			fields: fields{
				ID:    "1",
				Title: "Task 1",
				Desc: map[string]interface{}{
					"User":   "Kashif Khan",
					"Detail": "Complete New Microservice",
				},
				AddedOn: time.Now().UTC().Truncate(time.Minute),
				TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
				Status:  "Pending",
			},
			want: map[string]interface{}{
				"id":    "1",
				"title": "Task 1",
				"desc": map[string]interface{}{
					"User":   "Kashif Khan",
					"Detail": "Complete New Microservice",
				},
				"added_on": time.Now().UTC().Truncate(time.Minute),
				"todo_on":  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
				"status":   "Pending",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t1.Parallel()
			t := &Task{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				Desc:    tt.fields.Desc,
				AddedOn: tt.fields.AddedOn,
				TodoOn:  tt.fields.TodoOn,
				Status:  tt.fields.Status,
			}
			if got := t.Map(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Names(t1 *testing.T) {
	t1.Parallel()
	type fields struct {
		ID      string
		Title   string
		Desc    map[string]interface{}
		AddedOn time.Time
		TodoOn  time.Time
		Status  string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - Names of task struct field",
			fields: fields{
				ID:    "1",
				Title: "Task 1",
				Desc: map[string]interface{}{
					"User":   "Kashif Khan",
					"Detail": "Complete New Microservice",
				},
				AddedOn: time.Now().UTC().Truncate(time.Minute),
				TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
				Status:  "Pending",
			},
			want: []string{"id", "title", "desc", "added_on", "todo_on", "status"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t1.Parallel()
			t := &Task{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				Desc:    tt.fields.Desc,
				AddedOn: tt.fields.AddedOn,
				TodoOn:  tt.fields.TodoOn,
				Status:  tt.fields.Status,
			}
			if got := t.Names(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
