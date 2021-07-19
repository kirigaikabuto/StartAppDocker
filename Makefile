all: one
.PHONY: all

one:
	git add .
	git commit -m "sdsd"
	git push origin1 master

run:
	go run main.go