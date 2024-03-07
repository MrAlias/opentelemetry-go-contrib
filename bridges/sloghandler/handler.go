// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package sloghandler provides a bridge between the [log/slog] and
// OpenTelemetry logging.
package sloghandler // import "go.opentelemetry.io/contrib/bridges/sloghandler"

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"sync"

	"go.opentelemetry.io/otel/log"
)

const (
	bridgeName = "go.opentelemetry.io/contrib/bridge/sloghandler"
	// TODO: hook this into the release pipeline.
	bridgeVersion = "0.0.1-alpha"
)

type config struct{}

// Option configures a [Handler].
type Option interface {
	apply(config) config
}

// Handler is a [slog.Handler] that sends all logging records it receives to
// OpenTelemetry.
type Handler struct {
	// Ensure forward compatibility by explicitly making this not comparable.
	noCmp [0]func() //nolint: unused  // This is indeed used.

	attrs  []log.KeyValue
	groups groups
	logger log.Logger
}

// Compile-time check *Handler implements slog.Handler.
var _ slog.Handler = (*Handler)(nil)

// New returns a new [Handler] to be used as an [slog.Handler].
func New(lp log.LoggerProvider, opts ...Option) *Handler {
	return &Handler{
		logger: lp.Logger(
			bridgeName,
			log.WithInstrumentationVersion(bridgeVersion),
		),
	}
}

// Handle handles the Record.
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	var record log.Record
	record.SetTimestamp(r.Time)
	record.SetBody(log.StringValue(r.Message))

	const sevOffset = slog.Level(log.SeverityDebug) - slog.LevelDebug
	record.SetSeverity(log.Severity(r.Level + sevOffset))

	record.AddAttributes(h.attrs...)
	if h.groups.valid {
		h.groups.Record(record.AddAttributes, r.NumAttrs(), r.Attrs)
	} else {
		r.Attrs(func(a slog.Attr) bool {
			record.AddAttributes(convertAttr(a)...)
			return true
		})
	}

	h.logger.Emit(ctx, record)
	return nil
}

// Enable returns true if the Handler is enabled to log for the provided
// context and Level. Otherwise, false is returned if it is not enabled.
func (h Handler) Enabled(context.Context, slog.Level) bool {
	// TODO (MrAlias): The Logs Bridge API does not provide a way to retrieve
	// the current minimum logging level yet.
	// https://github.com/open-telemetry/opentelemetry-go/issues/4995
	return true
}

// WithAttrs returns a new [slog.Handler] based on h that will log using the
// passed attrs.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}

	h2 := *h
	if h2.groups.valid {
		h2.groups = h2.groups.AppendAttrs(attrs)
	} else {
		h2.attrs = slices.Grow(h2.attrs, len(attrs))
		for _, a := range attrs {
			h2.attrs = append(h2.attrs, convertAttr(a)...)
		}
	}
	return &h2
}

// WithGroup returns a new [slog.Handler] based on h that will log all messages
// and attributes within a group using name.
func (h *Handler) WithGroup(name string) slog.Handler {
	// Handlers should inline the Attrs of a group with an empty key.
	if name == "" {
		return h
	}

	h2 := *h
	h2.groups = h2.groups.Subgroup(name)
	return &h2
}

func convertAttr(attr slog.Attr) []log.KeyValue {
	if attr.Key == "" {
		if attr.Value.Kind() == slog.KindGroup {
			// A Handler should inline the Attrs of a group with an empty key.
			g := attr.Value.Group()
			out := make([]log.KeyValue, 0, len(g))
			for _, a := range g {
				out = append(out, convertAttr(a)...)
			}
			return out
		}

		if attr.Value.Any() == nil {
			// A Handler should ignore an empty Attr.
			return nil
		}
	}
	val := convertValue(attr.Value)
	return []log.KeyValue{{Key: attr.Key, Value: val}}
}

