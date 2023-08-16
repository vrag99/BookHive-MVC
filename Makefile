all: 
	go build -o bookhive ./cmd/main.go
	./bookhive

setup:
	chmod 777 ./scripts/setup.sh
	./scripts/setup.sh

clean:
	rm bookhive