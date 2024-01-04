package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Folder struct {
	name    string
	folders []*Folder
	files   []string
}

//Folders := make(map[string][]string)

type ByName []*Folder

func (n ByName) Len() int           { return len(n) }
func (n ByName) Less(i, j int) bool { return n[i].name < n[j].name }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func NewFolder(name string, files []string, folders []*Folder) *Folder {
	return &Folder{name: name, files: files, folders: folders}
}

func (f *Folder) GetFolder(key string) *Folder {
	return nil
}

func (f *Folder) InsertFolder(path []string) {
	for _, val := range path {
		if len(path) != 1 {
			if f.folders != nil {
				for _, folder := range f.folders {
					if folder.name == val {
						folder.InsertFolder(path[1:])
						return
					}
				}
			} else {
				folder := NewFolder(val, nil, nil)
				f.folders = append(f.folders, folder)
				folder.InsertFolder(path[1:])
				return
			}
		} else {
			if strings.Contains(val, ".txt") {
				f.files = append(f.files, val)
				sort.Strings(f.files)
			} else {
				folder := NewFolder(val, nil, nil)
				f.folders = append(f.folders, folder)
				sort.Sort(ByName(f.folders))
			}
		}
	}
}

func (f *Folder) SearchDuplicate() {

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := NewFolder("a", nil, nil)
	for scanner.Scan() {
		var path []string
		path = strings.Split(scanner.Text(), "/")
		fmt.Println(path)
		root.InsertFolder(path[1:])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
