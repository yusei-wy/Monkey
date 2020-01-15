run:
	go run main.go

test:
	go test -v ./ast
	go test -v ./lexer
	go test -v ./parser
