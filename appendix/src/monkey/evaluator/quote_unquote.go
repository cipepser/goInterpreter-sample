package evaluator

import (
	"github.com/cipepser/goInterpreter-sample/appendix/src/monkey/ast"
	"github.com/cipepser/goInterpreter-sample/appendix/src/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
