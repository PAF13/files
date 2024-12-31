package files

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type RipgrepMatch struct {
	Type string `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
		Lines struct {
			Text string `json:"text"`
		} `json:"lines"`
		LineNumber     int `json:"line_number"`
		AbsoluteOffset int `json:"absolute_offset"`
		Submatches     []struct {
			Match struct {
				Text string `json:"text"`
			} `json:"match"`
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"submatches"`
	} `json:"data"`
}

func rg(rootDir string, text string) []RipgrepMatch {
	results := []RipgrepMatch{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 10-second timeout
	defer cancel()

	cmd := exec.CommandContext(ctx, "rg", "--json", text, rootDir)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		return nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return nil
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadBytes('\n') // Read output line by line
		if err != nil {
			if err.Error() == "EOF" { // Break on end of file
				break
			}
			fmt.Println("Error reading output:", err)
			return nil
		}

		var match RipgrepMatch
		if err := json.Unmarshal(line, &match); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		if match.Type == "match" {
			fmt.Println("Found Match in:" + match.Data.Path.Text)
			results = append(results, match)
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
	}

	return results
}
