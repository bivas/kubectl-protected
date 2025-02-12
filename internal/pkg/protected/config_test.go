package protected

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProtectedPatterns(t *testing.T) {

	input := `
protected:
  - cluster1
  - .*pattern
`
	result, err := ReadProtectedPatterns(strings.NewReader(input))
	assert.NoError(t, err, "ReadProtectedPatterns should not return an error")
	assert.Len(t, result, 2, "ReadProtectedPatterns should return the expected result")
	assert.Equal(t, []string{"cluster1", ".*pattern"}, result, "ReadProtectedPatterns should return the expected result")
}
