package transformer

import (
	"fmt"
	"strings"

	"github.com/livensmi1e/jsonflower/internal/parser"
)

func TransformJSON2YAML(value parser.Value) string {
	return strings.TrimSuffix(transformYAML(value, 0), "\n")
}

func transformYAML(value parser.Value, indent int) string {
	var sb strings.Builder
	indentStr := strings.Repeat("  ", indent)
	switch v := value.(type) {
	case *parser.Object:
		sb.WriteString(transformObject(v, indent))
	case *parser.Array:
		for _, value := range v.Elements {
			sb.WriteString(indentStr + "- ")
			if obj, ok := value.(*parser.Object); ok {
				sb.WriteString(transformObjectInArray(obj, indent+1))
			} else {
				sb.WriteString(transformYAML(value, indent+1) + "\n")
			}
		}
		if sb.String() == "" {
			sb.WriteString("[]")
		}
	case parser.String:
		sb.WriteString(toRawString(v.Literal))
	case parser.Number:
		sb.WriteString(formatNumber(v.Value))
	case parser.Boolean:
		sb.WriteString(formatBool(v.Value))
	case parser.Null:
		sb.WriteString("null")
	}
	return sb.String()
}

func formatNumber(num float64) string {
	return fmt.Sprintf("%g", num)
}

func formatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func isComplex(value parser.Value) bool {
	switch v := value.(type) {
	case *parser.Object:
		return len(v.KeyValue) != 0
	case *parser.Array:
		return len(v.Elements) != 0
	default:
		return false
	}
}

func toRawString(quoted string) string {
	return strings.TrimPrefix(strings.TrimSuffix(quoted, "\""), "\"")
}

func transformObjectInArray(obj *parser.Object, indent int) string {
	var sb strings.Builder
	first := true
	for key, val := range obj.KeyValue {
		if first {
			if isComplex(val) {
				sb.WriteString("\n" + toRawString(key) + ": " + transformYAML(val, indent+1))
			} else {
				sb.WriteString(toRawString(key) + ": " + transformYAML(val, indent+1) + "\n")
			}
			first = false
		} else {
			sb.WriteString(strings.Repeat("  ", indent) + toRawString(key) + ": ")
			if isComplex(val) {
				sb.WriteString("\n" + transformYAML(val, indent+1))
			} else {
				sb.WriteString(transformYAML(val, indent+1) + "\n")
			}
		}
	}
	str := sb.String()
	if str == "" {
		return "{}"
	} else {
		return str
	}
}

func transformObject(obj *parser.Object, indent int) string {
	var sb strings.Builder
	for key, value := range obj.KeyValue {
		sb.WriteString(strings.Repeat("  ", indent) + toRawString(key) + ": ")
		if isComplex(value) {
			sb.WriteString("\n" + transformYAML(value, indent+1))
		} else {
			sb.WriteString(transformYAML(value, indent+1) + "\n")
		}
	}
	str := sb.String()
	if str == "" {
		return "{}"
	} else {
		return str
	}
}
