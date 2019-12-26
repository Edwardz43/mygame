test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...

docker:
	docker build -t mygame .

run:
	docker-compose up -d --build

stop:
	docker-compose down
