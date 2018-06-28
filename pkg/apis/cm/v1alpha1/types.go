package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartmuseumList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Chartmuseum `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Chartmuseum struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ChartmuseumSpec   `json:"spec"`
	Status            ChartmuseumStatus `json:"status,omitempty"`
}

type ChartmuseumDependency struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type ChartmuseumSpec struct {
	Git          string                  `json:"git"`
	Freq         int                     `json:"updateEveryMinutes,omitempty"`
	Dependencies []ChartmuseumDependency `json:"dependencies,omitempty"`
}
type ChartmuseumStatus struct {
	Sync     string `json:"sync"`
	LastSync string `json:"lastSync"`
}
