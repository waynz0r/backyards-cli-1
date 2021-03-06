// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphql

import (
	"context"
	"errors"

	"github.com/banzaicloud/istio-client-go/pkg/networking/v1alpha3"

	"github.com/MakeNowJust/heredoc"
)

// the only reason for not using the types from the Istio client go package here is that Istio uses snake case in YAML...
type Sidecar struct {
	Spec      SidecarSpec `json:"spec"`
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
}

type SidecarSpec struct {
	WorkloadSelector      *WorkloadSelector               `json:"workloadSelector,omitempty"`
	Ingress               []*IstioIngressListener         `json:"ingress,omitempty"`
	Egress                []*IstioEgressListener          `json:"egress"`
	OutboundTrafficPolicy *v1alpha3.OutboundTrafficPolicy `json:"outboundTrafficPolicy,omitempty"`
}

type WorkloadSelector struct {
	Labels map[string]string `json:"labels,omitempty"`
}

type IstioIngressListener struct {
	Port            *v1alpha3.Port       `json:"port"`
	Bind            string               `json:"bind,omitempty"`
	CaptureMode     v1alpha3.CaptureMode `json:"captureMode,omitempty"`
	DefaultEndpoint string               `json:"defaultEndpoint"`
}

type IstioEgressListener struct {
	Port        *v1alpha3.Port       `json:"port,omitempty"`
	Bind        string               `json:"bind,omitempty"`
	CaptureMode v1alpha3.CaptureMode `json:"capture_mode,omitempty"`
	Hosts       []string             `json:"hosts"`
}

type MeshWorkload struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

type MeshWorkloadSidecar struct {
	MeshWorkload `json:",inline"`

	Sidecars            []Sidecar `json:"sidecars,omitempty"`
	RecommendedSidecars []Sidecar `json:"recommendedSidecars,omitempty"`
}

func (c *client) GetWorkload(namespace, name string) (*MeshWorkload, error) {
	request := heredoc.Doc(`
	query($namespace: String!, $name: String!) {
      workload(namespace: $namespace, name: $name) {
        id
        name
        namespace
        labels
	  }
    }
`)

	type Response struct {
		Workload MeshWorkload `json:"workload"`
	}

	r := c.NewRequest(request)
	r.Var("name", name)
	r.Var("namespace", namespace)

	// run it and capture the response
	var respData Response
	if err := c.client.Run(context.Background(), r, &respData); err != nil {
		return nil, err
	}

	if respData.Workload.ID == "" {
		return nil, errors.New("not found")
	}

	return &respData.Workload, nil
}

func (c *client) GetWorkloadWithSidecar(namespace, name string) (*MeshWorkloadSidecar, error) {
	request := heredoc.Doc(`
	query($namespace: String!, $name: String!) {
      workload(namespace: $namespace, name: $name) {
        id
        name
        namespace
        labels
        sidecars {
          name
          namespace
          spec {
            workloadSelector {
              labels
            }
            egress {
              port {
                number
                protocol
                name
              }
              bind
              captureMode
              hosts
            }
            ingress {
              port {
                number
                protocol
                name
              }
              bind
              captureMode
              defaultEndpoint
            }
            outboundTrafficPolicy
          }
        }
      }
    }
`)

	type Response struct {
		Workload MeshWorkloadSidecar `json:"workload"`
	}

	r := c.NewRequest(request)
	r.Var("name", name)
	r.Var("namespace", namespace)

	// run it and capture the response
	var respData Response
	if err := c.client.Run(context.Background(), r, &respData); err != nil {
		return nil, err
	}

	if respData.Workload.ID == "" {
		return nil, errors.New("not found")
	}

	return &respData.Workload, nil
}

func (c *client) GetWorkloadWithSidecarRecommendation(namespace, name string, isolationLevel string, labelWhitelist []string) (*MeshWorkloadSidecar, error) {
	request := heredoc.Doc(`
	query($namespace: String!, $name: String!, $isolationLevel: IsolationLevel, $labelWhitelist: [String!]) {
      workload(namespace: $namespace, name: $name) {
        id
        name
        namespace
        labels
        recommendedSidecars(isolationLevel: $isolationLevel, labelWhitelist: $labelWhitelist) {
          name
          namespace
          spec {
            workloadSelector {
              labels
            }
            egress {
              port {
                number
                protocol
                name
              }
              bind
              captureMode
              hosts
            }
          }
        }
      }
    }
`)

	type Response struct {
		Workload MeshWorkloadSidecar `json:"workload"`
	}

	r := c.NewRequest(request)
	r.Var("name", name)
	r.Var("namespace", namespace)
	if isolationLevel != "" {
		r.Var("isolationLevel", isolationLevel)
	}
	r.Var("labelWhitelist", labelWhitelist)

	// run it and capture the response
	var respData Response
	if err := c.client.Run(context.Background(), r, &respData); err != nil {
		return nil, err
	}

	if respData.Workload.ID == "" {
		return nil, errors.New("not found")
	}

	return &respData.Workload, nil
}
