package server

import "time"

type option struct {
	Address   string
	GRPCPort  int
	BurstSize int
	Freq      time.Duration
}

func defaultOption() *option {
	return &option{
		Address:   "0.0.0.0",
		GRPCPort:  15015,
		Freq:      time.Microsecond * 10,
		BurstSize: 1000,
	}
}

type Option func(*option)

func WithAddress(addr string) Option {
	return func(op *option) {
		op.Address = addr
	}
}

func WithGRPCPort(port int) Option {
	return func(op *option) {
		op.GRPCPort = port
	}
}

func WithFreq(freq time.Duration) Option {
	return func(op *option) {
		op.Freq = freq
	}
}

func WithBurstSize(size int) Option {
	return func(op *option) {
		op.BurstSize = size
	}
}
