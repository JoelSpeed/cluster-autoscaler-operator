/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 Red Hat, Inc.
 *
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/cluster-autoscaler-operator/pkg/apis/autoscaling/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterAutoscalerLister helps list ClusterAutoscalers.
// All objects returned here must be treated as read-only.
type ClusterAutoscalerLister interface {
	// List lists all ClusterAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterAutoscaler, err error)
	// Get retrieves the ClusterAutoscaler from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterAutoscaler, error)
	ClusterAutoscalerListerExpansion
}

// clusterAutoscalerLister implements the ClusterAutoscalerLister interface.
type clusterAutoscalerLister struct {
	indexer cache.Indexer
}

// NewClusterAutoscalerLister returns a new ClusterAutoscalerLister.
func NewClusterAutoscalerLister(indexer cache.Indexer) ClusterAutoscalerLister {
	return &clusterAutoscalerLister{indexer: indexer}
}

// List lists all ClusterAutoscalers in the indexer.
func (s *clusterAutoscalerLister) List(selector labels.Selector) (ret []*v1.ClusterAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterAutoscaler))
	})
	return ret, err
}

// Get retrieves the ClusterAutoscaler from the index for a given name.
func (s *clusterAutoscalerLister) Get(name string) (*v1.ClusterAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterautoscaler"), name)
	}
	return obj.(*v1.ClusterAutoscaler), nil
}
