// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package model

//EnvVarModel structure
type EnvVarModel struct {
	SaltSecret  string `json:"saltSecret"`
	Environment string `json:"environment"`
}
