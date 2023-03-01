package mock

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"istio-mcp-server/resource"
	"istio-mcp-server/types"
)

var (
	ef = &mockenvoyfilter{}
	vs = &mockvirtualservices{}
	se = &mockserviceentry{}
)

type mock struct{}

func New(s types.Source) types.LogicServer {
	m := &mock{}

	ef.source = s
	vs.source = s
	se.source = s

	m.Registry()
	return m
}

func (m *mock) Registry() {
	resource.Registry(types.IstioCRDEnvoyFilter, ef)
	resource.Registry(types.IstioCRDVirtualService, vs)
	resource.Registry(types.IstioCRDServiceEntry, se)
}

func (m *mock) Start(stop <-chan struct{}) {
	serv := &http.Server{
		Addr: ":8080",
	}
	engin := gin.Default()
	engin.POST("/envoyfilter", ef.Update)
	engin.POST("/virtualservice", vs.Update)
	engin.POST("/serviceentry", se.Update)

	serv.Handler = engin
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			klog.Errorf("start http server failed:%s", err)
			return
		}
	}()
	<-stop
	serv.Shutdown(context.Background())
}
