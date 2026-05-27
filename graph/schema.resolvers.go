package graph

import (
    "context"
    "fmt"

    "example.com/pz11-graphql/graph/model"
    "example.com/pz11-graphql/internal/store"
)

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

// Tasks returns list of all tasks
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
    result := make([]*model.Task, 0, len(store.Tasks))
    for _, t := range store.Tasks {
        result = append(result, &model.Task{
            ID:          t.ID,
            Title:       t.Title,
            Description: t.Description,
            Done:        t.Done,
        })
    }
    return result, nil
}

// Task returns task by ID
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
    for _, t := range store.Tasks {
        if t.ID == id {
            return &model.Task{
                ID:          t.ID,
                Title:       t.Title,
                Description: t.Description,
                Done:        t.Done,
            }, nil
        }
    }
    return nil, nil
}

// CreateTask creates new task
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
    newID := fmt.Sprintf("t_%03d", len(store.Tasks)+1)
    task := &store.Task{
        ID:          newID,
        Title:       input.Title,
        Description: input.Description,
        Done:        false,
    }
    store.Tasks = append(store.Tasks, task)
    return &model.Task{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Done:        task.Done,
    }, nil
}

// UpdateTask updates existing task
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTaskInput) (*model.Task, error) {
    for _, t := range store.Tasks {
        if t.ID == id {
            if input.Title != nil {
                t.Title = *input.Title
            }
            if input.Description != nil {
                t.Description = input.Description
            }
            if input.Done != nil {
                t.Done = *input.Done
            }
            return &model.Task{
                ID:          t.ID,
                Title:       t.Title,
                Description: t.Description,
                Done:        t.Done,
            }, nil
        }
    }
    return nil, fmt.Errorf("task not found")
}

// DeleteTask deletes task
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
    for i, t := range store.Tasks {
        if t.ID == id {
            store.Tasks = append(store.Tasks[:i], store.Tasks[i+1:]...)
            return true, nil
        }
    }
    return false, nil
}
