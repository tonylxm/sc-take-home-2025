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

	for _, f := range f.folders {
		if f.Name == name {
			srcFolder = &f
		}
		if f.Name == dst {
			dstFolder = &f
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

	if strings.HasPrefix(srcFolder.Paths, dstFolder.Paths) {
		return nil, errors.New("cannot move a folder to one of its children")
	}

	newParentPath := dstFolder.Paths + "." + srcFolder.Name
	oldParentPath := srcFolder.Paths

	updatedFolders := []Folder{}

	for _, f := range f.folders {
		if strings.HasPrefix(f.Paths, oldParentPath) {
			newPath := strings.Replace(f.Paths, oldParentPath, newParentPath, 1)
			f.Paths = newPath
			updatedFolders = append(updatedFolders, f)
		}
	}

	return updatedFolders, nil
}
