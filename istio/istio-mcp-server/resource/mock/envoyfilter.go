package mock

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	ptypes "github.com/gogo/protobuf/types"
	mcp "istio.io/api/mcp/v1alpha1"
	networking "istio.io/api/networking/v1alpha3"

	"istio-mcp-server/resource"
	"istio-mcp-server/types"
	"istio-mcp-server/util"
)

type mockenvoyfilter struct {
	l      sync.Mutex
	snap   *types.ResourceSnap
	source types.Source
}

func (ef *mockenvoyfilter) All() (*types.ResourceSnap, error) {
	ef.l.Lock()
	defer ef.l.Unlock()

	if ef.snap != nil {
		return ef.snap, nil
	}

	ef.createNew()

	return ef.snap, nil
}

func (ef *mockenvoyfilter) createNew() {
	ef.snap = &types.ResourceSnap{
		Version:   resource.BuildVersion(),
		Resources: []*mcp.Resource{},
	}

	patchValue := `
name: "mosn-demo:20882"
virtual_hosts:
- name: "mosn.io.dubbo.DemoService:20882"
  retry_policy:
    num_retries: 3
  routes:
  - match:
      prefix: "/"
    route:
      timeout: 10s
      cluster: "outbound|20882||mosn.io.dubbo.DemoService"
      retry_policy:
        num_retries: 5
        per_try_timeout: 3s
`

	patch, _ := util.YAML2Struct(patchValue)
	data := &networking.EnvoyFilter{
		ConfigPatches: []*networking.EnvoyFilter_EnvoyConfigObjectPatch{
			{
				ApplyTo: networking.EnvoyFilter_ROUTE_CONFIGURATION,
				Match: &networking.EnvoyFilter_EnvoyConfigObjectMatch{
					ObjectTypes: &networking.EnvoyFilter_EnvoyConfigObjectMatch_RouteConfiguration{
						RouteConfiguration: &networking.EnvoyFilter_RouteConfigurationMatch{
							PortNumber: 20882,
							Vhost: &networking.EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch{
								Name: "mosn.io.dubbo.DemoService:80",
								Route: &networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch{
									Action: networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_ANY,
								},
							},
						},
					},
				},
				Patch: &networking.EnvoyFilter_Patch{
					Operation: networking.EnvoyFilter_Patch_MERGE,
					Value:     patch,
				},
			},
		},
	}
	b, _ := ptypes.MarshalAny(data)

	ef.snap.Resources = append(ef.snap.Resources, &mcp.Resource{
		Metadata: &mcp.Metadata{
			Name:    "dubbo-ef",
			Version: resource.BuildVersion(),
		},
		Body: b,
	})
}

func (ef *mockenvoyfilter) Update(c *gin.Context) {
	ef.l.Lock()
	defer ef.l.Unlock()

	ef.snap = &types.ResourceSnap{
		Version:   resource.BuildVersion(),
		Resources: []*mcp.Resource{},
	}

	patchValue := `
name: "mosn-demo:20882"
virtual_hosts:
- name: "mosn.io.dubbo.DemoService:20882"
  retry_policy:
    num_retries: 3
  routes:
  - match:
      prefix: "/"
    route:
      timeout: 20s
      cluster: "outbound|20882||mosn.io.dubbo.DemoService"
      retry_policy:
        num_retries: 3
        per_try_timeout: 5s
`

	patch, _ := util.YAML2Struct(patchValue)
	data := &networking.EnvoyFilter{
		ConfigPatches: []*networking.EnvoyFilter_EnvoyConfigObjectPatch{
			{
				ApplyTo: networking.EnvoyFilter_ROUTE_CONFIGURATION,
				Match: &networking.EnvoyFilter_EnvoyConfigObjectMatch{
					ObjectTypes: &networking.EnvoyFilter_EnvoyConfigObjectMatch_RouteConfiguration{
						RouteConfiguration: &networking.EnvoyFilter_RouteConfigurationMatch{
							PortNumber: 20882,
							Vhost: &networking.EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch{
								Name: "mosn.io.dubbo.DemoService:80",
								Route: &networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch{
									Action: networking.EnvoyFilter_RouteConfigurationMatch_RouteMatch_ANY,
								},
							},
						},
					},
				},
				Patch: &networking.EnvoyFilter_Patch{
					Operation: networking.EnvoyFilter_Patch_MERGE,
					Value:     patch,
				},
			},
		},
	}
	b, _ := ptypes.MarshalAny(data)

	ef.snap.Resources = append(ef.snap.Resources, &mcp.Resource{
		Metadata: &mcp.Metadata{
			Name:    "dubbo-ef",
			Version: resource.BuildVersion(),
		},
		Body: b,
	})

	ef.source.Push(types.IstioCRDEnvoyFilter, ef.snap)

	c.JSON(http.StatusOK, "ok")
}
