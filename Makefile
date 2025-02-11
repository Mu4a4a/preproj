.SILENT:

build:
	docker build -t preproj .

run: build
	docker compose up

shutdown:
	docker compose down

create-migration:
	mkdir -p ./schema
	goose create -dir ./schema $(NAME) sql

migration-up:
	goose -dir ./schema postgres "postgres://postgres:qwerty@host.docker.internal/postgres?SSLMode=disable" up

migration-down:
	goose -dir ./schema postgres "postgres://postgres:qwerty@172.18.0.2:5432/postgres?SSLMode=disable" down

test:

test-coverage:
	protoc --go_out=/home/mu4a/GolandProjects/preproj/internal/handler/grpcapi/gen --go-grpc_out=/home/mu4a/GolandProjects/preproj/internal/handler/grpcapi/gen .proto/product.proto
test:
	protoc --go_out=/home/mu4a/GolandProjects/preproj/internal/handler/grpcapi/gen --go-grpc_out=/home/mu4a/GolandProjects/preproj/internal/handler/grpcapi/gen ./user.proto
