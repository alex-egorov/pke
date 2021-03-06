/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// HelmRepoListItem struct for HelmRepoListItem
type HelmRepoListItem struct {
	Name string `json:"name,omitempty"`
	Cache string `json:"cache,omitempty"`
	Url string `json:"url,omitempty"`
	CertFile string `json:"certFile,omitempty"`
	KeyFile string `json:"keyFile,omitempty"`
	CaFile string `json:"caFile,omitempty"`
	PasswordSecretRef string `json:"passwordSecretRef,omitempty"`
	TlsSecretRef string `json:"tlsSecretRef,omitempty"`
}
