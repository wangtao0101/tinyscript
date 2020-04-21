package translator

import (
	"reflect"
	"strings"
)

func IsInstanceOfExpr(instance interface{}) bool {
	ival := reflect.ValueOf(instance)
	return strings.LastIndex(ival.Type().String(), "ast.Expr") != -1
}
