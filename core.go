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
	"go.uber.org/zap/zapcore"
)

// Core is a custom zapcore.Core implementation that writes logs to Google Cloud Logging.
type Core struct {
	out          *logging.Logger
	enc          zapcore.Encoder
	LevelEnabler zapcore.LevelEnabler
}

// NewCore creates a new Core based on the given configuration.
//
// Parameters:
// - out: The Google Cloud Logging logger to write logs to.
// - config: The configuration for the Encoder.
// - level: The logging level.
//
// Returns:
// - A new Core.
func newCore(out *logging.Logger, config EncoderConfig, level zapcore.LevelEnabler) *Core {
	return &Core{
		out:          out,
		enc:          newEncoder(config),
		LevelEnabler: level,
	}
}

// Level returns the current logging level.
//
// Returns:
// - The current logging level.
func (c *Core) Level() zapcore.Level {
	return zapcore.LevelOf(c.LevelEnabler)
}

// Enabled returns whether the given logging level is enabled.
//
// Parameters:
// - lvl: The logging level to check.
//
// Returns:
// - Whether the given logging level is enabled.
func (c *Core) Enabled(lvl zapcore.Level) bool {
	return c.LevelEnabler.Enabled(lvl)
}

// With returns a new Core with the given fields added.
//
// Parameters:
// - fields: The fields to add.
//
// Returns:
// - A new Core with the given fields added.
func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

// Check checks whether the given entry should be logged.
//
// Parameters:
// - ent: The entry to check.
// - ce: The checked entry.
//
// Returns:
// - The checked entry.
func (c *Core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

// Write writes the given entry and fields to the log buffer.
// If the log level is higher than ErrorLevel, the log buffer is flushed.
//
// Parameters:
// - ent: The entry to write.
// - fields: The fields to write.
//
// Returns:
// - An error if the entry could not be written, nil otherwise.
func (c *Core) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	defer buf.Free()
	if err != nil {
		return err
	}

	entry := logging.Entry{
		Timestamp: ent.Time,
		Severity:  toSeverity(ent.Level),
		Payload:   buf.String(),
	}

	// Write the log entry.
	c.out.Log(entry)

	// Since we may be crashing the program, sync the output.
	if ent.Level > zapcore.ErrorLevel {
		err := c.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

// Sync flushes the log buffer.
//
// Returns:
// - An error if the log buffer could not be flushed, nil otherwise.
func (c *Core) Sync() error {
	return c.out.Flush()
}

// clone returns a copy of the Core.
//
// Returns:
// - A copy of the Core.
func (c *Core) clone() *Core {
	return &Core{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
	}
}

// addFields adds the given fields to the encoder.
//
// Parameters:
// - enc: The encoder to add the fields to.
// - fields: The fields to add.
//
// Returns:
// - The encoder with the added fields.
func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

// toSeverity converts the given zapcore level to a Google Cloud Logging severity.
//
// Parameters:
// - l: The zapcore level to convert.
//
// Returns:
// - The converted logging severity.
func toSeverity(l zapcore.Level) logging.Severity {
	switch l {
	case zapcore.DebugLevel:
		return logging.Debug
	case zapcore.InfoLevel:
		return logging.Info
	case zapcore.WarnLevel:
		return logging.Warning
	case zapcore.ErrorLevel:
		return logging.Error
	case zapcore.DPanicLevel:
		return logging.Critical
	case zapcore.PanicLevel:
		return logging.Critical
	case zapcore.FatalLevel:
		return logging.Critical
	default:
		return logging.Default
	}
}
