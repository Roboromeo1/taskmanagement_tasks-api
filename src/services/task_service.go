package services

import (
	"github.com/Roboromeo1/taskmanagement_tasks-api/src/domain/queries"
	"github.com/Roboromeo1/taskmanagement_tasks-api/src/domain/tasks"
	"github.com/Roboromeo1/taskmanagement_utils/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &tasksService{}
)

type itemsServiceInterface interface {
	Create(task tasks.Task) (*tasks.Task, rest_errors.RestErr)
	Get(string) (*tasks.Task, rest_errors.RestErr)
	Search(queries.EsQuery) ([]tasks.Task, rest_errors.RestErr)
}

type tasksService struct{}

func (s *tasksService) Create(task tasks.Task) (*tasks.Task, rest_errors.RestErr) {
	if err := task.Save(); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *tasksService) Get(id string) (*tasks.Task, rest_errors.RestErr) {
	item := tasks.Task{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *tasksService) Search(query queries.EsQuery) ([]tasks.Task, rest_errors.RestErr) {
	dao := tasks.Task{}
	return dao.Search(query)
}
