package mcpx

import (
	"encoding/json"
	"fmt"
)

func ToMcpText(label string, data any) string {
	jsonData, _ := json.MarshalIndent(data, "", "  ")

	return fmt.Sprintf("### %s\n%s", label, string(jsonData))
}
