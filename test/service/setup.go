// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"fmt"
	"os"

	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/test/shared"
)

// To check whether IOTSS_TEST_SECRET is exported as environment variable and make it as globalVariable to be available in all integration tests
func Setup() {
	secret := ""
	secret = shared.Globalsecret
	fmt.Println("IOTSS_TEST_SECRET", secret)
	if secret == "" {
		secret = os.Getenv("IOTSS_TEST_SECRET")
	}
	fmt.Println("The value of IOTSS_TEST_SECRET is", secret)
	if secret == "" {
		fmt.Println("Seems like IOTSS_TEST_SECRET value is not set as environment variable \n Please set this value using \t export IOTSS_TEST_SECRET= \n Then re-run the test")
		os.Exit(1)
	} else {
		fmt.Println("The IOTSS_TEST_SECRET is already exported as global variable")
		shared.Globalsecret = secret
	}
}
