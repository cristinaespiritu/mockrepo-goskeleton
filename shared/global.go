// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package shared

import (
	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/model"
	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/util"
)

// Global environment variables, configuration and logging model ...
var (
	GlobalEnv    model.EnvVarModel
	GlobalConfig model.AppConfig
	GlobalLogger = util.NewLogger()
)
