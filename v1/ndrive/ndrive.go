package ndrive

import (
	"github.com/fengdotdev/golibs-nativedrive/ntrait"
	"os"
)

var _ SudoDrive = (*NDrive)(nil)

type NDrive struct {
	workingDir string
}

// CreateFileEmpty implements ntrait.SudoDrive.
func (n *NDrive) CreateFileEmpty(name string) error {
	panic("unimplemented")
}

// CreateFileWithContent implements ntrait.SudoDrive.
func (n *NDrive) CreateFileWithContent(name string, content []byte) error {
	panic("unimplemented")
}

// CreateFolder implements ntrait.SudoDrive.
func (n *NDrive) CreateFolder(name string) error {
	panic("unimplemented")
}

// DeleteFile implements ntrait.SudoDrive.
func (n *NDrive) DeleteFile(name string) error {
	panic("unimplemented")
}

// DeleteFolder implements ntrait.SudoDrive.
func (n *NDrive) DeleteFolder(name string) error {
	panic("unimplemented")
}

// GetFile implements ntrait.SudoDrive.
func (n *NDrive) GetFile(name string) (ntrait.File[T], error) {
	panic("unimplemented")
}

// GetFolder implements ntrait.SudoDrive.
func (n *NDrive) GetFolder(name string) (ntrait.Folder[F, E, T], error) {
	panic("unimplemented")
}

// GetList implements ntrait.SudoDrive.
func (n *NDrive) GetList() ([]ntrait.Element[E, F, T], error) {
	panic("unimplemented")
}

// UpdateFile implements ntrait.SudoDrive.
func (n *NDrive) UpdateFile(name string, newName string) error {
	panic("unimplemented")
}

// UpdateFileContent implements ntrait.SudoDrive.
func (n *NDrive) UpdateFileContent(name string, content []byte) error {
	panic("unimplemented")
}

// UpdateFolder implements ntrait.SudoDrive.
func (n *NDrive) UpdateFolder(name string, newName string) error {
	panic("unimplemented")
}

func NewNDrive(workingDir string) (*NDrive, error) {
	err := CheckDrive(workingDir)
	if err != nil {
		return nil, err
	}

	return &NDrive{
		workingDir: workingDir,
	}, nil
}

func NewNDriveWithFolder(workingDir string) (*NDrive, error) {
	CreateFolder(workingDir)

	return &NDrive{
		workingDir: workingDir,
	}, nil
}

func CreateFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (n *NDrive) GetName() string {
	os.Chdir(n.workingDir)

	// check if drive is accessible
	stats, err := os.Stat(n.workingDir)
	if err != nil {
		panic(err)
	}

	name := stats.Name()

	return name
}
