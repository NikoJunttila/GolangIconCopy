package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getResolution(imagePath string) string {

	if filepath.Ext(imagePath) == ".svg" {
		baseFolderName := getBaseFolderName(imagePath)
		if strings.Contains(baseFolderName, "@") {
			size := strings.Split(baseFolderName, "@")
			return size[0] + "x" + size[0]
		}
		return baseFolderName + "x" + baseFolderName
	}
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return "error"
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return "error"
	}
	if strings.Contains(imagePath, "@") {
		return strconv.Itoa(img.Width/2) + "x" + strconv.Itoa(img.Height/2)
	}
	return strconv.Itoa(img.Width) + "x" + strconv.Itoa(img.Height)
}

func neededResos(bigString string) []string {
	slice := strings.Split(bigString, ",")
	for index, reso := range slice {
		if !strings.Contains(reso, "x") {
			slice[index] = reso + "x" + reso
		}
	}
	return slice
}
func getBaseFolderName(filePath string) string {
	dir := filepath.Dir(filePath)
	base := filepath.Base(dir)
	return base
}
