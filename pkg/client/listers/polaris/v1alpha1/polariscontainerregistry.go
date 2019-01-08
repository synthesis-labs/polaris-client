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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/synthesis-labs/polaris-operator/pkg/apis/polaris/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolarisContainerRegistryLister helps list PolarisContainerRegistries.
type PolarisContainerRegistryLister interface {
	// List lists all PolarisContainerRegistries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PolarisContainerRegistry, err error)
	// PolarisContainerRegistries returns an object that can list and get PolarisContainerRegistries.
	PolarisContainerRegistries(namespace string) PolarisContainerRegistryNamespaceLister
	PolarisContainerRegistryListerExpansion
}

// polarisContainerRegistryLister implements the PolarisContainerRegistryLister interface.
type polarisContainerRegistryLister struct {
	indexer cache.Indexer
}

// NewPolarisContainerRegistryLister returns a new PolarisContainerRegistryLister.
func NewPolarisContainerRegistryLister(indexer cache.Indexer) PolarisContainerRegistryLister {
	return &polarisContainerRegistryLister{indexer: indexer}
}

// List lists all PolarisContainerRegistries in the indexer.
func (s *polarisContainerRegistryLister) List(selector labels.Selector) (ret []*v1alpha1.PolarisContainerRegistry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolarisContainerRegistry))
	})
	return ret, err
}

// PolarisContainerRegistries returns an object that can list and get PolarisContainerRegistries.
func (s *polarisContainerRegistryLister) PolarisContainerRegistries(namespace string) PolarisContainerRegistryNamespaceLister {
	return polarisContainerRegistryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolarisContainerRegistryNamespaceLister helps list and get PolarisContainerRegistries.
type PolarisContainerRegistryNamespaceLister interface {
	// List lists all PolarisContainerRegistries in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PolarisContainerRegistry, err error)
	// Get retrieves the PolarisContainerRegistry from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PolarisContainerRegistry, error)
	PolarisContainerRegistryNamespaceListerExpansion
}

// polarisContainerRegistryNamespaceLister implements the PolarisContainerRegistryNamespaceLister
// interface.
type polarisContainerRegistryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PolarisContainerRegistries in the indexer for a given namespace.
func (s polarisContainerRegistryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PolarisContainerRegistry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolarisContainerRegistry))
	})
	return ret, err
}

// Get retrieves the PolarisContainerRegistry from the indexer for a given namespace and name.
func (s polarisContainerRegistryNamespaceLister) Get(name string) (*v1alpha1.PolarisContainerRegistry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("polariscontainerregistry"), name)
	}
	return obj.(*v1alpha1.PolarisContainerRegistry), nil
}
