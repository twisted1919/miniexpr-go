package miniexpr

/*
BNF Grammar:

<expression> ::= <term>
<term>       ::= <factor> (("+" | "-") <factor>)*
<factor>     ::= <unary> (("*" | "/" | "<<" | ">>") <unary>)*
<unary>      ::= ("-" | "+") <unary> | <pow>
<pow>        ::= <primary> ( ("**" | "^") <pow> )?
<primary>    ::= [0-9]+ | "(" <expression> ")"

Notes:
* is actually a regular for loop
? is actually a regular if statement
*/
