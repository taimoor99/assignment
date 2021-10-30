run_containers:
	@docker container start mongodb || true
	@docker container start go-tuts || true

build_containers:
	docker-compose up -d --force-recreate --no-deps --build

run_tests:
	go test ./...