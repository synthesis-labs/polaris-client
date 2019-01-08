/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/synthesis-labs/polaris-client/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// PolarisBuildPipelines returns a PolarisBuildPipelineInformer.
	PolarisBuildPipelines() PolarisBuildPipelineInformer
	// PolarisBuildSteps returns a PolarisBuildStepInformer.
	PolarisBuildSteps() PolarisBuildStepInformer
	// PolarisContainerRegistries returns a PolarisContainerRegistryInformer.
	PolarisContainerRegistries() PolarisContainerRegistryInformer
	// PolarisSourceRepositories returns a PolarisSourceRepositoryInformer.
	PolarisSourceRepositories() PolarisSourceRepositoryInformer
	// PolarisStacks returns a PolarisStackInformer.
	PolarisStacks() PolarisStackInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// PolarisBuildPipelines returns a PolarisBuildPipelineInformer.
func (v *version) PolarisBuildPipelines() PolarisBuildPipelineInformer {
	return &polarisBuildPipelineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolarisBuildSteps returns a PolarisBuildStepInformer.
func (v *version) PolarisBuildSteps() PolarisBuildStepInformer {
	return &polarisBuildStepInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolarisContainerRegistries returns a PolarisContainerRegistryInformer.
func (v *version) PolarisContainerRegistries() PolarisContainerRegistryInformer {
	return &polarisContainerRegistryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolarisSourceRepositories returns a PolarisSourceRepositoryInformer.
func (v *version) PolarisSourceRepositories() PolarisSourceRepositoryInformer {
	return &polarisSourceRepositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolarisStacks returns a PolarisStackInformer.
func (v *version) PolarisStacks() PolarisStackInformer {
	return &polarisStackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}