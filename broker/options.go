package broker

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"os"
	"time"
)

type brokerOptions struct {
	// basic options
	ID               int64  `flag:"worker-id" cfg:"id"`
	Verbose          bool   `flag:"verbose"`
	TCPAddress       string `flag:"tcp-address"`
	HTTPAddress      string `flag:"http-address"`
	BroadcastAddress string `flag:"broadcast-address"`
	// NSQLookupdTCPAddresses []string `flag:"lookupd-tcp-address" cfg:"nsqlookupd_tcp_addresses"`

	// diskqueue options
	// MemQueueSize    int64         `flag:"mem-queue-size"`
	// MaxBytesPerFile int64         `flag:"max-bytes-per-file"`
	// SyncEvery       int64         `flag:"sync-every"`
	// SyncTimeout     time.Duration `flag:"sync-timeout"`

	// msg and command options
	MsgTimeout    time.Duration `flag:"msg-timeout" arg:"1ms"`
	MaxMsgTimeout time.Duration `flag:"max-msg-timeout"`
	MaxMsgSize    int64         `flag:"max-msg-size" deprecated:"max-message-size" cfg:"max_msg_size"`
	MaxBodySize   int64         `flag:"max-body-size"`
	ClientTimeout time.Duration

	// client overridable configuration options
	MaxHeartbeatInterval time.Duration `flag:"max-heartbeat-interval"`
	// MaxRdyCount            int64         `flag:"max-rdy-count"`
	MaxOutputBufferSize    int64         `flag:"max-output-buffer-size"`
	MaxOutputBufferTimeout time.Duration `flag:"max-output-buffer-timeout"`

	// statsd integration
	StatsdAddress  string        `flag:"statsd-address"`
	StatsdPrefix   string        `flag:"statsd-prefix"`
	StatsdInterval time.Duration `flag:"statsd-interval" arg:"1s"`
	StatsdMemStats bool          `flag:"statsd-mem-stats"`

	// e2e message latency
	// E2EProcessingLatencyWindowTime  time.Duration `flag:"e2e-processing-latency-window-time"`
	// E2EProcessingLatencyPercentiles []float64     `flag:"e2e-processing-latency-percentile" cfg:"e2e_processing_latency_percentiles"`

	// TLS config
	TLSCert string `flag:"tls-cert"`
	TLSKey  string `flag:"tls-key"`

	// compression
	DeflateEnabled  bool `flag:"deflate"`
	MaxDeflateLevel int  `flag:"max-deflate-level"`
	SnappyEnabled   bool `flag:"snappy"`
}

func NewBrokerOptions() *brokerOptions {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	o := &brokerOptions{
		TCPAddress:  "0.0.0.0:4150",
		HTTPAddress: "0.0.0.0:4151",

		MsgTimeout:    60 * time.Second,
		MaxMsgTimeout: 15 * time.Minute,
		MaxMsgSize:    1024768,
		MaxBodySize:   5 * 1024768,
		ClientTimeout: 60 * time.Second,

		MaxHeartbeatInterval:   60 * time.Second,
		MaxOutputBufferSize:    64 * 1024,
		MaxOutputBufferTimeout: 1 * time.Second,

		StatsdPrefix:   "broker.%s",
		StatsdInterval: 60 * time.Second,
		StatsdMemStats: true,

		//E2EProcessingLatencyWindowTime: time.Duration(10 * time.Minute),

		DeflateEnabled:  true,
		MaxDeflateLevel: 6,
		SnappyEnabled:   true,
	}

	h := md5.New()
	io.WriteString(h, hostname)
	o.ID = int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)

	return o
}
