package chunker

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestListFilesRecursively(t *testing.T) {
	// Setup test directory
	testDir := "./testdir"
	require.NoError(t, os.MkdirAll(testDir, os.ModePerm))
	defer os.RemoveAll(testDir) // Clean up after the test

	// Create dummy files in the test directory
	_, err := os.Create(filepath.Join(testDir, "file1.txt"))
	require.NoError(t, err)
	_, err = os.Create(filepath.Join(testDir, ".hiddenfile"))
	require.NoError(t, err)
	c := NewChunker()
	files, err := c.ListFilesRecursively(testDir)
	require.NoError(t, err)
	assert.Len(t, files, 2, "There should be two files listed in the directory")

	// Check if files are correctly identified
	for _, file := range files {
		assert.True(t, file.Size >= 0, "File size should be non-negative")
		if file.IsDotFile() {
			assert.Contains(t, file.Filename, ".hiddenfile", "Dot file should be '.hiddenfile'")
		} else {
			assert.Contains(t, file.Filename, "file1.txt", "Non-dot file should be 'file1.txt'")
		}
	}
}
