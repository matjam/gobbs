package template_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/matjam/gobbs/internal/template"
)

func TestParser(t *testing.T) {
	testState := map[string]interface{}{
		"sys_name":   "test",
		"sysop_name": "test_2",
	}

	buf := new(bytes.Buffer)
	parser := template.NewParser(testState, buf)
	err := parser.Parse("mecca_test_template.mec")
	if err != nil {
		t.Fatalf("failed: %v", err.Error())
	}

	fmt.Println(buf.String())
}
