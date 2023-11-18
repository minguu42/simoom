package cmdutil

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	fmt.Println(string(data))
	return nil
}
