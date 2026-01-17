build:
	go build -o explainit

install:
	go build -o explainit
	mv explainit $(PREFIX)/bin/

run:
	go run .

help:
	@echo "make build   - Build binary"
	@echo "make install - Install globally"
	@echo "make run     - Run locally"
