// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//nolint:golint,stylecheck
package model

//AppConfig structure
type AppConfig struct {
	ServiceName          string               `json:"serviceName"`
	ServiceVersion       string               `json:"serviceVersion"`
	ServiceMaintenance   bool                 `json:"serviceMaintenance"`
	ServiceConfiguration ServiceConfiguration `json:"serviceConfiguration"`
}

//ServiceConfiguration structure
type ServiceConfiguration struct {
	ServiceCertificates   ServiceCertificates   `json:"serviceCertificates"`
	ServiceLDAPConnection ServiceLDAPConnection `json:"serviceLDAPConnection"`
	ServiceLog            ServiceLog            `json:"serviceLog"`
	ServiceAPISecurity    ServiceAPISecurity    `json:"serviceAPISecurity"`
	ServiceVCenterDetails ServiceVCenterDetails `json:"serviceVCenterDetails"`
}

//ServiceVCenterDetails structure
type ServiceVCenterDetails struct {
	HostName       string `json:"hostName"`
	HTTPScheme     string `json:"httpScheme"`
	Path           string `json:"path"`
	ServiceAccount string `json:"serviceAccount"`
	Password       string `json:"password"`
}

//ServiceCertificates structure
type ServiceCertificates struct {
	CertLocation string `json:"certLocation"`
	KeyLocation  string `json:"keyLocation"`
}

//ServiceLDAPConnection structure
type ServiceLDAPConnection struct {
	DialURL            string `json:"serviceConfiguration.serviceLDAPConnection.dialUrl"`
	BindUsername       string `json:"serviceConfiguration.serviceLDAPConnection.bindUsername"`
	BindPassword       string `json:"serviceConfiguration.serviceLDAPConnection.bindPassword"`
	LdapBindDN         string `json:"serviceConfiguration.serviceLDAPConnection.ldapBindDN"`
	LdapSearchFilterDN string `json:"serviceConfiguration.serviceLDAPConnection.ldapSearchFilterDN"`
}

//ServiceLog structure
type ServiceLog struct {
	LogPath      string `json:"logPath"`
	LogFormatter string `json:"logFormatter"`
}

//ServiceAPISecurity structure
type ServiceAPISecurity struct {
	TokenExpiryDuration int    `json:"tokenExpiryDuration"`
	AllowOrigin         string `json:"allowOrigin"`
	AllowedMethods      string `json:"allowedMethods"`
}

//ServiceDbConnections structure
type ServiceDbConnections struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}
