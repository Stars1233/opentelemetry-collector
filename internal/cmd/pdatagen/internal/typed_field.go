// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/internal/cmd/pdatagen/internal"

import (
	"strings"
	"text/template"
)

const typedAccessorsTemplate = `// {{ .fieldName }} returns the {{ .lowerFieldName }} associated with this {{ .structName }}.
func (ms {{ .structName }}) {{ .fieldName }}() {{ .packageName }}{{ .returnType }} {
	return {{ .packageName }}{{ .returnType }}(ms.orig.{{ .originFieldName }})
}

// Set{{ .fieldName }} replaces the {{ .lowerFieldName }} associated with this {{ .structName }}.
func (ms {{ .structName }}) Set{{ .fieldName }}(v {{ .packageName }}{{ .returnType }}) {
	ms.state.AssertMutable()
	ms.orig.{{ .originFieldName }} = {{ .rawType }}(v)
}`

const typedAccessorsTestTemplate = `func Test{{ .structName }}_{{ .fieldName }}(t *testing.T) {
	ms := New{{ .structName }}()
	assert.Equal(t, {{ .packageName }}{{ .returnType }}({{ .defaultVal }}), ms.{{ .fieldName }}())
	testVal{{ .fieldName }} := {{ .packageName }}{{ .returnType }}({{ .testValue }})
	ms.Set{{ .fieldName }}(testVal{{ .fieldName }})
	assert.Equal(t, testVal{{ .fieldName }}, ms.{{ .fieldName }}())
}`

const typedSetTestTemplate = `tv.orig.{{ .originFieldName }} = {{ .testValue }}`

const typedCopyOrigTemplate = `dest.{{ .originFieldName }} = src.{{ .originFieldName }}`

const typedMarshalJSONTemplate = `if ms.orig.{{ .originFieldName }} != {{ .defaultVal }} {
		dest.WriteObjectField("{{ lowerFirst .originFieldName }}")
		{{- if .isType }}
		ms.orig.{{ .originFieldName }}.MarshalJSONStream(dest)
		{{- else if .isEnum }}
		dest.WriteInt32(int32(ms.orig.{{ .originFieldName }}))
		{{- else }}	
		dest.Write{{ upperFirst .rawType }}(ms.orig.{{ .originFieldName }})
		{{- end }}
	}`

const typedUnmarshalJSONTemplate = `case "{{ lowerFirst .originFieldName }}"{{ if needSnake .originFieldName -}}, "{{ toSnake .originFieldName }}"{{- end }}:
		{{- if .isType }}
		ms.orig.{{ .originFieldName }}.UnmarshalJSONIter(iter)
		{{- else if .isEnum }}
		ms.orig.{{ .originFieldName }} = {{ .rawType }}(iter.ReadEnumValue({{ .rawType }}_value))
		{{- else }}	
		ms.orig.{{ .originFieldName }} = iter.Read{{ upperFirst .rawType }}()
		{{- end }}`

// TypedField is a field that has defined a custom type (e.g. "type Timestamp uint64")
type TypedField struct {
	fieldName       string
	originFieldName string
	returnType      *TypedType
}

type TypedType struct {
	structName  string
	packageName string
	rawType     string
	isType      bool
	isEnum      bool
	defaultVal  string
	testVal     string
}

func (ptf *TypedField) GenerateAccessors(ms *messageStruct) string {
	t := template.Must(templateNew("typedAccessorsTemplate").Parse(typedAccessorsTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) GenerateAccessorsTest(ms *messageStruct) string {
	t := template.Must(templateNew("typedAccessorsTestTemplate").Parse(typedAccessorsTestTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) GenerateSetWithTestValue(ms *messageStruct) string {
	t := template.Must(templateNew("typedSetTestTemplate").Parse(typedSetTestTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) GenerateCopyOrig(ms *messageStruct) string {
	t := template.Must(templateNew("typedCopyOrigTemplate").Parse(typedCopyOrigTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) GenerateMarshalJSON(ms *messageStruct) string {
	t := template.Must(templateNew("typedMarshalJSONTemplate").Parse(typedMarshalJSONTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) GenerateUnmarshalJSON(ms *messageStruct) string {
	t := template.Must(templateNew("typedUnmarshalJSONTemplate").Parse(typedUnmarshalJSONTemplate))
	return executeTemplate(t, ptf.templateFields(ms))
}

func (ptf *TypedField) templateFields(ms *messageStruct) map[string]any {
	return map[string]any{
		"structName": ms.getName(),
		"defaultVal": ptf.returnType.defaultVal,
		"packageName": func() string {
			if ptf.returnType.packageName != ms.packageName {
				return ptf.returnType.packageName + "."
			}
			return ""
		}(),
		"isCommon":       usedByOtherDataTypes(ptf.returnType.packageName),
		"returnType":     ptf.returnType.structName,
		"fieldName":      ptf.fieldName,
		"lowerFieldName": strings.ToLower(ptf.fieldName),
		"testValue":      ptf.returnType.testVal,
		"rawType":        ptf.returnType.rawType,
		"isType":         ptf.returnType.isType,
		"isEnum":         ptf.returnType.isEnum,
		"originFieldName": func() string {
			if ptf.originFieldName == "" {
				return ptf.fieldName
			}
			return ptf.originFieldName
		}(),
	}
}

var _ Field = (*TypedField)(nil)
