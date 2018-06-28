package stub

import (
	"context"
	"fmt"

	"github.com/vtuson/cm-operator/pkg/apis/cm/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

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
				if err := sdk.Create(newCronJob(o)); err != nil {
					fmt.Println(err)
				}
			} else {
				//TODO - check for changes
			}
		} else {
			fmt.Println("deleting endpoing for", o.GetObjectMeta().GetName())
			deleteRepo(o.GetObjectMeta().GetName())
			objects[o.GetObjectMeta().GetName()] = nil
			if err := sdk.Delete(newCronJob(o)); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func newCronJob(cr *v1alpha1.Chartmuseum) *v1beta1.CronJob {

	min := cr.Spec.Freq
	if min < 10 {
		min = 10
	}

	schedule := fmt.Sprintf("*/%d * * * *", min)

	labels := map[string]string{
		"app": cr.GetObjectMeta().GetName(),
	}
	name := cr.GetObjectMeta().GetName()
	return &v1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
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
		Spec: v1beta1.CronJobSpec{
			Schedule: schedule,
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							RestartPolicy: "Never",
							Containers: []v1.Container{
								v1.Container{
									Name:    "refresh",
									Image:   "vtuson/busybox",
									Command: []string{"curl"},
									Args:    []string{getServiceURL() + "/" + name + "/update"},
								},
							},
						},
					},
				},
			},
		},
	}
}
