package parser

import (
	"agora-aggregator/pkg/fetcher"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name          string
		content       string
		expectedStatus string
	}{
		{
			name: "Simple Case",
			content: `
# ADR 1: Use Go
Status: PROPOSED
Some content here.
`,
			expectedStatus: "PROPOSED",
		},
		{
			name: "Lowercase Status",
			content: `
# ADR 2: Use Cobra
status: DRAFT
More content.
`,
			expectedStatus: "DRAFT",
		},
		{
			name: "Extra Whitespace",
			content: `
# ADR 3: Monorepo
Status:   ACCEPTED
Even more content.
`,
			expectedStatus: "ACCEPTED",
		},
		{
			name: "No Status",
			content: `
# ADR 4: No Status
This one is missing a status line.
`,
			expectedStatus: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			adr := &fetcher.ADR{Content: tc.content}
			Parse(adr)
			if adr.Status != tc.expectedStatus {
				t.Errorf("expected status %q, got %q", tc.expectedStatus, adr.Status)
			}
		})
	}
}
