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
)

// New creates a new zap.Logger that writes logs to the given Google Cloud Logging logger.
//
// Parameters:
// - out: The Google Cloud Logging logger to write logs to.
// - config: The configuration for the zap.Logger.
// - options: Additional options for the zap.Logger.
//
// Returns:
// - A new zap.Logger that writes logs to the given Google Cloud Logging logger.
func New(out *logging.Logger, config Config, options ...zap.Option) *zap.Logger {
	core := newCore(out, config.EncoderConfig, config.Level)

	return zap.New(core, options...)
}

// NewProduction creates a new zap.Logger that writes logs to the given Google Cloud Logging logger.
// It uses the default configuration for the Core.
//
// Parameters:
// - logger: The Google Cloud Logging logger to write logs to.
//
// Returns:
// - A new zap.Logger that writes logs to the given Google Cloud Logging logger.
func NewProduction(logger *logging.Logger) *zap.Logger {
	return NewProductionConfig().Build(logger)
}

// NewDevelopment creates a new zap.Logger that writes logs to the given Google Cloud Logging logger.
// It uses the default configuration for the Core.
//
// Parameters:
// - logger: The Google Cloud Logging logger to write logs to.
//
// Returns:
// - A new zap.Logger that writes logs to the given Google Cloud Logging logger.
func NewDevelopment(logger *logging.Logger) *zap.Logger {
	return NewProductionConfig().Build(logger)
}
