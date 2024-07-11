dependencies:
	rm -f go.sum go.mod
	go mod init github.com/ebsouza/todo-app

	go get github.com/gin-gonic/gin@v1.9.0
	go get github.com/google/uuid@v1.3.0
	go get github.com/stretchr/testify@v1.8.1
	go get gorm.io/driver/postgres@v1.5.2
	go get gorm.io/driver/sqlite@v1.5.1
	go get gorm.io/gorm@v1.25.1
	
	go get -t github.com/ebsouza/todo-app/tasks

docker-build: dependencies
	docker-compose up --force-recreate

docker-run:
	docker-compose up --force-recreate

test:
	go test ./...

format:
	gofmt -s -w .
	go fmt github.com/...
	go fmt ...