package chunker

import (
	"io/fs"
	"path/filepath"
)

type Chunker struct{}

func NewChunker() Chunker {
	return Chunker{}
}

// ListFilesRecursively lists all files in a given directory and stores them in FileInfo
func (c *Chunker) ListFilesRecursively(dir string) ([]FileInfo, error) {
	var files []FileInfo
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, FileInfo{Filename: path, Size: info.Size()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
