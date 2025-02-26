test:
	@ go test ./...

auto_test:
	@ fswatch -or . | xargs -n1 -I{} go test ./...

run:
	@ go run main.go

build:
	@ go build -o bin/loks cmd/loks.go
	@ chmod +x bin/loks

gen:
	@ flog -f json -d 500ms -l
