/*
Copyright AppsCode Inc. and Contributors

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

package lib

import (
	api "go.bytebuilders.dev/audit/api/v1"

	"kmodules.xyz/client-go/discovery"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AuditEventCreator struct {
	Mapper discovery.ResourceMapper
}

func (p *AuditEventCreator) CreateEvent(obj client.Object) (*api.Event, error) {
	rid, err := p.Mapper.ResourceIDForGVK(obj.GetObjectKind().GroupVersionKind())
	if err != nil {
		return nil, err
	}

	return &api.Event{
		Resource:   obj,
		ResourceID: *rid,
	}, nil
}
