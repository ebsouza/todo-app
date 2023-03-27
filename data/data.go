package data

import "todo/server/domain"


var Tasks = []domain.Task{
	{ID: "1", Title: "Task 1", Description: "Do something", Status: "OK"},
	{ID: "2", Title: "Task 2", Description: "Do something again", Status: "OK"},
	{ID: "3", Title: "Task 3", Description: "Ok, do it right", Status: "OK"},
}
