package parser

import (
	"adr-aggregator/pkg/fetcher"
	"bufio"
	"strings"
)

// Parse extracts the status from the ADR content.
func Parse(adr *fetcher.ADR) {
	scanner := bufio.NewScanner(strings.NewReader(adr.Content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.ToLower(line), "status:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				status := strings.TrimSpace(parts[1])
				adr.Status = status
				return
			}
		}
	}
}
