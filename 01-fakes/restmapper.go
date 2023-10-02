package fakes

import (
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func addGV(restMapper *meta.DefaultRESTMapper, gv metav1.GroupVersion, kind string) {
	restMapper.Add(
		schema.GroupVersionKind{
			Group:   gv.Group,
			Version: gv.Version,
			Kind:    kind,
		},
		meta.RESTScopeRoot,
	)
}

func addSpecificGV(restMapper *meta.DefaultRESTMapper, gv metav1.GroupVersion, kind, resourcePlural string) {
	restMapper.AddSpecific(
		schema.GroupVersionKind{
			Group:   gv.Group,
			Version: gv.Version,
			Kind:    kind,
		},
		schema.GroupVersionResource{
			Group:    gv.Group,
			Version:  gv.Version,
			Resource: resourcePlural,
		},
		schema.GroupVersionResource{
			Group:    gv.Group,
			Version:  gv.Version,
			Resource: strings.ToLower(kind),
		},
		meta.RESTScopeRoot,
	)
}
