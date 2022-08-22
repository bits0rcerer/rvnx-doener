package test

import (
	"log"
	"os"
	"testing"
)

func CommonTestMain(m *testing.M) {
	SetTestENVs()
    exitVal := m.Run()
    os.Exit(exitVal)
}

func SetTestENVs() {
	panicOnError(os.Setenv("TEST_DATABASE_URL", "postgresql://pg:pg@localhost:5433/pg"))
    panicOnError(os.Setenv("DEBUG", "TRUE"))
    panicOnError(os.Setenv("SESSION_ENCRYPTION_SECRET", "DO-NOT-USE-IN-PRODUCTION"))
    panicOnError(os.Setenv("SESSION_SECRET", "DO-NOT-USE-IN-PRODUCTION"))
    panicOnError(os.Setenv("TWITCH_CLIENT_ID", "DO-NOT-USE-IN-PRODUCTION"))
    panicOnError(os.Setenv("TWITCH_CLIENT_SECRET", "DO-NOT-USE-IN-PRODUCTION"))
}

func panicOnError(err error)  {
	if err != nil {
		log.Panicln(err)
	}
}