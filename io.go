package gears

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PathGenAsDate() (s string, err error) {
	b := time.Now()
	s1 := fmt.Sprintf("%d", b.Year())
	s2 := fmt.Sprintf("%02d%02d", b.Month(), b.Day())
	s = filepath.Join(s1, s2)
	if !Exists(s) {
		err = os.MkdirAll(s, 0755)
		if err != nil {
			fmt.Println("[-] gears.PathGenAsDate() error!")
			log.Fatal(err)
			return
		}
	}
	return
}

func MakeDirAll(path string) {
	if !Exists(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("[-] gears.MakeDirAll() error")
			log.Fatal(err)
		}
	}
}

func FileRemoveRoutine(root string) error {
	a := time.Now().AddDate(0, 0, -2)
	b := fmt.Sprintf("[%02d%02d]", a.Month(), a.Day())
	err := os.Remove(filepath.Join(root, b))
	if err != nil {
		fmt.Println("[-] gears.RemoveRoutine() error")
		return err
	}

	return nil
}

func RemoveRoutine(root string) error {
	a := time.Now().AddDate(0, 0, -2)
	b := fmt.Sprintf("%02d%02d", a.Month(), a.Day())
	err := os.RemoveAll(filepath.Join(root, b))
	if err != nil {
		fmt.Println("[-] gears.RemoveRoutine() error")
		return err
	}

	return nil
}

// Exists judge the path, return true while exist, whatever it's a file or folder
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func FileCodeDetector(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("[-] gears.FileCodeDetector() error")
		log.Fatal(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	fdetect := StrDetector(string(fd))

	return fdetect
}

// GetInput get string from os.Stdin filted \r and \n
func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	return input
}

// GetUnPrefixedFiles return a files slice from walk over the folder path and filtered by without prefix string.
func GetUnPrefixedFiles(folder, prefix string) (files []string, err error) {
	if !Exists(folder) {
		fmt.Printf("\n[-] Folder's not exist!\n%s\n", folder)
		return
	}
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && !strings.HasPrefix(info.Name(), prefix) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", folder, err)
		return
	}

	return
}

// GetPrefixedFiles return a files slice from walk over the folder path and filtered by prefix string.
func GetPrefixedFiles(folder, prefix string) (files []string, err error) {
	if !Exists(folder) {
		fmt.Printf("\n[-] Folder's not exist!\n%s\n", folder)
		return
	}
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasPrefix(info.Name(), prefix) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", folder, err)
		return
	}

	return
}
