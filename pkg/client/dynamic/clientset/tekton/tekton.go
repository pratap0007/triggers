/*
Copyright 2021 The Tekton Authors

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

package tekton

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"github.com/tektoncd/triggers/pkg/apis/triggers"
	"github.com/tektoncd/triggers/pkg/client/dynamic/clientset"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

var (
	allowedPipelineTypes = map[string][]string{
		"v1alpha1": {"pipelineresources", "pipelineruns", "taskruns", "pipelines", "clustertasks", "tasks", "conditions", "runs"},
		"v1beta1":  {"pipelineruns", "taskruns", "pipelines", "clustertasks", "tasks"},
	}
	allowedTriggersTypes = map[string][]string{
		"v1alpha1": {"clusterinterceptors", "interceptors"},
		"v1beta1":  {"clustertriggerbindings", "eventlisteners", "triggerbindings", "triggers", "triggertemplates"},
	}
)

// WithClient adds Tekton related clients to the Dynamic client.
func WithClient(client dynamic.Interface) clientset.Option {
	return func(cs *clientset.Clientset) {
		for version, resources := range allowedPipelineTypes {
			for _, resource := range resources {
				r := schema.GroupVersionResource{
					Group:    pipeline.GroupName,
					Version:  version,
					Resource: resource,
				}
				cs.Add(r, client)
			}
		}
		for version, resources := range allowedTriggersTypes {
			for _, resource := range resources {
				r := schema.GroupVersionResource{
					Group:    triggers.GroupName,
					Version:  version,
					Resource: resource,
				}
				cs.Add(r, client)
			}
		}
	}
}
