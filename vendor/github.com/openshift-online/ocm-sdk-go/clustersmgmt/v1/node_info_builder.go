/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// NodeInfoBuilder contains the data and logic needed to build 'node_info' objects.
//
// Provides information about a node from specific type in the cluster.
type NodeInfoBuilder struct {
	amount *int
	type_  *NodeType
}

// NewNodeInfo creates a new builder of 'node_info' objects.
func NewNodeInfo() *NodeInfoBuilder {
	return new(NodeInfoBuilder)
}

// Amount sets the value of the 'amount' attribute to the given value.
//
//
func (b *NodeInfoBuilder) Amount(value int) *NodeInfoBuilder {
	b.amount = &value
	return b
}

// Type sets the value of the 'type' attribute to the given value.
//
// Type of node received via telemetry.
func (b *NodeInfoBuilder) Type(value NodeType) *NodeInfoBuilder {
	b.type_ = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *NodeInfoBuilder) Copy(object *NodeInfo) *NodeInfoBuilder {
	if object == nil {
		return b
	}
	b.amount = object.amount
	b.type_ = object.type_
	return b
}

// Build creates a 'node_info' object using the configuration stored in the builder.
func (b *NodeInfoBuilder) Build() (object *NodeInfo, err error) {
	object = new(NodeInfo)
	object.amount = b.amount
	object.type_ = b.type_
	return
}
