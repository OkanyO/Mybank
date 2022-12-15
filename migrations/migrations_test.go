package migrations

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Test_createAccounts(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createAccounts()
		})
	}
}

func TestMigrate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Migrate()
		})
	}
}
