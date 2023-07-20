package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func containsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./copyfiles.exe path/to/srcCode path/to/iconSrc path/to/destination *optional: resolutions separated with ,*")
		return
	}
	iconSrc := os.Args[2]
	_, err := os.Stat(iconSrc)
	if err != nil {
		fmt.Printf("The folder '%s' does not exist.\n", iconSrc)
		return
	}
	srcCode := os.Args[1]
	_, err = os.Stat(srcCode)
	if err != nil {
		fmt.Printf("The folder '%s' does not exist.\n", srcCode)
		return
	}
	destination := os.Args[3]
	_, err = os.Stat(destination)
	if err != nil {
		err := os.MkdirAll(destination, 0755)
		if err != nil {
			fmt.Printf("error with creating destination folder %s", err)
			return
		}
	}

	iconFiles := getIcons(iconSrc, "")
	neededIcons := srcCodeFind(srcCode)
	if len(os.Args) < 5 {
		fmt.Print("no reso checks")
		for _, file := range iconFiles {
			fileName := filepath.Base(file)
			fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			if containsString(neededIcons, fileNameWithoutExtension) {
				absPath, _ := filepath.Abs(file)
				relativePath, _ := filepath.Rel(iconSrc, absPath)
				newPath := destination + "/" + relativePath
				err = copyFile(file, newPath)
				if err != nil {
					fmt.Println("error happened")
					panic(err)
				}
			}
		}
	} else {
		resolutionsNeeded := os.Args[4]
		resolutions := neededResos(resolutionsNeeded)
		for _, file := range iconFiles {
			fileName := filepath.Base(file)
			fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			if containsString(neededIcons, fileNameWithoutExtension) && containsString(resolutions, getResolution(file)) {
				absPath, _ := filepath.Abs(file)
				relativePath, _ := filepath.Rel(iconSrc, absPath)
				newPath := destination + "/" + relativePath
				err = copyFile(file, newPath)
				if err != nil {
					fmt.Println("error happened")
					panic(err)
				}
			}
		}
	}
	themes := getThemes(iconSrc)
	for _, theme := range themes {
		absPath, _ := filepath.Abs(theme)
		relativePath, _ := filepath.Rel(iconSrc, absPath)
		newPath := destination + "/" + relativePath
		err = copyFile(theme, newPath)
		if err != nil {
			fmt.Println("error happened")
			panic(err)
		}
	}
	println("Files copied successfully!")
}
