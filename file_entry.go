package chunker

import (
	"path/filepath"
	"strings"
)

// FileInfo struct to hold details about each file
type FileInfo struct {
	Filename string
	Size     int64
}

// IsDotFile checks if the file is a dot file
func (f *FileInfo) IsDotFile() bool {
	return strings.HasPrefix(filepath.Base(f.Filename), ".")
}

// SizeInKB returns the file size in kilobytes
func (f *FileInfo) SizeInKB() float64 {
	return float64(f.Size) / 1024
}

// SizeInMB returns the file size in megabytes
func (f *FileInfo) SizeInMB() float64 {
	return float64(f.Size) / (1024 * 1024)
}
