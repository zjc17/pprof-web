package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateFileID(t *testing.T) {
	fileID := GenerateFileID()
	assert.True(t, len(fileID) == fileIDLength)
}
