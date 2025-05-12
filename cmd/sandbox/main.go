package main

import (
	"fmt"
	"os"

	"github.com/fengdotdev/golibs-nativedrive/v1/ndrive"
)

func main() {
	workingDirRaw, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	workingDir := workingDirRaw + "/testfolder"

	drive, err := ndrive.NewNDriveWithFolder(workingDir)
	if err != nil {
		panic(err)
	}

	fmt.Println("Drive Name:", drive.GetName())

}
