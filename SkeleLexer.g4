lexer grammar SkeleLexer;

PKG     : 'pkg';
FOL     : 'fol';
DOC     : 'doc';
FILE    : 'file';
FUN     : 'fun';
PRE     : 'pre';
POS     : 'pos';
FSM     : 'fsm';
STS     : 'states';

LINE : '\\' ~[\r\n#]*;
COMMENT : '#' ~[\r\n]*;

FILENAME: ID (ID | [0-9] | '_')* '_test.go';

WORD: ID+ (ID | '_')*;

NEWLINE : ('\r'? '\n' | '\r')+ ;

WS : [ \n\t\r] -> skip;
ID : [a-zA-Z];