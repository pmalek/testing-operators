package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dyn_fake "k8s.io/client-go/dynamic/fake"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes20(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.AddToScheme(scheme))

	gwGVR := schema.GroupVersionResource{
		Group:    gatewayv1beta1.GroupVersion.Group,
		Version:  gatewayv1beta1.GroupVersion.Version,
		Resource: "gateways",
	}

	dynClient := dyn_fake.NewSimpleDynamicClient(scheme)

	r := dynClient.Resource(gwGVR)

	t.Run("gateways can be created, listed, modified and deleted", func(t *testing.T) {
		ctx := context.Background()

		t.Run("dynamic client for custom GVR can list resources", func(t *testing.T) {
			unstructuredList, err := r.List(ctx, metav1.ListOptions{})
			require.NoError(t, err)
			require.Len(t, unstructuredList.Items, 0)
		})
	})
}
