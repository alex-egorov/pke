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
// EksEncryptionConfig EKS encryption configuration object describing an encryption provider and the corresponding resources to encrypt. More information can be found at https://docs.aws.amazon.com/eks/latest/APIReference/API_EncryptionConfig.html.
type EksEncryptionConfig struct {
	Provider EksEncryptionConfigProvider `json:"provider,omitempty"`
	// Resource kinds to encrypt with the corresponding encryption provider.
	Resources []string `json:"resources,omitempty"`
}
