run-start:
	go run cmd/start/main.go

run-end:
	go run cmd/end/main.go

gotags:
	gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .

test:
	go test -v ./math
