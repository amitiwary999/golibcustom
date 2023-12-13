package filereader

import (
	"fmt"
	"os"
)

func GetAllFilesName() {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		fmt.Print(homeDir)
		dirList, err := os.ReadDir(homeDir)
		if err == nil {
			for _, dir := range dirList {
				fmt.Printf("------name of sub dir %v--------\n", dir)
				subDir := fmt.Sprintf("%v/%v", homeDir, dir.Name())
				subDirList, errSubDir := os.ReadDir(subDir)
				if errSubDir == nil {
					for _, subDir := range subDirList {
						fmt.Printf("sub dir name %v\n", subDir)
					}
				} else {
					fmt.Printf("error printing sub dir of all dir of home %v\n", errSubDir)
				}
			}
		} else {
			fmt.Printf("error in reading directory %v\n", err)
		}
	} else {
		fmt.Printf("error in getting home dir %v\n", err)
	}
}
