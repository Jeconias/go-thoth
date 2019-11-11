package rules

import myasthurts "github.com/lab259/go-my-ast-hurts"

// RequiredWithInput TODO
type RequiredWithInput struct {
	CurrentField *myasthurts.Field
	CurrentTag   myasthurts.TagParam
	CurrentRef   string

	ActualField *myasthurts.Field
	ActualTag   myasthurts.TagParam
	ActualRef   string

	Expressions []string
}

// EqInput TODO
type EqInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// RequiredInput TODO
type RequiredInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// Condition TODO
var Condition = make(map[string]string, 10)

// Expression TODO
var Expression string
