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

import "go.uber.org/zap/zapcore"

// EncoderConfig is a configuration struct for the Encoder
// used by the custom Core implementation.
type EncoderConfig struct {
	LineEnding     string
	EncodeTime     zapcore.TimeEncoder
	EncodeDuration zapcore.DurationEncoder
	EncodeCaller   zapcore.CallerEncoder
}

// DefaultEncoderConfig returns the default configuration for the Encoder.
//
// Returns:
// - The default configuration for the Encoder.
func DefaultEncoderConfig() EncoderConfig {
	return EncoderConfig{
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewEncoder creates a new Encoder based on the given configuration.
// The Encoder is used by the custom Core implementation,
// to log messages in the Google Cloud Logging structured logging format.
//
// Parameters:
// - config: The configuration for the Encoder.
//
// Returns:
// - A new Encoder based on the given configuration.
func newEncoder(config EncoderConfig) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "severity",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     config.LineEnding,
		EncodeLevel:    encodeLevel(),
		EncodeTime:     config.EncodeTime,
		EncodeDuration: config.EncodeDuration,
		EncodeCaller:   config.EncodeCaller,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// encodeLevel returns a function that encodes the given zapcore level to a string,
// based on the Google Cloud Logging structured logging format.
//
// Returns:
// - A function that encodes the given zapcore level to a string.
func encodeLevel() zapcore.LevelEncoder {
	// https://cloud.google.com/logging/docs/structured-logging
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		case zapcore.InvalidLevel:
			panic("encountered invalid log level")
		default:
			enc.AppendString("UNKNOWN")
		}
	}
}
