package folder

import (
	"errors"
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if len(f.folders) == 0 {
		return nil, errors.New("no folders found for the given organisation")
	}

	var srcFolder *Folder
	var dstFolder *Folder

	for _, folder := range f.folders {
		if folder.Name == name {
			srcFolder = &folder
		}
		if folder.Name == dst {
			dstFolder = &folder
		}
	}

	if srcFolder == nil {
		return nil, fmt.Errorf("folder %s not found", name)
	}

	if dstFolder == nil {
		return nil, fmt.Errorf("folder %s not found", dst)
	}

	if dstFolder.OrgId != srcFolder.OrgId {
		return nil, errors.New("folders belong to different organisations")
	}

	if srcFolder.Name == dst || strings.HasPrefix(srcFolder.Paths, dstFolder.Paths) {
		return nil, errors.New("cannot move a folder to itself or to one of its children")
	}

	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("cannot move a folder under its own descendant")
	}

	newParentPath := dstFolder.Paths + "." + srcFolder.Name
	oldParentPath := srcFolder.Paths

	updatedFolders := []Folder{}
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, oldParentPath) {
			newPath := strings.Replace(folder.Paths, oldParentPath, newParentPath, 1)
			folder.Paths = newPath
		}
		updatedFolders = append(updatedFolders, folder)
	}

	srcFolder.Paths = newParentPath

	return updatedFolders, nil
}
