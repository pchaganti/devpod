package v1

import (
	clusterv1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type ProjectMembers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Teams holds all the teams that have access to the cluster
	Teams []ProjectMember `json:"teams,omitempty"`

	// Users holds all the users that have access to the cluster
	Users []ProjectMember `json:"users,omitempty"`
}

type ProjectMember struct {
	// Info about the user or team
	// +optional
	Info clusterv1.EntityInfo `json:"info,omitempty"`
}
