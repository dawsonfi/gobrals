package address

import (
	"errors"
	"os"
	"testing"
)

func Test_FromJson(t *testing.T) {
	t.TempDir()
}

func createTempFile(tempDir string, tempFile string) (file *os.File, err0 error) {
	file, err := os.CreateTemp(tempDir, tempFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errors.Join(err, file.Close())
	}()

	return file, nil
}
