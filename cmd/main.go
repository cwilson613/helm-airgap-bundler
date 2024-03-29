package main

import (
	"helm-airgap-bundler/pkg/bundler"
	"log"
)

func main() {
	chartURL := "https://example.com/path/to/chart"
	auth := bundler.AuthConfig{
		Username: "user", // Your registry username
		Password: "pass", // Your registry password
	}

	if err := bundler.BundleChartDependencies(chartURL, auth); err != nil {
		log.Fatalf("Failed to bundle chart dependencies: %v", err)
	}
}
