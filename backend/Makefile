wire:
	@cd internal/bootstrap && wire
migrate-up:
	migrate -path migrations -database "postgres://oss:${db_password}@localhost:${db_port}/oss?sslmode=disable" -verbose up
