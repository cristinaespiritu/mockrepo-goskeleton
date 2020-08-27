// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// +build integration

package service

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/test/service"
	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/test/shared"
	"testing"
)

func TestSetupIntegration(t *testing.T) {
	//Once the IOTSS_TEST_SECRET is exposed as globalVariable access it using below
	fmt.Println("shared.Globalsecret---", shared.Globalsecret)
	var a string = "This is your sample integration test setup"
	var b string = "This is your sample integration test setup"

	err := errors.New("two words should be the same")

	assert.Equal(t, a, b, err)

	fmt.Println("integration test pass")
}

// To be added in all integration tests to get the IOTSS_TEST_SECRET value
func init() {

	service.Setup()
}
