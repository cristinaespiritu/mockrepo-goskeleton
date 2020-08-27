// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// +build integration

package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMainIntegration(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	err := errors.New("two words should be the same")

	assert.Equal(t, a, b, err)

	fmt.Println("integration test pass")
}