func convertValue(v slog.Value) log.Value {
	switch v.Kind() {
	case slog.KindAny:
		return log.StringValue(fmt.Sprintf("%+v", v.Any()))
	case slog.KindBool:
		return log.BoolValue(v.Bool())
	case slog.KindDuration:
		return log.Int64Value(v.Duration().Nanoseconds())
	case slog.KindFloat64:
		return log.Float64Value(v.Float64())
	case slog.KindInt64:
		return log.Int64Value(v.Int64())
	case slog.KindString:
		return log.StringValue(v.String())
	case slog.KindTime:
		return log.Int64Value(v.Time().UnixNano())
	case slog.KindUint64:
		return log.Int64Value(int64(v.Uint64()))
	case slog.KindGroup:
		g := v.Group()
		kvs := make([]log.KeyValue, 0, len(g))
		for _, a := range g {
			kvs = append(kvs, convertAttr(a)...)
		}
		return log.MapValue(kvs...)
	case slog.KindLogValuer:
		return convertValue(v.Resolve())
	default:
		panic(fmt.Sprintf("unhandled attribute kind: %s", v.Kind()))
	}
}

var groupPool = sync.Pool{
	New: func() any { return new(group) },
}

type group struct {
	name  string
	attrs []log.KeyValue
	prev  *group
}

func newGroup(name string, attrs []log.KeyValue, prev *group) *group {
	grp := groupPool.Get().(*group)
	grp.name = name
	if len(attrs) > 0 {
		grp.attrs = attrs
	} else {
		grp.attrs = grp.attrs[:0]
	}
	grp.prev = prev
	return grp
}

func (g *group) KeyValue() log.KeyValue {
	/*
		var out log.KeyValue
		// A Handler should not output groups if there are no attributes.
		if len(g.attrs) > 0 {
			out = log.Map(g.name, g.attrs...)
		}
		return out
	*/
	return log.Map(g.name, g.attrs...)
}

func (g *group) AddAttr(n int, f func(func(slog.Attr) bool)) {
	if n == 0 {
		return
	}

	g.attrs = slices.Grow(g.attrs, n)
	f(func(a slog.Attr) bool {
		g.attrs = append(g.attrs, convertAttr(a)...)
		return true
	})
}

type groups struct {
	subGroups []string
	kv        log.KeyValue
	valid     bool
}

func newGroups(grp *group) groups {
	var n int
	curr := grp
	for curr != nil {
		n++
		curr = curr.prev
	}

	switch n {
	case 0:
		return groups{valid: true}
	case 1:
		return groups{kv: grp.KeyValue(), valid: true}
	}

	n-- // Don't include top kv in subGroups
	g := groups{
		subGroups: make([]string, n),
		kv:        grp.KeyValue(),
		valid:     true,
	}

	n-- // Use as last index of g.subGroups.
	curr = grp.prev
	for curr != nil {
		g.subGroups[n] = g.kv.Key
		n--

		curr.attrs = append(curr.attrs, g.kv)
		g.kv = curr.KeyValue()

		old := curr
		curr = curr.prev

		groupPool.Put(old)
	}

	return g
}

func (g groups) repack(*group) {
	// TODO
}

func (g groups) unpack() *group {
	if !g.valid {
		return nil
	}

	grp := newGroup(g.kv.Key, g.kv.Value.AsMap(), nil)
	for _, name := range g.subGroups {
		idx := slices.IndexFunc(grp.attrs, func(a log.KeyValue) bool {
			return a.Key == name
		})
		if idx < 0 {
			// This should never happen. Drop the group if it does.
			continue
		}

		next := grp.attrs[idx]
		grp.attrs = append(grp.attrs[:idx], grp.attrs[idx+1:]...)
		grp = newGroup(next.Key, next.Value.AsMap(), grp)
	}

	return grp
}

func (g groups) Subgroup(name string) groups {
	return newGroups(newGroup(name, nil, g.unpack()))
}

func (g groups) AppendAttrs(attrs []slog.Attr) groups {
	grp := g.unpack()
	grp.AddAttr(len(attrs), func(f func(slog.Attr) bool) {
		for _, a := range attrs {
			if !f(a) {
				return
			}
		}
	})
	return newGroups(grp)
}

func (g groups) Record(sync func(...log.KeyValue), n int, f func(func(slog.Attr) bool)) {
	if n > 0 {
		sync(g.WithAttrs(n, f).kv)
		return
	}
	grp := g.unpack()
	grp = trimEmpty(grp)
	if grp == nil {
		// No attributes to sync.
		return
	}
	sync(newGroups(grp).kv)
}

func trimEmpty(grp *group) *group {
	if grp == nil || len(grp.attrs) > 0 {
		return grp
	}
	return trimEmpty(grp.prev)
}

func (g groups) WithAttrs(n int, f func(func(slog.Attr) bool)) groups {
	grp := g.unpack()
	grp.AddAttr(n, f)
	return newGroups(grp)
}
