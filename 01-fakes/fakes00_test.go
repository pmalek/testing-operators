package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestFakes00(t *testing.T) {
	t.Run("pods can be created, listed, modified and deleted", func(t *testing.T) {
		ctx := context.Background()
		fakeClient := fakectrlruntimeclient.NewClientBuilder().Build()

		podList := &corev1.PodList{}
		assert.NoError(t, fakeClient.List(ctx, podList))
		assert.Len(t, podList.Items, 0)

		pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "pod"}}
		assert.NoError(t, fakeClient.Create(ctx, pod))

		assert.NoError(t, fakeClient.List(ctx, podList))
		assert.Len(t, podList.Items, 1)

		assert.NoError(t, fakeClient.Delete(ctx, pod))
	})

	t.Run("pods prepopulate fake clients object store", func(t *testing.T) {
		ctx := context.Background()
		pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "pod"}}
		fakeClient := fakectrlruntimeclient.NewClientBuilder().WithObjects(pod).Build()

		assert.NoError(t, fakeClient.Get(ctx, types.NamespacedName{Name: "pod"}, pod))

		podList := corev1.PodList{}
		assert.NoError(t, fakeClient.List(ctx, &podList))
		assert.Len(t, podList.Items, 1)
	})
}
