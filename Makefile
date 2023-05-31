dependencies:
	go mod init github.com/ebsouza/todo-app
	go get .
	go get -t github.com/ebsouza/todo-app/tasks

docker: dependencies
	docker-compose up --force-recreate

test:
	go test ./...

format:
	gofmt -s -w .
	go fmt github.com/...
	go fmt ...