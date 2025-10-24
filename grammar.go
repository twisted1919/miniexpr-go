package miniexpr

/*
BNF Grammar:

<expression> ::= <term>
<term>       ::= <factor> (("+" | "-") <factor>)*
<factor>     ::= <unary> (("*" | "/" | "<<" | ">>") <unary>)*
<unary>      ::= ("-" | "+") <unary> | <pow>
<pow>        ::= <primary> ( ("**" | "^") <pow> )?
<primary>    ::= NUMBER | "(" <expression> ")"

Notes:
NUMBER is actually [0-9]+ where . and _ are allowed.
* is actually a regular for loop
? is actually a regular if statement
*/
