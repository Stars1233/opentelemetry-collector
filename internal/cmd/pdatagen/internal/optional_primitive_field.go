// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/internal/cmd/pdatagen/internal"
import (
	"strings"
	"text/template"
)

const optionalPrimitiveAccessorsTemplate = `// {{ .fieldName }} returns the {{ .lowerFieldName }} associated with this {{ .structName }}.
func (ms {{ .structName }}) {{ .fieldName }}() {{ .returnType }} {
	return ms.orig.Get{{ .fieldName }}()
}

// Has{{ .fieldName }} returns true if the {{ .structName }} contains a
// {{ .fieldName }} value, false otherwise.
func (ms {{ .structName }}) Has{{ .fieldName }}() bool {
	return ms.orig.{{ .fieldName }}_ != nil
}

// Set{{ .fieldName }} replaces the {{ .lowerFieldName }} associated with this {{ .structName }}.
func (ms {{ .structName }}) Set{{ .fieldName }}(v {{ .returnType }}) {
	ms.state.AssertMutable()
	ms.orig.{{ .fieldName }}_ = &{{ .originStructType }}{{ "{" }}{{ .fieldName }}: v}
}

// Remove{{ .fieldName }} removes the {{ .lowerFieldName }} associated with this {{ .structName }}.
func (ms {{ .structName }}) Remove{{ .fieldName }}() {
	ms.state.AssertMutable()
	ms.orig.{{ .fieldName }}_ = nil
}`

const optionalPrimitiveAccessorsTestTemplate = `func Test{{ .structName }}_{{ .fieldName }}(t *testing.T) {
	ms := New{{ .structName }}()
	{{- if eq .returnType "float64" }}
	assert.InDelta(t, {{ .defaultVal }}, ms.{{ .fieldName }}() , 0.01)
	{{- else }}
	assert.Equal(t, {{ .defaultVal }}, ms.{{ .fieldName }}())
	{{- end }}
	ms.Set{{ .fieldName }}({{ .testValue }})
	assert.True(t, ms.Has{{ .fieldName }}())
	{{- if eq .returnType "float64" }}
	assert.InDelta(t, {{.testValue }}, ms.{{ .fieldName }}(), 0.01)
	{{- else }}
	assert.Equal(t, {{ .testValue }}, ms.{{ .fieldName }}())
	{{- end }}
	ms.Remove{{ .fieldName }}()
	assert.False(t, ms.Has{{ .fieldName }}())
	dest := New{{ .structName }}()
	dest.Set{{ .fieldName }}({{ .testValue }})
	ms.CopyTo(dest)
	assert.False(t, dest.Has{{ .fieldName }}())
}`

const optionalPrimitiveSetTestTemplate = `tv.orig.{{ .fieldName }}_ = &{{ .originStructType }}{
{{- .fieldName }}: {{ .testValue }}}`

const optionalPrimitiveCopyOrigTemplate = `if src{{ .fieldName }}, ok := src.{{ .fieldName }}_.(*{{ .originStructType }}); ok {
	dest{{ .fieldName }}, ok := dest.{{ .fieldName }}_.(*{{ .originStructType }})
	if !ok {
		dest{{ .fieldName }} = &{{ .originStructType }}{}
		dest.{{ .fieldName }}_ = dest{{ .fieldName }}
	}
	dest{{ .fieldName }}.{{ .fieldName }} = src{{ .fieldName }}.{{ .fieldName }}
} else {
	dest.{{ .fieldName }}_ = nil
}`

const optionalPrimitiveMarshalJSONTemplate = `if ms.Has{{ .fieldName }}() {
		dest.WriteObjectField("{{ lowerFirst .fieldName }}")
		dest.Write{{ upperFirst .returnType }}(ms.{{ .fieldName }}())
	}`

const optionalPrimitiveUnmarshalJSONTemplate = `case "{{ lowerFirst .fieldName }}"{{ if needSnake .fieldName -}}, "{{ toSnake .fieldName }}"{{- end }}:
		ms.orig.{{ .fieldName }}_ = &{{ .originStructType }}{{ "{" }}{{ .fieldName }}: iter.Read{{ upperFirst .returnType }}()}`

type OptionalPrimitiveField struct {
	fieldName string
	protoType ProtoType
}

func (opv *OptionalPrimitiveField) GenerateAccessors(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveAccessorsTemplate").Parse(optionalPrimitiveAccessorsTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) GenerateAccessorsTest(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveAccessorsTestTemplate").Parse(optionalPrimitiveAccessorsTestTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) GenerateSetWithTestValue(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveSetTestTemplate").Parse(optionalPrimitiveSetTestTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) GenerateCopyOrig(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveCopyOrigTemplate").Parse(optionalPrimitiveCopyOrigTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) GenerateMarshalJSON(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveMarshalJSONTemplate").Parse(optionalPrimitiveMarshalJSONTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) GenerateUnmarshalJSON(ms *messageStruct) string {
	t := template.Must(templateNew("optionalPrimitiveUnmarshalJSONTemplate").Parse(optionalPrimitiveUnmarshalJSONTemplate))
	return executeTemplate(t, opv.templateFields(ms))
}

func (opv *OptionalPrimitiveField) templateFields(ms *messageStruct) map[string]any {
	return map[string]any{
		"structName":       ms.getName(),
		"packageName":      "",
		"defaultVal":       opv.protoType.defaultValue(),
		"fieldName":        opv.fieldName,
		"lowerFieldName":   strings.ToLower(opv.fieldName),
		"testValue":        opv.protoType.testValue(opv.fieldName),
		"returnType":       opv.protoType.goType(),
		"originStructName": ms.originFullName,
		"originStructType": ms.originFullName + "_" + opv.fieldName,
	}
}

var _ Field = (*OptionalPrimitiveField)(nil)
