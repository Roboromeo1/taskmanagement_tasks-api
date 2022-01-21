package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	elasticsearch "github.com/Roboromeo1/taskmanagement_tasks-api/src/clients"
	"github.com/Roboromeo1/taskmanagement_tasks-api/src/domain/queries"
	"github.com/Roboromeo1/taskmanagement_utils/rest_errors"
	"strings"
)

const (
	indexTasks = "items"
	typeTask   = "_doc"
)

func (i *Task) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexTasks, typeTask, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Task) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexTasks, typeTask, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Task) Search(query queries.EsQuery) ([]Task, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexTasks, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Task, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var task Task
		if err := json.Unmarshal(bytes, &task); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		task.Id = hit.Id
		items[index] = task
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
