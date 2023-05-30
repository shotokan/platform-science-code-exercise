package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadPath(title string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s: ", title)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}

	path := scanner.Text()
	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("paht file %s is not a valid path", path)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("file %s does not exist", path)
	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("it is not a valid file: %s", path)
	}

	return scanner.Text(), nil
}
