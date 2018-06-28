package stub

import (
	"context"
	"fmt"

	"github.com/vtuson/cm-operator/pkg/apis/cm/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

var objects map[string]*v1alpha1.Chartmuseum

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	if objects == nil {
		objects = make(map[string]*v1alpha1.Chartmuseum)
	}
	switch o := event.Object.(type) {
	case *v1alpha1.Chartmuseum:
		// err := sdk.Create(newbusyBoxPod(o))
		if !event.Deleted {
			exists := objects[o.GetObjectMeta().GetName()]
			if exists == nil {
				objects[o.GetObjectMeta().GetName()] = o
				fmt.Println("creating endpoing for", o.GetObjectMeta().GetName())
				for _, dep := range o.Spec.Dependencies {
					addDependency(dep.Name, dep.Url)
				}
				addRepo(o.GetObjectMeta().GetName(), o.Spec.Git)
			} else {
				//TODO - check for changes
			}
		} else {
			fmt.Println("deleting endpoing for", o.GetObjectMeta().GetName())
			deleteRepo(o.GetObjectMeta().GetName())
			objects[o.GetObjectMeta().GetName()] = nil
		}
	}
	return nil
}

/* newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha1.Chartmuseum) *corev1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Chartmuseum",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
*/
