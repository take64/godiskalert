.PHONY: build
build:
	go build cmd/godiskalert/godiskalert.go

.PHONY: test
test:
	./godiskalert -slack="<<slack hook URL>>"
