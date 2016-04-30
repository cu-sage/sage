build:
	go get ./...
	go build -o bin/sage

test:
	go test ./...

docker:
	docker build -t sage-assess:$(BUILD_NUMBER) .

deploy:
	-docker stop sage
	-docker rm sage
	docker run -p 8081:8080 --name sage --net sagenetwork -d sage-assess:$(DEPLOY_TAG)	
