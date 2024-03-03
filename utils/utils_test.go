// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

package utils

import (
	"testing"
)

// For testing
// $ cd utils
// $ go test -v

func TestGenerateName(t *testing.T) {
	// Act
	var generated = GenerateName()
	t.Logf("String generated:\t%+v", generated)

	// Assert
	if generated == "" {
		t.Errorf("Expected string not empty")
	}
}
