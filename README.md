# TODO App

This project was written in Go language and Gin framework.

The app features are very basic. However, the project structure received more effort on development.


### Building and running

```
go mod init todo/server
go get .
GIN_MODE=release go run .
```


### CURL


**GET:** /tasks/id
```
curl localhost:8080/tasks/<task_id>
```

**GET:** /tasks
```
curl localhost:8080/tasks
```

**POST:** /tasks
```
curl -X POST localhost:8080/tasks -H 'Content-Type: application/json'  -d '{"id": "4", "title": "Hard Task", "description":"Do something", "status": "OK}'
```

**DELETE:** /tasks/id
```
curl -X DELETE localhost:8080/tasks/<task_id>
```
