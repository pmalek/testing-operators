package fakes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes10(t *testing.T) {
	t.Run("gateways can be created, listed, modified and deleted", func(t *testing.T) {
		ctx := context.Background()

		fakeClient := fakectrlruntimeclient.
			NewClientBuilder().
			Build()

		gateway := &gatewayv1beta1.Gateway{ObjectMeta: metav1.ObjectMeta{Name: "gateway"}}

		// Error: no kind is registered for the type v1beta1.Gateway in scheme "pkg/runtime/scheme.go:100
		require.NoError(t, fakeClient.Create(ctx, gateway))

		gatewayList := &gatewayv1beta1.GatewayList{}
		require.NoError(t, fakeClient.List(ctx, gatewayList))
		require.Len(t, gatewayList.Items, 0)

		require.NoError(t, fakeClient.List(ctx, gatewayList))
		require.Len(t, gatewayList.Items, 1)

		require.NoError(t, fakeClient.Delete(ctx, gateway))
	})
}
