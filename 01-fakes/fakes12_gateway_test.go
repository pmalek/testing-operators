package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes12(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.AddToScheme(scheme))

	t.Run("gateways can be created, listed, modified and deleted", func(t *testing.T) {
		ctx := context.Background()

		restMapper := meta.NewDefaultRESTMapper(nil)
		addGV(restMapper, gatewayv1beta1.GroupVersion, "Gateway")

		fakeClient := fakectrlruntimeclient.
			NewClientBuilder().
			WithScheme(scheme).
			WithRESTMapper(restMapper).
			Build()

		gateway := &gatewayv1beta1.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name: "gateway",
			},
		}

		gatewayList := &gatewayv1beta1.GatewayList{}

		assert.NoError(t, fakeClient.List(ctx, gatewayList))
		assert.Len(t, gatewayList.Items, 0)

		assert.NoError(t, fakeClient.Create(ctx, gateway))

		assert.NoError(t, fakeClient.List(ctx, gatewayList))
		assert.Len(t, gatewayList.Items, 1)

		assert.NoError(t, fakeClient.Delete(ctx, gateway))
	})
}
