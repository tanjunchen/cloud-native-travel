package pool

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/transport"
)

type pool struct {
	size int
	ttl  time.Duration
	tr   transport.Transport

	sync.Mutex
	conns map[string][]*poolConn
}

type poolConn struct {
	transport.Client
	id      string
	created time.Time
}

func newPool(options Options) *pool {
	return &pool{
		size:  options.Size,
		tr:    options.Transport,
		ttl:   options.TTL,
		conns: make(map[string][]*poolConn),
	}
}

// NoOp the Close since we manage it
func (p *poolConn) Close() error {
	return nil
}

func (p *poolConn) Id() string {
	return p.id
}

func (p *poolConn) Created() time.Time {
	return p.created
}

func (p *pool) Close() error {
	p.Lock()
	for k, c := range p.conns {
		for _, conn := range c {
			conn.Client.Close()
		}
		delete(p.conns, k)
	}
	p.Unlock()
	return nil
}

func (p *pool) Get(addr string, opts ...transport.DialOption) (Conn, error) {
	p.Lock()
	conns := p.conns[addr]

	for len(conns) > 0 {
		conn := conns[len(conns)-1]
		conns = conns[:len(conns)-1]
		p.conns[addr] = conns

		if d := time.Since(conn.Created()); d > p.ttl {
			conn.Client.Close()
			continue
		}
		p.Unlock()
		return conn, nil
	}
	p.Unlock()

	c, err := p.tr.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	return &poolConn{
		Client:  c,
		id:      uuid.New().String(),
		created: time.Now(),
	}, nil
}
func (p *pool) Release(conn Conn, err error) error {
	if err != nil {
		return conn.(*poolConn).Client.Close()
	}

	p.Lock()
	conns := p.conns[conn.Remote()]
	if len(conns) >= p.size {
		p.Unlock()
		return conn.(*poolConn).Client.Close()
	}
	p.conns[conn.Remote()] = append(conns, conn.(*poolConn))
	p.Unlock()

	return nil
}
