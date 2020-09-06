.PHONY: build

# go build command
build:
	@go build -v -o SMMOB ./main.go

# go run command
run:
	make build
	@./SMMOB

# gcloud deploy command
deploy:
	@gcloud app deploy
