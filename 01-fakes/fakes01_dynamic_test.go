package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dyn_fake "k8s.io/client-go/dynamic/fake"
	"k8s.io/kubectl/pkg/scheme"
)

func TestFakes01(t *testing.T) {
	ctx := context.Background()

	dynClient := dyn_fake.NewSimpleDynamicClient(scheme.Scheme)
	r := dynClient.Resource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	})

	t.Run("dynamic client for custom GVR can list resources", func(t *testing.T) {
		unstructuredList, err := r.List(ctx, metav1.ListOptions{})
		require.NoError(t, err)
		require.Len(t, unstructuredList.Items, 0)
	})

	t.Run("dynamic client for custom GVR can create and then list resources", func(t *testing.T) {
		_, err := r.Create(ctx, &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Pod",
				"metadata": map[string]interface{}{
					"name": "pod",
				},
			},
		}, metav1.CreateOptions{})
		require.NoError(t, err)
		pod, err := r.Get(ctx, "pod", metav1.GetOptions{})
		require.NoError(t, err)
		require.Equal(t, "pod", pod.Object["metadata"].(map[string]interface{})["name"])

		unstructuredList, err := r.List(ctx, metav1.ListOptions{})
		require.NoError(t, err)
		require.Len(t, unstructuredList.Items, 1)
	})
}
