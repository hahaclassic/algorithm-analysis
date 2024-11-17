package scrapper

import (
	"fmt"
	"os"
)

func inputParams(baseURL, dirPath string, maxWorkers, maxPages int) {
	fmt.Println("[PARAMETERS]",
		"\n- baseURL:", baseURL,
		"\n- dirPath:", dirPath,
		"\n- maxWorkers:", maxWorkers,
		"\n- maxPages:", maxPages,
	)
}

func results(visited int) {
	fmt.Println("\n[RESULT]",
		"\n- visited:", visited,
		//"\n- saved:", saved,
	)
}

func createDir(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("%w: %w", ErrCreateDir, err)
	}

	return nil
}
