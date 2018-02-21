package server

import (
	"github.com/sirupsen/logrus"
	"net"
	"github.com/gorilla/mux"
)


// HTTPServer is used to wrap the services and expose them over an HTTP interface
type HTTPServer struct {
	mux        *mux.Router
	listener   net.Listener
	listenerCh chan struct{}
	logger     *logrus.Entry
	Addr       string
}

// NewHTTPServer starts new HTTP server over the services
func NewHTTPServer(config *Config) (*HTTPServer, error) {
	// Start the listener
	lnAddr, err := net.ResolveTCPAddr("tcp", config.normalizedAddrs.HTTP)
	if err != nil {
		return nil, err
	}
	ln, err := config.Listener("tcp", lnAddr.IP.String(), lnAddr.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to start HTTP listener: %v", err)
	}

	// If TLS is enabled, wrap the listener with a TLS listener
	if config.TLSConfig.EnableHTTP {
		tlsConf := &tlsutil.Config{
			VerifyIncoming:       config.TLSConfig.VerifyHTTPSClient,
			VerifyOutgoing:       true,
			VerifyServerHostname: config.TLSConfig.VerifyServerHostname,
			CAFile:               config.TLSConfig.CAFile,
			CertFile:             config.TLSConfig.CertFile,
			KeyFile:              config.TLSConfig.KeyFile,
			KeyLoader:            config.TLSConfig.GetKeyLoader(),
		}
		tlsConfig, err := tlsConf.IncomingTLSConfig()
		if err != nil {
			return nil, err
		}
		ln = tls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, tlsConfig)
	}

	// Create the mux
	mux := http.NewServeMux()

	// Create the server
	srv := &HTTPServer{
		agent:      agent,
		mux:        mux,
		listener:   ln,
		listenerCh: make(chan struct{}),
		logger:     agent.logger,
		Addr:       ln.Addr().String(),
	}
	srv.registerHandlers(config.EnableDebug)

	// Handle requests with gzip compression
	gzip, err := gziphandler.GzipHandlerWithOpts(gziphandler.MinSize(0))
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(srv.listenerCh)
		http.Serve(ln, gzip(mux))
	}()

	return srv, nil
}