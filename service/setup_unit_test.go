// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// +build unit

package service

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSetup(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	//err := errors.New("")

	assert.Equal(t, a, b, "two words should be the same")

	fmt.Println("unit test pass")
}
