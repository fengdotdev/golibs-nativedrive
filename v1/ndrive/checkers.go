package ndrive

import (
	"errors"
	"fmt"
	"os"

	"github.com/fengdotdev/golibs-helperfuncs/unique"
)

var (
	ErrorDriveNotFound      = errors.New("drive not found")
	ErrorDriveNotAccessible = errors.New("drive not accessible")
	ErrorWriteFileFailed    = errors.New("write file failed")
	ErrorReadFileFailed     = errors.New("read file failed")
	ErrorDeleteFileFailed   = errors.New("delete file failed")
	ErrorCreateFileFailed   = errors.New("create file failed")
	ErrorCreateFolderFailed = errors.New("create folder failed")
	ErrorDeleteFolderFailed = errors.New("delete folder failed")
	ErrorMoveFileFailed     = errors.New("move file failed")
	ErrorMoveFolderFailed   = errors.New("move folder failed")
	ErrorCopyFileFailed     = errors.New("copy file failed")
	ErrorCopyFolderFailed   = errors.New("copy folder failed")
	ErrorRenameFileFailed   = errors.New("rename file failed")
	ErrorNotPermission      = errors.New("not permission")
)

// CheckDrive checks if the drive is valid and accessible.
// It returns an error if the drive is not valid or not accessible.


func CheckDrive(workinPath string) error {

	err := CheckDriveAccess(workinPath)
	if err != nil {
		return err
	}
	err = CheckFileCRUD(workinPath)
	if err != nil {
		return err
	}
	err = CheckFolderCRUD(workinPath)
	if err != nil {
		return err
	}

	return nil
}

// CheckDriveAccess checks if the drive is valid and accessible.
// It returns an error if the drive is not valid or not accessible.
func CheckDriveAccess(workinPath string) error {

	os.Chdir(workinPath)
	// check if drive is accessible
	_, err := os.Stat(workinPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrorDriveNotFound
		}
		if os.IsPermission(err) {
			return ErrorDriveNotAccessible
		}
	}

	return nil
}

// CheckFileCRUD checks if the file is valid and accessible.
// It returns an error if the file is not valid or not accessible.
func CheckFileCRUD(workingPath string) error {
	dummy := "hello world"
	dummydata := []byte(dummy)
	id := unique.RamdomUUID()
	filename := fmt.Sprintf("testfile-%s.txt", id)
	// create file
	file, err := os.Create(filename)
	if err != nil {
		return ErrorCreateFileFailed
	}
	defer os.Remove(filename)
	defer file.Close()
	// write to file
	_, err = file.Write(dummydata)
	if err != nil {
		return ErrorWriteFileFailed
	}
	// close file
	err = file.Close()
	if err != nil {
		return ErrorWriteFileFailed
	}
	// open file
	file, err = os.Open(filename)
	if err != nil {
		return ErrorReadFileFailed
	}
	// read from file
	data := make([]byte, len(dummydata))
	_, err = file.Read(data)
	if err != nil {
		return ErrorReadFileFailed
	}

	// check if data is same
	if string(data) != dummy {
		return ErrorReadFileFailed
	}
	// delete file
	err = os.Remove(filename)
	if err != nil {
		return ErrorDeleteFileFailed
	}

	return nil
}


// CheckFolderCRUD checks if the folder is valid and accessible.
// It returns an error if the folder is not valid or not accessible.
func CheckFolderCRUD(workingPath string) error {

	// create folder

	id := unique.RamdomUUID()
	foldername := fmt.Sprintf("testfolder-%s", id)
	err := os.Mkdir(foldername, 0755)
	if err != nil {
		return ErrorCreateFolderFailed
	}
	// write in folder
	err = CheckFileCRUD(foldername)
	if err != nil {
		return ErrorNotPermission
	}

	// move folder
	err = os.Rename(foldername, foldername+"-moved")
	if err != nil {
		return ErrorMoveFolderFailed
	}

	// copy folder
	err = os.Rename(foldername+"-moved", foldername)
	if err != nil {
		return ErrorCopyFolderFailed
	}

	// delete folder
	err = os.RemoveAll(foldername)
	if err != nil {
		return ErrorDeleteFolderFailed
	}

	// deleted copy folder
	err = os.RemoveAll(foldername + "-moved")
	if err != nil {
		return ErrorDeleteFolderFailed
	}

	return nil
}
