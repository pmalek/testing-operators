package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dyn_fake "k8s.io/client-go/dynamic/fake"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes21(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.AddToScheme(scheme))

	gwGVR := schema.GroupVersionResource{
		Group:    gatewayv1beta1.GroupVersion.Group,
		Version:  gatewayv1beta1.GroupVersion.Version,
		Resource: "gateways",
	}

	dynClient := dyn_fake.NewSimpleDynamicClientWithCustomListKinds(
		scheme,
		map[schema.GroupVersionResource]string{
			gwGVR: "GatewayList",
		},
	)

	r := dynClient.Resource(gwGVR)

	t.Run("gateways can be created, listed and deleted", func(t *testing.T) {
		ctx := context.Background()

		t.Run("dynamic client for custom GVR can list resources", func(t *testing.T) {
			unstructuredList, err := r.List(ctx, metav1.ListOptions{})
			require.NoError(t, err)
			assert.Len(t, unstructuredList.Items, 0)
		})

		t.Run("dynamic client for custom GVR can create resources", func(t *testing.T) {
			u := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": gatewayv1beta1.GroupVersion.Group + "/" + gatewayv1beta1.GroupVersion.Version,
					"kind":       "Gateway",
					"metadata": map[string]interface{}{
						"name":      "gateway",
						"namespace": "ns",
					},
				},
			}

			_, err := r.Namespace("ns").Create(ctx, u, metav1.CreateOptions{})
			require.NoError(t, err)
		})

		t.Run("dynamic client for custom GVR can delete resources", func(t *testing.T) {
			err := r.Namespace("ns").Delete(ctx, "gateway", metav1.DeleteOptions{})
			require.NoError(t, err)
		})
	})
}
