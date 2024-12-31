package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(anyStruct any, dir string, fileName string) {
	var path string
	// Check if the last character matches
	if len(dir) > 0 && rune(dir[len(dir)-1]) == '\\' {
		path = dir + fileName
	} else {
		path = dir + "\\" + fileName
	}

	// Serialize the Path struct to JSON
	fileJSON, err := json.MarshalIndent(anyStruct, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON to a file
	err = os.WriteFile(path, fileJSON, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
