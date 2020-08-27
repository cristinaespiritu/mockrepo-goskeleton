// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// +build unit

package util

import (
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plainText := "MyPassword123"
	enc, _ := Encrypt(plainText, "TEST_SECRET")
	sEnc := EncodeBase64String(enc)
	fmt.Println(sEnc)
	dec, err := DecodeAndDecrypt(sEnc, "TEST_SECRET")
	if err != nil {
		t.Errorf("Error in decryption")
	}
	sDec := string(dec)
	if plainText != sDec {
		t.Errorf("Expected string doesn't match")
	}
}
