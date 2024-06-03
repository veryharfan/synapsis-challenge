migration_up:
	go run main.go --migration up

migration_rollback:
	go run main.go --migration rollback

run:
	go run main.go