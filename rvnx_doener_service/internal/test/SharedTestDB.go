package test

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/internal/data"
	"strconv"
	"strings"
)

const TestDatabaseURLKey = "TEST_DATABASE_URL"

func OpenSharedTestDB() (client *ent.Client, cleanUp func(), err error) {
	dbUrl := os.Getenv(TestDatabaseURLKey)
	if dbUrl == "" {
		log.Panicln(TestDatabaseURLKey + " is not defined")
	}

	cfg, err := pgx.ParseConfig(dbUrl)
	if err != nil {
		log.Printf("unable to parse %s(=%s): %v\n", TestDatabaseURLKey, dbUrl, err)
		return nil, nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		log.Printf("unable to connect to test database: %v\n", err)
		return nil, nil, err
	}
	defer func() {
		_ = conn.Close(context.Background())
	}()

	dbUid, err := uuid.NewV4()
	if err != nil {
		log.Printf("unable to generate dbUid: %v\n", err)
		return nil, nil, err
	}

	databaseName := fmt.Sprintf("rvnx_doener_test_%s", strings.ReplaceAll(dbUid.String(), "-", ""))
	_, err = conn.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s;", databaseName))
	if err != nil {
		log.Printf("unable to create new database (%s): %v\n", databaseName, err)
		return nil, nil, err
	}

	closeDB, client, err := data.OpenPostgresWithConnectionString(data.BuildPostgresConnectionString(cfg.Host, strconv.Itoa(int(cfg.Port)), cfg.User, cfg.Password, databaseName, "disable"))
	if err != nil {
		log.Printf("unable to connect to test database instance (%s): %v\n", databaseName, err)
		return nil, nil, err
	}

	return client, func() {
		err := closeDB()
		if err != nil {
			log.Printf("unable to close test database: %v\n", err)
			// do not return
		}

		conn, err := pgx.ConnectConfig(context.Background(), cfg)
		if err != nil {
			log.Printf("unable to connect to test database: %v\n", err)
			return
		}
		defer func() {
			_ = conn.Close(context.Background())
		}()

		_, err = conn.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %s;", databaseName))
		if err != nil {
			log.Printf("unable to cleanup database (%s): %v\n", databaseName, err)
			return
		}
	}, nil
}
