run:
	go run main.go
build:
	go build -o bin/main main.go
docker-build: 
	sudo su docker build -t buddy_test .
docker-run:
	sudo su docker run -p -d 8080:8080 buddy_test
docker-compose:
	sudo docker-compose up -d