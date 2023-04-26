package files

import (
	"fmt"
	"github.com/zjc17/pprof-web/pkg/random"
	"os"
	"path"
)

const (
	fileIDLength = 8
)

func GenerateFileID() string {
	return random.GenerateRandomDigitalWithUpperCaseLetterCode(fileIDLength)
}

// GetFilePathByFileID generate file path
// TODO consider the file conflict
func GetFilePathByFileID(fileID string) string {
	return path.Join(os.TempDir(), fmt.Sprintf("pprof-web-%s", fileID))
}
