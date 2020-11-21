package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)

	if err == nil {
		return true
	}

	return false
}

// ReadFile: Reads... a file
func ReadFile(path string) []byte {
	if EmptyString(path) {
		Print("Path indefinido")

		return nil
	}

	res, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte("")
	}

	return res
}

func SeparateFileByLine(file []byte) ([]string, int) {
	lines := make([]string, 1)
	start := 0
	end := 0
	for pos, b := range file {
		if BtsSingle(b) == "\n" {
			end = pos
			lines = append(lines, Bts(file[start:end]))
			start = pos + 1
		}
	}

	return lines, end
}

func createFile(path string) bool {
	if FileExists(path) {
		Print("File exists!")

		return false
	}

	_, err := os.Create(path)
	Panic(err)

	return true
}

func WriteBytesToFile(file string, text []byte) {
	if !FileExists(file) {
		createFile(file)
	}

	err := ioutil.WriteFile(file, text, 0644)
	Panic(err)
}

func WriteToFile(file string, text string) {
	if !FileExists(file) {
		createFile(file)
	}

	err := ioutil.WriteFile(file, []byte(text), 0644)
	Panic(err)
}

func AppendToFile(file string, text []string) int {
	if !FileExists(file) {
		Print("File don't exist")

		return 0
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	StayCalm(err)

	defer StayCalm(f.Close())
	var num int
	for _, line := range text {
		if _, e := f.WriteString(line + "\n"); err != nil {
			log.Println(e)
			num++
		}
	}

	return num
}

func ReadFilesInDir(path string) []string {
	c, err := ioutil.ReadDir(path)
	StayCalm(err)

	var files []string
	for _, entry := range c {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files
}

type FileInfo struct {
	Name     string
	Size     int64
	Dir      bool
	Modified time.Time
}

func ReadDir(path string) []FileInfo {
	c, err := ioutil.ReadDir(path)
	StayCalm(err)

	var files []FileInfo
	for _, entry := range c {
		entry.Mode()
		files = append(files, FileInfo{Name: entry.Name(), Size: entry.Size(), Dir: entry.IsDir(), Modified: entry.ModTime()})
	}

	return files
}
