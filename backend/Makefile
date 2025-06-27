air:
	go build -o ./tmp/server cmd/server/main.go

live:
	go run github.com/air-verse/air@latest -c .air.toml

generate:
	rm -rf repository/ent/entgen/
	go generate ./...
