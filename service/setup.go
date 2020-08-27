// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gitlab.eng.vmware.com/dell-iot/iotss-utils/util"
	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/model"
	"gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/shared"
	fileutil "gitlab.eng.vmware.com/dell-iot/iotss/go-skeleton-project/util"
)

var appConfig model.AppConfig
var standardLogger = shared.GlobalLogger

const (
	configPath    = "asset/config/"
	configFileExt = "_application.json"
)

// AddArgs adds command line args to shared global struct
func AddArgs() {
	secret, stage := initFlag()

	if *secret == "" || *stage == "" {
		standardLogger.InvalidArg("startup")
		os.Exit(-1)
	}

	shared.GlobalEnv.SaltSecret = *secret
	shared.GlobalEnv.Environment = *stage

	standardLogger.Info("args added")
}

// AddConfig adds admin password to shared global config
func AddConfig() {
	byteValue, err := fileutil.ReadFile(getConfigFile())
	if err != nil {
		standardLogger.Error(fmt.Sprintf("unable to read config file, %v", err))
	}

	err = json.Unmarshal(byteValue, &appConfig)
	if err != nil {
		standardLogger.Error(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	shared.GlobalConfig = appConfig
	adminPass, err := util.DecodeAndDecrypt(shared.GlobalConfig.ServiceConfiguration.ServiceVCenterDetails.Password, shared.GlobalEnv.SaltSecret)
	if err != nil {
		standardLogger.Error(err.Error())
	}
	shared.GlobalConfig.ServiceConfiguration.ServiceVCenterDetails.Password = string(adminPass)
	standardLogger.Info("config added")
}

func getConfigFile() (configFile string) {
	return configPath + shared.GlobalEnv.Environment + configFileExt
}

func initFlag() (*string, *string) {
	secret := flag.String("secret", "", "secret string")
	stage := flag.String("stage", "", "stage string")
	flag.Parse()

	return secret, stage
}
