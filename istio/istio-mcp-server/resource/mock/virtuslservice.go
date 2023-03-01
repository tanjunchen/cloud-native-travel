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
)

type mockvirtualservices struct {
	l      sync.Mutex
	snap   *types.ResourceSnap
	source types.Source
}

func (vs *mockvirtualservices) All() (*types.ResourceSnap, error) {
	vs.l.Lock()
	defer vs.l.Unlock()

	if vs.snap != nil {
		return vs.snap, nil
	}

	vs.createNew()

	return vs.snap, nil
}

func (vs *mockvirtualservices) createNew() {
	vs.snap = &types.ResourceSnap{
		Version:   resource.BuildVersion(),
		Resources: []*mcp.Resource{},
	}

	data := &networking.VirtualService{
		Hosts: []string{"istiosvc"},
		Http: []*networking.HTTPRoute{
			{
				Route: []*networking.HTTPRouteDestination{
					{
						Destination: &networking.Destination{
							Host:   "istiosvc",
							Subset: "v1",
						},
						Weight: 50,
					},
					{
						Destination: &networking.Destination{
							Host:   "istiosvc",
							Subset: "v2",
						},
						Weight: 50,
					},
				},
			},
		},
	}
	b, _ := ptypes.MarshalAny(data)

	vs.snap.Resources = append(vs.snap.Resources, &mcp.Resource{
		Metadata: &mcp.Metadata{
			Name:    "aabb-server",
			Version: resource.BuildVersion(),
		},
		Body: b,
	})
}

func (vs *mockvirtualservices) Update(c *gin.Context) {
	// TODO update VirtualServices resource
	c.JSON(http.StatusOK, "ok")
}
