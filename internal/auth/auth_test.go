package auth

import (
	"fmt"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	fmt.Println("Test Case 1 - Invalid Auth string")
	headers := map[string][]string{
		"Authorization": {"ApiKeyabc123xyz345def678hij91011"},
	}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("Expecting 'error' but found 'nil'")
	}
	fmt.Println("Test Case 2 - Valid Auth string")
	headers = map[string][]string{
		"Authorization": {"ApiKey abc123xyz345def678hij91011"},
	}
	_, err = GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expecting 'nil' but found 'error'")
	}
	fmt.Println("Test Case 3 - Invalid Authentication method")
	headers = map[string][]string{
		"Authorization": {"Basic abc123xyz345def678hij91011"},
	}
	_, err = GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expecting 'error' but found 'nil'")
	}
}
