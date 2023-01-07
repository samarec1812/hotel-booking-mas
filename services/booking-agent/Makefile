.PHONY: migrate-up
migrate-up:
	migrate -path ./schema -database 'postgres://admin:qwerty@localhost:5436/booking-agent-db?sslmode=disable' up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema -database 'postgres://admin:qwerty@localhost:5436/booking-agent-db?sslmode=disable' down

.PHONY: migrate-force
migrate-force:
	migrate -path ./schema -database 'postgres://admin:qwerty@localhost:5436/booking-agent-db?sslmode=disable' force 1