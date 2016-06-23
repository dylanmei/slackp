VERSION := 0.1
TARGET  := slackp

default: build

deps:
	go get -v -u ./...

build: clean
	go build -v -o bin/$(TARGET)

release: clean
	 GOARCH=amd64 GOOS=linux go build -v -a -ldflags '-w -linkmode external -extldflags "-static" -X main.Version=$(VERSION)' -o bin/$(TARGET) .

docker-publish: release
	docker build -t dylanmei/$(TARGET):$(VERSION) .
	docker push dylanmei/$(TARGET):$(VERSION)
	docker tag dylanmei/$(TARGET):$(VERSION) dylanmei/$(TARGET):latest
	docker push dylanmei/$(TARGET):latest

clean:
	rm -rf bin
