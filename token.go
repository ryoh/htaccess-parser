package main

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT

	// Misc
	ASTERISK // *
	COMMA    // ,

	// Keywords
	SELECT
	FROM
)
