package ndrive

// GetWorkingDir implements ntrait.SudoDrive.
func (n *NDrive) GetWorkingDir() string {
	return n.workingDir
}

// SetWorkingDir implements ntrait.SudoDrive.
func (n *NDrive) SetWorkingDir(path string) error {

	err := CheckDrive(path)
	if err != nil {
		return err
	}

	n.workingDir = path
	return nil
}
