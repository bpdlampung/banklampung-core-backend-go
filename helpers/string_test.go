package helpers

import (
	"fmt"
	"testing"
)

func TestRandomAlphaNumeric(t *testing.T) {
	fmt.Println(RandomAlphaNumeric(21))
}

func TestRandomNumericString(t *testing.T) {
	fmt.Println(RandomNumericString(6))
}
