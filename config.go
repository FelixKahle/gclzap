// Copyright (c) 2024 Felix Kahle.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package gclzap

import (
	"cloud.google.com/go/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config is a configuration struct for the zap.Logger that writes logs to Google Cloud Logging.
type Config struct {
	EncoderConfig EncoderConfig
	Level         zapcore.Level
}

// Build creates a new zap.Logger that writes logs to Google Cloud Logging.
//
// Parameters:
// - logger: The logger to use for logging.
//
// Returns:
// - A new zap.Logger that writes logs to Google Cloud Logging.
func (c Config) Build(logger *logging.Logger) *zap.Logger {
	return New(logger, c)
}

// NewProductionConfig returns a new configuration for the zap.Logger that writes logs to Google Cloud Logging.
//
// Returns:
// - A new configuration for the zap.Logger that writes logs to Google Cloud Logging.
func NewProductionConfig() Config {
	return Config{
		EncoderConfig: DefaultEncoderConfig(),
		Level:         zapcore.InfoLevel,
	}
}

// NewDevelopmentConfig returns a new configuration for the zap.Logger that writes logs to Google Cloud Logging.
//
// Returns:
// - A new configuration for the zap.Logger that writes logs to Google Cloud Logging.
func NewDevelopmentConfig() Config {
	return Config{
		EncoderConfig: DefaultEncoderConfig(),
		Level:         zapcore.DebugLevel,
	}
}
