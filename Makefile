all: 
	go build -o bookhive ./cmd/main.go
	./bookhive

clean:
	rm bookhive