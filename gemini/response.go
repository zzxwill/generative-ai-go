package gemini

import (
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func ResponseToString(resp *genai.GenerateContentResponse) string {
	return responseString(resp)
}

// responseString is copied from https://github.com/google/generative-ai-go/blob/2813ecb3f680192c8742eccfbf44d1eebb659829/genai/client_test.go#L529
func responseString(resp *genai.GenerateContentResponse) string {
	var b strings.Builder
	for i, cand := range resp.Candidates {
		if len(resp.Candidates) > 1 {
			fmt.Fprintf(&b, "%d:", i+1)
		}
		b.WriteString(contentString(cand.Content))
	}
	return b.String()
}

func contentString(c *genai.Content) string {
	var b strings.Builder
	if c == nil || c.Parts == nil {
		return ""
	}
	for i, part := range c.Parts {
		if i > 0 {
			fmt.Fprintf(&b, ";")
		}
		fmt.Fprintf(&b, "%v", part)
	}
	return b.String()
}
