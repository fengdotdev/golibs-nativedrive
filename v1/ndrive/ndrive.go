package ndrive

import "os"

var _ SudoDrive = (*NDrive)(nil)

type NDrive struct {
	workingDir string
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
