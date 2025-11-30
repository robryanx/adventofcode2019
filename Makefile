buildall:
	go build -o bin ./days/...
.PHONY: buildall

testall: buildall
	go test ./
.PHONY: testall
