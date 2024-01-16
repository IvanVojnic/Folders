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
	visited bool
	path    string
}

type ArrayFolder struct {
	content []string
	paths   []string
}

func NewArrayFolder(files []string, folders []*Folder, path string) *ArrayFolder {
	content := make([]string, 0)
	for _, file := range files {
		content = append(content, file)
	}
	for _, currentFolder := range folders {
		content = append(content, currentFolder.name)
	}
	var paths []string
	paths = append(paths, path)
	return &ArrayFolder{content: content, paths: paths}
}

func NewFolder(name string, files []string, folders []*Folder, path string) *Folder {
	return &Folder{name: name, files: files, folders: folders, path: path}
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
			}

			folder := NewFolder(val, nil, nil, fmt.Sprintf("%s%s/", f.path, val))
			f.folders = append(f.folders, folder)
			folder.InsertFolder(path[1:])
			return

		}
		if strings.Contains(val, ".txt") {
			f.files = append(f.files, val)
		} else {
			isExists := false
			for _, childFolder := range f.folders {
				if val == childFolder.name {
					isExists = true
				}
			}
			if !isExists {
				folder := NewFolder(val, nil, nil, fmt.Sprintf("%s%s/", f.path, val))
				f.folders = append(f.folders, folder)
			}
		}
	}
}

func SearchDuplicate(folder *Folder) {
	queue := make([]*Folder, 0)
	var duplFolders []ArrayFolder
	queue = append(queue, folder)
	folder.visited = true
	duplFolders = append(duplFolders, *NewArrayFolder(folder.files, folder.folders, folder.path))
	for len(queue) > 0 {
		v := queue[0]
		if !v.visited {
			foldersNames := make([]string, 0, len(v.folders))
			for _, currentFolder := range v.folders {
				foldersNames = append(foldersNames, currentFolder.name)
			}
			isDupl := false
			for i, arrayFolders := range duplFolders {
				if arrayFolders.CompareArrays(v.files, foldersNames) {
					duplFolders[i].paths = append(duplFolders[i].paths, v.path)
					isDupl = true
				}
			}
			if !isDupl {
				duplFolders = append(duplFolders, *NewArrayFolder(v.files, v.folders, v.path))
			}
			v.visited = true
		}
		for _, currentFolder := range v.folders {
			queue = append(queue, currentFolder)
		}
		queue = queue[1:]
	}
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	for _, duplFolder := range duplFolders {
		content := fmt.Sprintf("\ncontent - %s", strings.Join(duplFolder.content, ","))
		file.WriteString(content)

		for _, path := range duplFolder.paths {
			pathString := fmt.Sprintf("\n path - %s", path)
			file.WriteString(pathString)
		}
		file.WriteString("\n")
	}
}

func (a *ArrayFolder) CompareArrays(files []string, folders []string) bool {
	for _, file := range files {
		folders = append(folders, file)
	}

	sort.Strings(folders)
	sort.Strings(a.content)
	matches := 0

	for _, contentValue := range a.content {
		index := sort.Search(len(folders), func(i int) bool {
			return folders[i] >= contentValue
		})

		if index < len(folders) && folders[index] == contentValue {
			matches++
		}
	}

	var similarityPercent float64
	foldersLen := len(folders)
	AContentLen := len(a.content)

	if foldersLen > AContentLen {
		similarityPercent = float64(matches) / float64(foldersLen) * 100
	} else {
		similarityPercent = float64(matches) / float64(AContentLen) * 100
	}

	if similarityPercent >= 90 {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := NewFolder("root", nil, nil, "root/")
	for scanner.Scan() {
		var path []string
		path = strings.Split(scanner.Text(), "/")
		pathWithRoot := append([]string{"root"}, path...)
		root.InsertFolder(pathWithRoot[1:])
	}
	SearchDuplicate(root)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
