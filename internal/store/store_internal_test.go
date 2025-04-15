package store

import (
	"fmt"
	"go-rest-api/internal/config"
	"os"
	"strings"
	"testing"
)

var databaseURL string

func TestMain (m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=darki password=1234 dbname=my_project sslmode=disable"
	}
	
	os.Exit(m.Run())
}

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
    t.Helper()
    
    config := config.NewStoreConfig()
    config.DatabaseURL = databaseURL
    
    s := New(config)
    err := s.Open()
    if err != nil {
        t.Fatal(err)
    }
    
    return s, func(tables ...string) {
        if len(tables) > 0 {
            _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
            if err != nil {
                t.Fatal(err)
            }
            
            s.Close()
        }
    }
}
