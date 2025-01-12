package filemanager

import (
	"io/fs"
	"os"
	"path/filepath"
)

// FileInfo represents information about a file
type FileInfo struct {
	Name    string
	Size    int64
	Mode    fs.FileMode
	ModTime string
	IsDir   bool
}

// FileManager handles file operations
type FileManager struct {
	rootPath string
}

// NewFileManager creates a new file manager instance
func NewFileManager(rootPath string) *FileManager {
	return &FileManager{
		rootPath: rootPath,
	}
}

// List returns a list of files in the given directory
func (fm *FileManager) List(path string) ([]FileInfo, error) {
	fullPath := filepath.Join(fm.rootPath, path)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var files []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		files = append(files, FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime().String(),
			IsDir:   info.IsDir(),
		})
	}

	return files, nil
}

// CreateDir creates a new directory
func (fm *FileManager) CreateDir(path string) error {
	fullPath := filepath.Join(fm.rootPath, path)
	return os.MkdirAll(fullPath, 0755)
}

// Delete removes a file or directory
func (fm *FileManager) Delete(path string) error {
	fullPath := filepath.Join(fm.rootPath, path)
	return os.RemoveAll(fullPath)
}

// Read reads the contents of a file
func (fm *FileManager) Read(path string) ([]byte, error) {
	fullPath := filepath.Join(fm.rootPath, path)
	return os.ReadFile(fullPath)
}

// Write writes content to a file
func (fm *FileManager) Write(path string, content []byte) error {
	fullPath := filepath.Join(fm.rootPath, path)
	return os.WriteFile(fullPath, content, 0644)
}
