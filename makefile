build:
	go build -o server cmd/main/main.go

run: build
	./server

watch:
	clear
	ulimit -n 1000
	reflex -s -r '\.go$$' make run