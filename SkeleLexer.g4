lexer grammar SkeleLexer;

PKG     : 'pkg';
FOL     : 'fol';
DOC     : 'doc';
FILE    : 'file';
FUN     : 'fun';
PRE     : 'pre';
POS     : 'pos';

LINE : '\\' ~[\r\n]*;

FILENAME: WORD '_test.go';

WORD: ID+;

NEWLINE : ('\r'? '\n' | '\r')+ ;

FS : '/';

WS : [ \n\t\r] -> skip;
ID : [a-zA-Z];