package transformer

import (
	"fmt"
	"strings"

	"github.com/livensmi1e/jsonflower/internal/parser"
)

func TransformJSON2YAML(value parser.Value) string {
	return transformYAML(value, 0)
}

func transformYAML(value parser.Value, indent int) string {
	var sb strings.Builder
	indentStr := strings.Repeat("  ", indent)
	switch v := value.(type) {
	case *parser.Object:
		for key, value := range v.KeyValue {
			sb.WriteString(indentStr + toRawString(key) + ": ")
			if isComplex(value) {
				sb.WriteString("\n" + transformYAML(value, indent+1))
			} else {
				sb.WriteString(transformYAML(value, indent+1) + "\n")
			}
		}
	case *parser.Array:
		for _, value := range v.Elements {

			sb.WriteString(indentStr + "- " + transformYAML(value, indent+1) + "\n")
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
	switch value.(type) {
	case *parser.Object, *parser.Array:
		return true
	default:
		return false
	}
}

func toRawString(quoted string) string {
	return strings.TrimPrefix(strings.TrimSuffix(quoted, "\""), "\"")
}
