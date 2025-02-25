test:
	@ go test ./...

auto_test:
	@ fswatch -or . | xargs -n1 -I{} go test ./...

run:
	@ go run main.go

build:
	@ go build -o bin/slut cmd/slut.go
	@ chmod +x bin/slut

gen:
	@ flog -f json -d 500ms -l
