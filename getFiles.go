package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getIcons(folderPath string, searchTerm string) []string {
	var icons []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasSuffix(strings.ToLower(info.Name()), ".svg") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".png") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".jpeg") {
				if searchTerm == "" || strings.Contains(strings.ToLower(info.Name()), strings.ToLower(searchTerm)) {
					icons = append(icons, path)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return icons
}

func getThemes(folderPath string) []string {
	var themes []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasSuffix(strings.ToLower(info.Name()), ".theme") {
				themes = append(themes, path)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	return themes
}
