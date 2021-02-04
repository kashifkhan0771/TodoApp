package mongo

import (
	"context"
	"github.com/kashifkhan0771/TodoApp/db"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kashifkhan0771/TodoApp/models"
)

func Test_client_AddTask(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "task-management-mongo-db")

	t.Parallel()
	type args struct {
		ctx  context.Context
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - add new task into db",
			args: args{task: &models.Task{
				Title: "Task 1",
				Desc: map[string]interface{}{
					"User":   "Kashif Khan",
					"Detail": "Drink Water Daily",
				},
				AddedOn: time.Now().UTC().Truncate(time.Minute),
				TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
				Status:  "InProcess",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, _ := NewClient(db.Option{})
			_, err := c.AddTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}

func Test_client_DeleteTask(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "task-management-mongo-db")

	c, _ := NewClient(db.Option{})
	task := &models.Task{
		Title: "Task To Delete",
		Desc: map[string]interface{}{
			"User":   "Anonymous",
			"Detail": "Delete this task",
		},
		AddedOn: time.Now().UTC().Truncate(time.Minute),
		TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
		Status:  "ToDelete",
	}
	_, _ = c.AddTask(context.TODO(), task)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success - delete task from database",
			args:    args{id: task.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := c.DeleteTask(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetTaskByID(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "task-management-mongo-db")
	c, _ := NewClient(db.Option{})
	task := &models.Task{
		Title: "New Task",
		Desc: map[string]interface{}{
			"User":   "Shahzad Haider",
			"Detail": "Get This Task From Database",
		},
		AddedOn: time.Now().UTC().Truncate(time.Minute),
		TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
		Status:  "Pending",
	}

	taskID, _ := c.AddTask(context.TODO(), task)

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Task
		wantErr bool
	}{
		{
			name:    "Success - get task from database",
			args:    args{id: taskID},
			want:    task,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := c.GetTaskByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaskByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateTask(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "task-management-mongo-db")
	c, _ := NewClient(db.Option{})
	task := &models.Task{
		Title: "New Task",
		Desc: map[string]interface{}{
			"User":   "Shahzad Haider",
			"Detail": "Update this task",
		},
		AddedOn: time.Now().UTC().Truncate(time.Minute),
		TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
		Status:  "Pending",
	}

	updateTaskID, _ := c.AddTask(context.TODO(), task)
	type args struct {
		ctx  context.Context
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - update task in database",
			args: args{task: &models.Task{
				ID:    updateTaskID,
				Title: "New Task Updated",
				Desc: map[string]interface{}{
					"User":   "Shahzad Haider",
					"Detail": "Task Updated",
				},
				AddedOn: time.Now().UTC().Truncate(time.Minute),
				TodoOn:  time.Date(2021, 02, 15, 12, 00, 00, 00, time.UTC),
				Status:  "Pending",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
