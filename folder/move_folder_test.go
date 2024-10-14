package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	sampleData := []folder.Folder{
		{Name: "alpha", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha"},
		{Name: "bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.delta"},
		{Name: "echo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.delta.echo"},
		{Name: "foxtrot", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")), Paths: "foxtrot"},
		{Name: "golf", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "golf"},
	}

	f := folder.NewDriver(sampleData)

	tests := []struct {
		name          string
		folderName    string
		dst           string
		expectedPaths []string
		expectErr     bool
		}{
			{
				name:          "Move bravo under delta",
				folderName:    "bravo",
				dst:           "delta",
				expectedPaths: []string{"alpha", "alpha.delta.bravo", "alpha.delta.bravo.charlie", "alpha.delta", "alpha.delta.echo", "foxtrot", "golf"},
				expectErr:     false,
			},
			{
				name:          "Move non-existent folder",
				folderName:    "nonexistent",
				dst:           "alpha",
				expectedPaths: nil,
				expectErr:     true,
			},
			{
				name:          "Move to non-existent destination",
				folderName:    "bravo",
				dst:           "nonexistent",
				expectedPaths: nil,
				expectErr:     true,
			},
			{
				name:          "Move bravo to itself",
				folderName:    "bravo",
				dst:           "bravo",
				expectedPaths: nil,
				expectErr:     true,
			},
			{
				name:          "Move bravo under charlie",
				folderName:    "bravo",
				dst:           "charlie",
				expectedPaths: nil,
				expectErr:     true,
			},
			{
				name:          "Move bravo to golf (different org)",
				folderName:    "bravo",
				dst:           "foxtrot",
				expectedPaths: nil,
				expectErr:     true,
			},
			{
				name:          "Move to an existing destination",
				folderName:    "bravo",
				dst:           "golf",
				expectedPaths: []string{"alpha", "golf.bravo", "golf.bravo.charlie", "alpha.delta", "alpha.delta.echo", "foxtrot", "golf"},
				expectErr:     false,
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedFolders, err := f.MoveFolder(tt.folderName, tt.dst)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				paths := []string{}
				for _, folder := range updatedFolders {
					paths = append(paths, folder.Paths)
				}
				assert.Equal(t, tt.expectedPaths, paths)
			}
		})
	}
}

