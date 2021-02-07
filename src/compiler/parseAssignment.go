/*
parseAssignment.go
Author: @BraydonKains
*/

package compiler

import (
	"errors"
)

func parseAssignment(p *Parser) error {
	switch token := p.peekToken(); token.id {
	case TYPEINT:
		return parseIntAssignment()
	case TYPESTR:
		return parseStringAssignment()
	case TYPEBOOL:
		return parseBoolAssignment()
	}
	return errors.New("Cannot parse assignment for this type yet")
}

func parseIntAssignment(p *Parser) error {
	typeToken := p.readToken()

	if identifier := p.readToken(); identifier.id != IDENT {
		return errors.New("Missing identifier")
	}

	if assignment := p.readToken(); assignment.id != ASSIGN {
		return errors.New("Expected assignment operator")
	}

	value := p.readToken()
	negativeFlag := false
	if value.id == SUB {
		negativeFlag = true
		value = p.readToken()
	}
	if value.id != INT {
		return errors.New("Tried to assign non-integer value to int")
	}

	return errors.New("Int assignments are not implemented")
}

func parseStringAssignment(p *Parser) error {
	return errors.New("String assignments are not implemented")
}

func parseBoolAssignment(p *Parser) error {
	return errors.New("String assignments are not implemented")
}
