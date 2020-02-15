package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, stmt := range program.Statements {
		if isMacroDefinition(stmt) {
			addMacro(stmt, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		index := definitions[i]
		program.Statements = append(
			program.Statements[:index],
			program.Statements[index+1:]...,
		)
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStmt, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStmt.Value.(*ast.MacroLiteral)
	return ok
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStmt, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStmt.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Body:       macroLiteral.Body,
		Env:        env,
	}

	env.Set(letStmt.Name.Value, macro)
}
