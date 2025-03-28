package beautify

import (
	"fmt"
	"strings"

	"github.com/livensmi1e/jsonflower/internal/parser"
)

func BeautifyAST(value parser.Value) string {
	return formatAST(value, 0)
}

func formatAST(value parser.Value, indent int) string {
	var sb strings.Builder
	indentStr := strings.Repeat("  ", indent)
	switch v := value.(type) {
	case *parser.Object:
		sb.WriteString("{\n")
		for key, value := range v.KeyValue {
			sb.WriteString(indentStr + "  " + keyStyle.Render(key) + ": ")
			sb.WriteString(formatAST(value, indent+1) + ",\n")
		}
		sb.WriteString(indentStr + "}")
	case *parser.Array:
		sb.WriteString("[\n")
		for _, value := range v.Elements {
			sb.WriteString(indentStr + "  " + formatAST(value, indent+1) + ",\n")
		}
		sb.WriteString(indentStr + "]")
	case parser.String:
		sb.WriteString(strStyle.Render(v.Literal))
	case parser.Number:
		sb.WriteString(numStyle.Render(formatNumber(v.Value)))
	case parser.Boolean:
		sb.WriteString(boolStyle.Render(formatBool(v.Value)))
	case parser.Null:
		sb.WriteString(nullStyle.Render("null"))
	}
	return sb.String()
}

func formatNumber(num float64) string {
	return numStyle.Render(fmt.Sprintf("%g", num))
}

func formatBool(b bool) string {
	if b {
		return boolStyle.Render("true")
	}
	return boolStyle.Render("false")
}
