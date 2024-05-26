/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AuthorizationModelRequestSpec defines the desired state of AuthorizationModelRequest
type AuthorizationModelRequestSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	AuthorizationModel string `json:"authorizationModel,omitempty"`
	Version            string `json:"version,omitempty"`
}

// AuthorizationModelRequestStatus defines the observed state of AuthorizationModelRequest
type AuthorizationModelRequestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AuthorizationModelRequest is the Schema for the authorizationmodelrequests API
type AuthorizationModelRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthorizationModelRequestSpec   `json:"spec,omitempty"`
	Status AuthorizationModelRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AuthorizationModelRequestList contains a list of AuthorizationModelRequest
type AuthorizationModelRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthorizationModelRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AuthorizationModelRequest{}, &AuthorizationModelRequestList{})
}
