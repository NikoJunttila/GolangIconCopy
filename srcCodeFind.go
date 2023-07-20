package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func srcCodeFind(folderPath string) []string {
	var icons []string
	filePaths := findPaths(folderPath)

	for _, filePath := range filePaths {
		sourceCode, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filePath, err)
			continue
		}

		lines := strings.Split(string(sourceCode), "\n")

		for _, line := range lines {
			// Look for icons in different formats
			if strings.Contains(line, "QIcon::fromTheme") {
				iconName := strings.Split(strings.Split(line, "::fromTheme(\"")[1], "\")")[0]
				icons = append(icons, iconName)
			}
			if strings.Contains(line, "iconset theme=") {
				iconName := strings.Split(strings.Split(line, "theme=\"")[1], "\"")[0]
				icons = append(icons, iconName)
			}
		}
	}
	return icons
}

func findPaths(folderPath string) []string {
	filePaths := []string{}
	cppPaths, err := filepath.Glob(filepath.Join(folderPath, "*.cpp"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(cppPaths) != 0 {
		for _, path := range cppPaths {
			filePaths = append(filePaths, path)
		}
	}
	uiPaths, err := filepath.Glob(filepath.Join(folderPath, "*.ui"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(uiPaths) != 0 {
		for _, path := range uiPaths {
			filePaths = append(filePaths, path)
		}
	}
	ccPaths, err := filepath.Glob(filepath.Join(folderPath, "*.cc"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(ccPaths) != 0 {
		for _, path := range ccPaths {
			filePaths = append(filePaths, path)
		}
	}
	cxxPaths, err := filepath.Glob(filepath.Join(folderPath, "*.cxx"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(cxxPaths) != 0 {
		for _, path := range cxxPaths {
			filePaths = append(filePaths, path)
		}
	}

	return filePaths
}
