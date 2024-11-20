package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/damirpavlik/proxy-pilot/internal/config"
)

func Run() error {
	appConfig, err := config.LoadConfig("settings", "config")
	if err != nil {
		return fmt.Errorf("skill issue loading configuration: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	for _, resource := range appConfig.Resources {
		targetURL, err := url.Parse(resource.DestinationURL)
		if err != nil {
			log.Printf("skipping resource %s: invalid URL %s - %v", resource.Name, resource.DestinationURL, err)
			continue
		}

		proxy := NewReverseProxy(targetURL)

		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, targetURL, resource.Endpoint))
		log.Printf("registered proxy route: %s -> %s", resource.Endpoint, resource.DestinationURL)
	}

	serverAddr := fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Listen_port)
	log.Printf("starting server on %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
