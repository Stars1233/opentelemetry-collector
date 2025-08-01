// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestFunction_MoveTo(t *testing.T) {
	ms := generateTestFunction()
	dest := NewFunction()
	ms.MoveTo(dest)
	assert.Equal(t, NewFunction(), ms)
	assert.Equal(t, generateTestFunction(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestFunction(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newFunction(&otlpprofiles.Function{}, &sharedState)) })
	assert.Panics(t, func() { newFunction(&otlpprofiles.Function{}, &sharedState).MoveTo(dest) })
}

func TestFunction_CopyTo(t *testing.T) {
	ms := NewFunction()
	orig := NewFunction()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestFunction()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newFunction(&otlpprofiles.Function{}, &sharedState)) })
}

func TestFunction_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestFunction()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	// Append an unknown field at the start to ensure unknown fields are skipped
	// and the unmarshal logic continues.
	buf := stream.Buffer()
	assert.EqualValues(t, '{', buf[0])
	iter := json.BorrowIterator(append([]byte(`{"unknown": "string",`), buf[1:]...))
	defer json.ReturnIterator(iter)
	dest := NewFunction()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestFunction_NameStrindex(t *testing.T) {
	ms := NewFunction()
	assert.Equal(t, int32(0), ms.NameStrindex())
	ms.SetNameStrindex(int32(13))
	assert.Equal(t, int32(13), ms.NameStrindex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newFunction(&otlpprofiles.Function{}, &sharedState).SetNameStrindex(int32(13)) })
}

func TestFunction_SystemNameStrindex(t *testing.T) {
	ms := NewFunction()
	assert.Equal(t, int32(0), ms.SystemNameStrindex())
	ms.SetSystemNameStrindex(int32(13))
	assert.Equal(t, int32(13), ms.SystemNameStrindex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newFunction(&otlpprofiles.Function{}, &sharedState).SetSystemNameStrindex(int32(13)) })
}

func TestFunction_FilenameStrindex(t *testing.T) {
	ms := NewFunction()
	assert.Equal(t, int32(0), ms.FilenameStrindex())
	ms.SetFilenameStrindex(int32(13))
	assert.Equal(t, int32(13), ms.FilenameStrindex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newFunction(&otlpprofiles.Function{}, &sharedState).SetFilenameStrindex(int32(13)) })
}

func TestFunction_StartLine(t *testing.T) {
	ms := NewFunction()
	assert.Equal(t, int64(0), ms.StartLine())
	ms.SetStartLine(int64(13))
	assert.Equal(t, int64(13), ms.StartLine())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newFunction(&otlpprofiles.Function{}, &sharedState).SetStartLine(int64(13)) })
}

func generateTestFunction() Function {
	tv := NewFunction()
	fillTestFunction(tv)
	return tv
}

func fillTestFunction(tv Function) {
	tv.orig.NameStrindex = int32(13)
	tv.orig.SystemNameStrindex = int32(13)
	tv.orig.FilenameStrindex = int32(13)
	tv.orig.StartLine = int64(13)
}
