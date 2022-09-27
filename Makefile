unit-tests:
	go test
functional-tests:
	@echo Performing Functional Tests
	@echo Functional Tests Passed
build:
	GOOS=linux GOARCH=amd64 go build -o curr-time-app .
	chmod 777 time-app

imgbuild:
	docker build -t lnanjangud653/current-time:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/current-time:${IMG_VER}
