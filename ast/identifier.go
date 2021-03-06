package ast

import (
	"bytes"
)

type Identifier struct {
	TokenAble
	Callee         *Identifier
	Value          string
	OriginalCallee *Identifier // So robot.Avatar.Name the OriginalCallee will be robot
}

func (il *Identifier) validIfCondition() bool { return true }

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	out := &bytes.Buffer{}
	if i.Callee != nil {
		out.WriteString(i.Callee.String())
		out.WriteString(".")
	}
	out.WriteString(i.Value)
	return out.String()
}
