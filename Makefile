build:
	go build -o opsctl

run:
	go run main.go

docker-build:
	docker build -t opsctl .

clean:
	rm -f opsctl
