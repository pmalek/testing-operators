package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dyn_fake "k8s.io/client-go/dynamic/fake"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// NOTE: Broken:
// https://kubernetes.slack.com/archives/C0EG7JC6T/p1696267521516559

func TestFakes22(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.AddToScheme(scheme))

	dynClient := dyn_fake.NewSimpleDynamicClientWithCustomListKinds(
		scheme,
		map[schema.GroupVersionResource]string{
			{
				Group:    "gateway.networking.k8s.io",
				Version:  "v1beta1",
				Resource: "gateways",
			}: "GatewayList",
		},
	)

	r := dynClient.Resource(schema.GroupVersionResource{
		Group:    "gateway.networking.k8s.io",
		Version:  "v1beta1",
		Resource: "gateways",
	})

	ctx := context.Background()

	t.Run("dynamic client for custom GVR can list resources", func(t *testing.T) {
		unstructuredList, err := r.List(ctx, metav1.ListOptions{})
		require.NoError(t, err)
		assert.Len(t, unstructuredList.Items, 0)
	})

	t.Run("dynamic client for custom GVR can create and get resources", func(t *testing.T) {
		u := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "gateway.networking.k8s.io/v1beta1",
				"kind":       "Gateway",
				"metadata": map[string]interface{}{
					"name": "gateway",
				},
			},
		}
		u.SetAPIVersion("gateway.networking.k8s.io/v1beta1")
		u.SetKind("Gateway")
		u.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "gateway.networking.k8s.io",
			Version: "v1beta1",
			Kind:    "Gateway",
		})

		_, err := r.Create(ctx, u, metav1.CreateOptions{})
		require.NoError(t, err)

		gateway, err := r.Get(ctx, "gateway", metav1.GetOptions{})
		require.NoError(t, err)
		assert.Equal(t, "gateway", gateway.Object["metadata"].(map[string]interface{})["name"])
	})

	t.Run("dynamic client for custom GVR can list resources", func(t *testing.T) {
		unstructuredList, err := r.List(ctx, metav1.ListOptions{})
		require.NoError(t, err)
		assert.Len(t, unstructuredList.Items, 1)
	})
}
