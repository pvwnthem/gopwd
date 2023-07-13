package util

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	return homeDir
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
func CopyFile(src, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	// Copy the contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}

	return nil
}
func RemoveDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func CreateTempFileFromBytes(content []byte) *os.File {
	tmpfile, _ := os.CreateTemp("", "tempfile")
	tmpfile.Write(content)
	tmpfile.Close()

	_ = os.Chmod(tmpfile.Name(), 0644) // Set file permission to read-only

	return tmpfile
}

// ReadBytesFromFile reads the content of a file and returns it as a byte slice.
func ReadBytesFromFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadDirectory(path string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile(path string, data string) error {
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

func WriteBytesToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

func PrintDirectoryTree(dirPath string, indent string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		if entry.IsDir() {
			if i == len(entries)-1 {
				fmt.Printf("%s└── %s\n", indent, entry.Name())
			} else {
				fmt.Printf("%s├── %s\n", indent, entry.Name())
			}
		}
		if entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			err := PrintDirectoryTree(subDirPath, indent+"│   ")
			if err != nil {
				fmt.Printf("Error printing subdirectory '%s': %v\n", subDirPath, err)
			}
		}
	}

	return nil
}

// Utility functions to replace ioutil package functions

func createTempFile(dir, prefix string) (*os.File, error) {
	tmpfile, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return nil, err
	}
	return tmpfile, nil
}

func writeFile(filename string, data []byte, perm os.FileMode) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
