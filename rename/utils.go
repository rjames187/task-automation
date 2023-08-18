package rename

import (
	"fmt"
	"regexp"
	"strings"
)

func segregateFilePath(fullPath string) (string, string, string, error) {
	runes := []rune(fullPath)

	extSlice := []rune{}
	nameSlice := []rune{}
	pathSlice := []rune{'/'}

	mode := "ext"

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '.' {
			mode = "name"
			continue
		}
		if runes[i] == '/' {
			mode = "path"
			continue
		}
		if mode == "ext" {
			extSlice = append([]rune{runes[i]}, extSlice...)
			continue
		}
		if mode == "name" {
			nameSlice = append([]rune{runes[i]}, nameSlice...)
			continue
		}
		if mode == "path" {
			pathSlice = append([]rune{runes[i]}, pathSlice...)
		}
	}

	ext := string(extSlice)
	name := string(nameSlice)
	path := string(pathSlice)

	fmt.Println(path)

	return ext, name, path, nil
}

func flipSlashes(path string) string {
	runes := []rune(path)
	for i, char := range runes {
		if char == '\\' {
			runes[i] = '/'
		}
	}
	return string(runes)
}

func matches(name string, rgx string) (bool, error) {
	rgx = strings.ReplaceAll(rgx, `\`, `\\`)
	regex, err := regexp.Compile(rgx)
	if err != nil {
		return false, err
	}
	return regex.MatchString(name), nil
}