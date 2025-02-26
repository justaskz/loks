test:
	@ go test ./...

auto_test:
	@ fswatch -or . | xargs -n1 -I{} go test ./...

build:
	@ go build -o bin/loks cmd/loks/main.go
	@ chmod +x bin/loks

gen:
	@ flog -f json -d 500ms -l
