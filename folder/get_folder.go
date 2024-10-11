package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

// GetFoldersByOrgID returns all folders for a given orgID
func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

// GetAllChildFolders returns all child folders of a folder name for a given orgID
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.folders

	if len(folders) == 0 {
		return nil, errors.New("No folders found for the given organisation")
	}

	var parentFolder *Folder
	for _, f := range folders {
		if f.Name == name {
			parentFolder = &f
			break
		}
	}

	if parentFolder == nil {
		return nil, errors.New("No folder found with the given name")
	}

	childFolders := []Folder{}
	for _, f := range folders {
		if strings.HasPrefix(f.Paths, parentFolder.Paths) && f.Paths != parentFolder.Paths {
			childFolders = append(childFolders, f)
		}
	}

	return childFolders, nil
}
