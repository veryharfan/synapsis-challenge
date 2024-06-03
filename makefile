migration_up:
	go run migration/main/main.go up

migration_rollback:
	go run migration/main/main.go rollback

run:
	go run main.go