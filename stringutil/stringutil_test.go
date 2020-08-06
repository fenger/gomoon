package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessString(t *testing.T) {
	var tmpl = "hello, {{.Name}}"
	var value = "github"
	vars := make(map[string]interface{})
	vars["Name"] = value
	result, _ := ParseString(tmpl, vars)
	assert.Equal(t, "hello, github", result)
}
