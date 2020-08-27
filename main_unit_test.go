// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// +build unit

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "two words should be the same")

	fmt.Println("unit test pass")
}
