package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Folder struct {
	name    string
	folders []*Folder
	files   []string
	visited bool
}

type ArrayFolder struct {
	content []string
	paths   []string
}

func NewArrayFolder(content []string, paths []string) *ArrayFolder {
	return &ArrayFolder{content: content, paths: paths}
}

var DuplicateFolders []ArrayFolder

func NewFolder(name string, files []string, folders []*Folder) *Folder {
	return &Folder{name: name, files: files, folders: folders}
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
				//sort.Sort(ByName(f.folders))
				folder.InsertFolder(path[1:])
				return
			}
		} else {
			if strings.Contains(val, ".txt") {
				f.files = append(f.files, val)
				//sort.Strings(f.files)
			} else {
				folder := NewFolder(val, nil, nil)
				f.folders = append(f.folders, folder)
				//sort.Sort(ByName(f.folders))
			}
		}
	}
}

func (f *Folder) SearchDuplicate(folder *Folder) {
	queue := make([]*Folder, 10, 20)
	queue = append(queue, folder)
	folder.visited = true
	content := make([]string, len(folder.files)+len(folder.folders))
	for _, file := range folder.files {
		content = append(content, file)
	}
	for _, currentFolder := range folder.folders {
		content = append(content, currentFolder.name)
	}
	var paths []string
	paths = append(paths, fmt.Sprintf(folder.name, "/"))
	DuplicateFolders = append(DuplicateFolders, *NewArrayFolder(content, paths))
	for len(queue) > 0 {
		v := queue[0]
		for _, folder := range v.folders {
			//if()
		}
	}
}

func CompareArrays(files []string, folders []string, filesFolders []string) bool {

	return false
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

//Folders := make(map[string][]string)
/*type ByName []*Folder

func (n ByName) Len() int           { return len(n) }
func (n ByName) Less(i, j int) bool { return n[i].name < n[j].name }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (f *Folder) GetFolder(key string) *Folder {
	return nil
}
*/
