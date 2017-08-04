// Code generated by yy. DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
)

// CommaOpt represents data reduced by productions:
//
//	CommaOpt:
//	        /* empty */
//	|       ','          // Case 1
type CommaOpt struct {
	Token Token
}

func (n *CommaOpt) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *CommaOpt) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *CommaOpt) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// ExprList represents data reduced by productions:
//
//	ExprList:
//	        Expression
//	|       ExprList ',' Expression  // Case 1
type ExprList struct {
	Case       int
	ExprList   *ExprList
	Expression *Expression
	Token      Token
}

func (n *ExprList) reverse() *ExprList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.ExprList
	for nb != nil {
		nc := nb.ExprList
		nb.ExprList = na
		na = nb
		nb = nc
	}
	n.ExprList = nil
	return na
}

func (n *ExprList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *ExprList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *ExprList) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 1:
		return n.ExprList.Pos()
	case 0:
		return n.Expression.Pos()
	default:
		panic("internal error")
	}
}

// Expression represents data reduced by productions:
//
//	Expression:
//	        '!' Expression
//	|       '-' Expression               // Case 1
//	|       Expression "&&" Expression   // Case 2
//	|       Expression "->" Expression   // Case 3
//	|       Expression "/." Expression   // Case 4
//	|       Expression "//" Expression   // Case 5
//	|       Expression "//." Expression  // Case 6
//	|       Expression "/;" Expression   // Case 7
//	|       Expression "/@" Expression   // Case 8
//	|       Expression ":=" Expression   // Case 9
//	|       Expression ":>" Expression   // Case 10
//	|       Expression "<=" Expression   // Case 11
//	|       Expression "<>" Expression   // Case 12
//	|       Expression "=!=" Expression  // Case 13
//	|       Expression "==" Expression   // Case 14
//	|       Expression "===" Expression  // Case 15
//	|       Expression ">=" Expression   // Case 16
//	|       Expression "@@" Expression   // Case 17
//	|       Expression "||" Expression   // Case 18
//	|       Expression '*' Expression    // Case 19
//	|       Expression '+' Expression    // Case 20
//	|       Expression '-' Expression    // Case 21
//	|       Expression '.' Expression    // Case 22
//	|       Expression '/' Expression    // Case 23
//	|       Expression ':' Expression    // Case 24
//	|       Expression ';'               // Case 25
//	|       Expression ';' Expression    // Case 26
//	|       Expression '<' Expression    // Case 27
//	|       Expression '=' Expression    // Case 28
//	|       Expression '>' Expression    // Case 29
//	|       Expression '?' Expression    // Case 30
//	|       Expression '@' Expression    // Case 31
//	|       Expression '^' Expression    // Case 32
//	|       Expression '|' Expression    // Case 33
//	|       Factor                       // Case 34
//	|       Factor ':' Expression        // Case 35
type Expression struct {
	Case        int
	Expression  *Expression
	Expression2 *Expression
	Factor      *Factor
	Token       Token
}

func (n *Expression) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Expression) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Expression) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33:
		return n.Expression.Pos()
	case 34, 35:
		return n.Factor.Pos()
	case 0, 1:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// Factor represents data reduced by productions:
//
//	Factor:
//	        Term
//	|       Term Factor  // Case 1
type Factor struct {
	Case   int
	Factor *Factor
	Term   *Term
}

func (n *Factor) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Factor) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Factor) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Term.Pos()
}

// Tag represents data reduced by productions:
//
//	Tag:
//	        IDENT
//	|       STRING  // Case 1
type Tag struct {
	Case  int
	Token Token
}

func (n *Tag) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tag) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tag) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// Term represents data reduced by productions:
//
//	Term:
//	        FLOAT
//	|       '(' Expression ')'                // Case 1
//	|       '{' '}'                           // Case 2
//	|       '{' ExprList CommaOpt '}'         // Case 3
//	|       IDENT                             // Case 4
//	|       IDENT "::" Tag                    // Case 5
//	|       IDENT "::" Tag "::" Tag           // Case 6
//	|       INT                               // Case 7
//	|       PATTERN                           // Case 8
//	|       SLOT                              // Case 9
//	|       STRING                            // Case 10
//	|       Term "[[" ExprList CommaOpt "]]"  // Case 11
//	|       Term '!'                          // Case 12
//	|       Term '&'                          // Case 13
//	|       Term '[' ']'                      // Case 14
//	|       Term '[' ExprList CommaOpt ']'    // Case 15
type Term struct {
	Case       int
	CommaOpt   *CommaOpt
	ExprList   *ExprList
	Expression *Expression
	Tag        *Tag
	Tag2       *Tag
	Term       *Term
	Token      Token
	Token2     Token
	Token3     Token
}

func (n *Term) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Term) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Term) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 11, 12, 13, 14, 15:
		return n.Term.Pos()
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// start represents data reduced by production:
//
//	start:
//	        Expression
type start struct {
	Expression *Expression
}

func (n *start) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *start) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *start) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Expression.Pos()
}
