# TODO App

This project was written in Go language and it has Gin and GORM as essential parts.

The app features are very basic. However, it shows how everything work together.


### Building & Running

```
make docker-build

make docker-run
```

#### Docker version
- Docker version 27.0.3
- Docker compose version v2.28.1

### CURL


**GET:** /tasks/id
```
curl localhost:8080/tasks/<task_id>
```

**GET:** /tasks
```
curl localhost:8080/tasks?limit=<limit>&offset=<offset>&status=<Task.status>
```

**POST:** /tasks
```
curl -X POST localhost:8080/tasks -H 'Content-Type: application/json'  -d '{"title": "Hard Task", "description":"Do something"}'
```

**UPDATE:** /tasks/id
```
curl -X PUT localhost:8080/tasks/<task_id> -H 'Content-Type: application/json'  -d '{"title": "Easy Task", "description":"Do nothing", "status": "OK"}'
```

**DELETE:** /tasks/id
```
curl -X DELETE localhost:8080/tasks/<task_id>
```
