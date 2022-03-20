dev:
	go run .

test:
	go test ./... 
test_cover:
	go test ./... --cover

test_cover_profile:
	go test ./... --cover -coverprofile=cover.tmp
	go tool cover -html=cover.tmp -o cover.html

.PHONY: dev test test_cover