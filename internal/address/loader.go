package address

import (
	"errors"
	"fmt"
	"os"
)

type JsonZipDatabaseLoader struct {
	DbPath string
}

func NewJsonDatabaseLoader(dbPath string) *JsonZipDatabaseLoader {
	return &JsonZipDatabaseLoader{
		DbPath: dbPath,
	}
}

func (loader JsonZipDatabaseLoader) Load() (zipDatabase ZipDatabase, err0 error) {
	dbFile, err := os.Open(loader.DbPath)
	if err != nil {
		return nil, fmt.Errorf("open file '%s': %w", loader.DbPath, err)
	}

	defer func() {
		if err := dbFile.Close(); err != nil {
			err0 = errors.Join(fmt.Errorf("close file '%s': %w", loader.DbPath, err), err0)
		}
	}()

	zipDatabase, err = FromJson(dbFile)
	if err != nil {
		return nil, fmt.Errorf("parsing file '%s': %w", loader.DbPath, err)
	}

	return zipDatabase, nil
}
