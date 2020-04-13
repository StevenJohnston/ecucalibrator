package printer_test

import (
	"testing"

	"github.com/StevenJohnston/ecucalibrator/src/bazel-go-test/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestPrinter(t *testing.T) {
	str := "test string"
	printed := printer.Println(str)
	assert.Equal(t, "well this is wrong", printed)
}
