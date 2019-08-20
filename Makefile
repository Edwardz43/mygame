BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

# engine:
# 	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

# clean:
# 	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t mygame .

run:
	docker-compose up -d --build

stop:
	docker-compose down
