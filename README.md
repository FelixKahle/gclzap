# Google Cloud Logging Zap

## Overview

This Go package provides an easy way to write logs to Google Cloud Logging. It leverages the official Google Cloud client library for Go, ensuring compatibility and ease of use.

## Usage

### Installation

`go get github.com/FelixKahle/gclzap`

### Example

```go
package main

import (
    "github.com/FelixKahle/gclzap"
    "go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "YOUR_PROJECT_ID"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "my-log"

	logger := client.Logger(logName)

	// Creates a new Zap logger.
	zapLogger := gclzap.NewProduction(logger)

	// Logs a message.
	zapLogger.Info("Hello, world!")
}
```
