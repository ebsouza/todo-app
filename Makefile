dependencies:
	rm -f go.sum go.mod
	go mod init github.com/ebsouza/todo-app
	go get .
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