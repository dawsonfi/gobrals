package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewJsonDatabaseLoader(t *testing.T) {
	tests := []struct {
		name    string
		dbPath  string
		succeed bool
	}{
		{
			name:    "create instance with valid path",
			dbPath:  "data/br_address_database.json",
			succeed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewJsonDatabaseLoader(tt.dbPath))
		})
	}
}
