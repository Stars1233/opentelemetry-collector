// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptrace // import "go.opentelemetry.io/collector/pdata/ptrace"

import (
	"slices"

	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/internal/otlp"
)

// JSONMarshaler marshals Traces to JSON bytes using the OTLP/JSON format.
type JSONMarshaler struct{}

// MarshalTraces to the OTLP/JSON format.
func (*JSONMarshaler) MarshalTraces(td Traces) ([]byte, error) {
	dest := json.BorrowStream(nil)
	defer json.ReturnStream(dest)
	td.marshalJSONStream(dest)
	return slices.Clone(dest.Buffer()), dest.Error()
}

// JSONUnmarshaler unmarshals OTLP/JSON formatted-bytes to Traces.
type JSONUnmarshaler struct{}

// UnmarshalTraces from OTLP/JSON format into Traces.
func (*JSONUnmarshaler) UnmarshalTraces(buf []byte) (Traces, error) {
	iter := json.BorrowIterator(buf)
	defer json.ReturnIterator(iter)
	td := NewTraces()
	td.unmarshalJSONIter(iter)
	if iter.Error() != nil {
		return Traces{}, iter.Error()
	}
	otlp.MigrateTraces(td.getOrig().ResourceSpans)
	return td, nil
}
