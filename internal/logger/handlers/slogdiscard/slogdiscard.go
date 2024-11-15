package slogdiscard

import (
	"context"
	"log/slog"
)

// DiscardHandler imlements slog.Handler interface with discard methods
type DiscardHandler struct {
}

// Enabled reports whether the handler hadles records at the given level
func (d *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

// Handle handles the Record
func (d *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
func (d *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return d
}

// WithGroup returns a new Handler with given group appended to
// the receiver's existing groups.
func (d *DiscardHandler) WithGroup(_ string) slog.Handler {
	return d
}
