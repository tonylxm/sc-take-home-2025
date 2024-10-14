package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllChildFolders(t *testing.T) {
	sampleData := []folder.Folder{
		{Name: "alpha", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha"},
		{Name: "bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.delta"},
		{Name: "echo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "echo"},
		{Name: "foxtrot", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")), Paths: "foxtrot"},
	}

	f := folder.NewDriver(sampleData)

	tests := []struct {
		name      string
		orgID     uuid.UUID
		folder    string
		expected  []folder.Folder
		expectErr bool
	}{
		{
			name:   "Get all children of alpha",
			orgID:  uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folder: "alpha",
			expected: []folder.Folder{
				{Name: "bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.delta"},
			},
			expectErr: false,
		},
		{
			name:   "Get all children of bravo",
			orgID:  uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folder: "bravo",
			expected: []folder.Folder{
				{Name: "charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "alpha.bravo.charlie"},
			},
			expectErr: false,
		},
		{
			name:      "Folder not found",
			orgID:     uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folder:    "invalid_folder",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "No children for folder",
			orgID:     uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folder:    "echo",
			expected:  []folder.Folder{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := f.GetAllChildFolders(tt.orgID, tt.folder)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

