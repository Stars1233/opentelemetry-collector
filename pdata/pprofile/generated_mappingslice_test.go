// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestMappingSlice(t *testing.T) {
	es := NewMappingSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newMappingSlice(&[]*otlpprofiles.Mapping{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewMapping()
	testVal := generateTestMapping()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestMapping(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestMappingSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newMappingSlice(&[]*otlpprofiles.Mapping{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewMappingSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestMappingSlice_CopyTo(t *testing.T) {
	dest := NewMappingSlice()
	// Test CopyTo empty
	NewMappingSlice().CopyTo(dest)
	assert.Equal(t, NewMappingSlice(), dest)

	// Test CopyTo larger slice
	src := generateTestMappingSlice()
	src.CopyTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)

	// Test CopyTo same size slice
	src.CopyTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)

	// Test CopyTo smaller size slice
	NewMappingSlice().CopyTo(dest)
	assert.Equal(t, 0, dest.Len())

	// Test CopyTo larger slice with enough capacity
	src.CopyTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
}

func TestMappingSlice_EnsureCapacity(t *testing.T) {
	es := generateTestMappingSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestMappingSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestMappingSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestMappingSlice(), es)
}

func TestMappingSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestMappingSlice()
	dest := NewMappingSlice()
	src := generateTestMappingSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestMappingSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestMappingSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}

	dest.MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestMappingSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewMappingSlice()
	emptySlice.RemoveIf(func(el Mapping) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestMappingSlice()
	pos := 0
	filtered.RemoveIf(func(el Mapping) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestMappingSlice_RemoveIfAll(t *testing.T) {
	got := generateTestMappingSlice()
	got.RemoveIf(func(el Mapping) bool {
		return true
	})
	assert.Equal(t, 0, got.Len())
}

func TestMappingSliceAll(t *testing.T) {
	ms := generateTestMappingSlice()
	assert.NotEmpty(t, ms.Len())

	var c int
	for i, v := range ms.All() {
		assert.Equal(t, ms.At(i), v, "element should match")
		c++
	}
	assert.Equal(t, ms.Len(), c, "All elements should have been visited")
}

func TestMappingSlice_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestMappingSlice()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	iter := json.BorrowIterator(stream.Buffer())
	defer json.ReturnIterator(iter)
	dest := NewMappingSlice()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestMappingSlice_Sort(t *testing.T) {
	es := generateTestMappingSlice()
	es.Sort(func(a, b Mapping) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Less(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b Mapping) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Greater(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestMappingSlice() MappingSlice {
	es := NewMappingSlice()
	fillTestMappingSlice(es)
	return es
}

func fillTestMappingSlice(es MappingSlice) {
	*es.orig = make([]*otlpprofiles.Mapping, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpprofiles.Mapping{}
		fillTestMapping(newMapping((*es.orig)[i], es.state))
	}
}
