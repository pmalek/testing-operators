package fakes

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	gatewayv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFakes99(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, gatewayv1beta1.AddToScheme(scheme))

	restMapper := meta.NewDefaultRESTMapper(nil)
	as := func(gv metav1.GroupVersion, kind, resourcePlural string) {
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
	as(gatewayv1beta1.GroupVersion, "Gateway", "gateways")

	fakeClient := fakectrlruntimeclient.
		NewClientBuilder().
		WithRESTMapper(restMapper).
		WithScheme(scheme).
		Build()

	ctx := context.Background()
	gateway := &gatewayv1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name: "gateway",
		},
	}

	gatewayList := &gatewayv1beta1.GatewayList{}

	assert.NoError(t, fakeClient.Create(ctx, gateway))
	assert.NoError(t, fakeClient.List(ctx, gatewayList))
	assert.NoError(t, fakeClient.Delete(ctx, gateway))
}
