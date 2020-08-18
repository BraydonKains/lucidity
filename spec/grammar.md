# Grammar
### Author: @BraydonKains

Mostly for my own benefit, I'm using this document to track my efforts at drafting a grammar for my language.

## Plain List of Reserved Words and Symbols

This will likely be an evolving list of the reserved words for the language so far.

### Condition control
*if*
* For conditional branching statements
*unless*
* The negative of if
*else*
* If the 1 or more previous conditional check(s) have failed
* TODO: How am I collecting those if checks?
*and*, *or*, *all*, *one_of*
* boolean logic operators
*equals*, *greater*, *less*, *than*
* Comparison operators
*is_a*, *is_an*
* Type checking operators

### Loops
*for*
* begin a loop
*to*
* begins a range loop
*step*
* Specify a step amount in a range loop 
* TODO: Allow a function to be used to calculate step?
*each*
* begins an element loop

### Types
*int*
* Integer number
*float*
* Floating point number
*string*
* String of text 
*mix*
* A type that is completely unchecked

## Context Free Grammar

Notes: 
* All rules are uppercase, all lower case will be reserved words
* Whitespace in a rule is significant

Iden -> [a-Z]([a-Z]|[0-9])\*
Expr -> Expr(OpExpr)\*
Op -> (+|-|\*|/)
LP -> 
Code -> My bad way of saying "any amount of valid code"
ConditionStatement -> (if|unless) ConditionExpr {Code}
ConditionExpr -> (Iden|BoolExpr)
BoolExpr -> (CompareExpr|TypeCheckExpr|BoolExpr)(BoolOp BoolExpr)*
CompareExpr -> (Iden|Expr) (
