package data

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"errors"
	"log"
	"net/url"
	"os"
	"rvnx_doener_service/ent"
	"strings"

	_ "github.com/lib/pq"
)

const DatabaseURLKey = "DATABASE_URL"

func OpenPostgres(sslMode string) (closeDB func() error, client *ent.Client, err error) {
	databaseURL := os.Getenv(DatabaseURLKey)
	if databaseURL == "" {
		return nil, nil, errors.New(DatabaseURLKey + " is not defined")
	}

	return OpenPostgresWithURL(databaseURL, sslMode)
}

func OpenPostgresWithURL(databaseURL string, sslMode string) (closeDB func() error, client *ent.Client, err error) {
	dsn, err := url.Parse(databaseURL)
	if err != nil || dsn == nil {
		return nil, nil, errors.New("unable to parse " + DatabaseURLKey)
	}

	pw, _ := dsn.User.Password()
	return OpenPostgresWithConnectionString(BuildPostgresConnectionString(dsn.Hostname(), dsn.Port(), dsn.User.Username(), pw, strings.TrimLeft(dsn.Path, "/"), sslMode))
}

func OpenPostgresWithConnectionString(connectionString string) (closeDB func() error, client *ent.Client, err error) {
	client, err = ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background(),
		schema.WithGlobalUniqueID(true),
	); err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}

	return client.Close, client, err
}

func BuildPostgresConnectionString(host, port, user, password, dbname, sslMode string) string {
	return "host=" + host + " port=" + port + " user=" + user +
		" dbname=" + dbname + " password=" + password + " sslmode=" + sslMode
}
