package resource

import (
	"time"

	"istio.io/istio/pkg/config/schema/collections"
	"k8s.io/klog/v2"

	"istio-mcp-server/types"
)

var FactorySnap map[string]types.Snap

func init() {
	FactorySnap = make(map[string]types.Snap)
}

// Registry registry need care resource
func Registry(ele string, snap types.Snap) {
	if _, ok := FactorySnap[ele]; ok {
		klog.Errorf("duplicate registry resource:%s", ele)
		return
	}
	FactorySnap[ele] = snap
}

// GetAllResource get all pilot watch resource
func GetAllResource() []string {
	var cols []string
	for _, col := range collections.Pilot.All() {
		cols = append(cols, col.Name().String())
	}
	return cols
}

// BuildVersion build resource snap version
func BuildVersion() string {
	return time.Now().Format("20060102150405")
}
