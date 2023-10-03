package fakes

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dyn_fake "k8s.io/client-go/dynamic/fake"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes22_HTTPRoute(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.Install(scheme))

	gwGVR := schema.GroupVersionResource{
		Group:    gatewayv1beta1.GroupVersion.Group,
		Version:  gatewayv1beta1.GroupVersion.Version,
		Resource: "httproutes",
	}

	client := dyn_fake.NewSimpleDynamicClient(scheme,
		newUnstructured(metav1.GroupVersion{Group: "unknown", Version: "v1"}, "HTTPRoute", "ns", "name2-baz"),
		newUnstructured(gatewayv1beta1.GroupVersion, "HTTPRoute", "ns", "name-bar"),
		newUnstructured(gatewayv1beta1.GroupVersion, "HTTPRoute", "ns", "name-baz"),
	)
	list, err := client.Resource(gwGVR).List(ctx, metav1.ListOptions{})
	require.NoError(t, err)

	expected := []unstructured.Unstructured{
		*newUnstructured(gatewayv1beta1.GroupVersion, "HTTPRoute", "ns", "name-bar"),
		*newUnstructured(gatewayv1beta1.GroupVersion, "HTTPRoute", "ns", "name-baz"),
	}
	if !equality.Semantic.DeepEqual(list.Items, expected) {
		t.Fatal(cmp.Diff(expected, list.Items))
	}
}

func TestFakes22_Gateway(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.Install(scheme))

	gwGVR := schema.GroupVersionResource{
		Group:    gatewayv1beta1.GroupVersion.Group,
		Version:  gatewayv1beta1.GroupVersion.Version,
		Resource: "gateways",
		// NOTE: https://github.com/kubernetes/client-go/issues/1082
		// Resource: "gatewaies",
	}

	client := dyn_fake.NewSimpleDynamicClient(scheme,
		newUnstructured(metav1.GroupVersion{Group: "unknown", Version: "v1"}, "Gateway", "ns", "name2-baz"),
		newUnstructured(gatewayv1beta1.GroupVersion, "Gateway", "ns", "name-bar"),
		newUnstructured(gatewayv1beta1.GroupVersion, "Gateway", "ns", "name-baz"),
	)
	list, err := client.Resource(gwGVR).List(ctx, metav1.ListOptions{})
	require.NoError(t, err)

	expected := []unstructured.Unstructured{
		*newUnstructured(gatewayv1beta1.GroupVersion, "Gateway", "ns", "name-bar"),
		*newUnstructured(gatewayv1beta1.GroupVersion, "Gateway", "ns", "name-baz"),
	}
	if !equality.Semantic.DeepEqual(list.Items, expected) {
		t.Fatal(cmp.Diff(expected, list.Items))
	}
}
