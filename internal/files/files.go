package files

import "github.com/zjc17/pprof-web/pkg/random"

func GenerateFileID() string {
	return random.GenerateRandomDigitalWithUpperCaseLetterCode(8)
}

// GetFilePathByFileID generate file path
// TODO consider the file conflict
func GetFilePathByFileID(fileID string) string {
	return "/tmp/pprof-web-" + fileID
}
