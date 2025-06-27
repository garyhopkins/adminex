package app

import (
	"context"
	"log"

	"github.com/garyhopkins/adminex/repository/ent/entgen"
)

func NewTestDbClient() *entgen.Client {
	client, err := entgen.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
