package app

import (
	"context"
	"log"

	"github.com/garyhopkins/adminex/repository/ent/entgen"
	_ "github.com/mattn/go-sqlite3"
)

type AppSettings struct {
	EntClient *entgen.Client
}

var CoreApp *AppSettings

func setupApp() {
	ctx := context.Background()

	// set up a db client. testing for now
	client := NewTestDbClient()
	RunDbMigrations(ctx, client)

	a := AppSettings{
		EntClient: client,
	}
	CoreApp = &a
}

func NewTestDbClient() *entgen.Client {
	// client, err := entgen.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := entgen.Open("sqlite3", "file:sqlite.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func RunDbMigrations(ctx context.Context, client *entgen.Client) {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func GetApp() *AppSettings {
	if CoreApp == nil {
		setupApp()
	}
	return CoreApp
}
