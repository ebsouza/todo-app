# TODO App

This project was written in Go language and it has Gin and GORM as essential parts.

The app features are very basic. However, it shows how everything work together.


### Running

```
docker-compose up
```


### CURL


**GET:** /tasks/id
```
curl localhost:8080/tasks/<task_id>
```

**GET:** /tasks
```
curl localhost:8080/tasks/
```

**POST:** /tasks
```
curl -X POST localhost:8080/tasks/ -H 'Content-Type: application/json'  -d '{"id": "1", "title": "Hard Task", "description":"Do something", "status": "OK"}'
```

**DELETE:** /tasks/id
```
curl -X DELETE localhost:8080/tasks/<task_id>
```





# Links

https://aprendagolang.com.br/2022/09/01/como-fazer-teste-unitario-no-gorm-com-testify-e-sqlmock/
https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj
https://gorm.io/docs/
https://dev.to/devniklesh/how-to-read-env-variables-in-golang-using-viper-2jd1
https://github.com/peter-evans/docker-compose-healthcheck