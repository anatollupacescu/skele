lexer grammar SkeleLexer;

FUN: 'fun';
FILE: 'file';
ITEM: 'item';

PRE: 'pre';
POS: 'pos';

COMMENT : '#' ~[\r\n]*;
ARG_NAME: '(' ~[)]*;

FILENAME: WORD '.go';

WORD: (ID | '_')+;

NEWLINE : ('\r'? '\n' | '\r')+ ;

CP : ')';
FS : '/';
DT : '.';
ML: '\\';

WS : [ \n\t\r] -> skip;
ID : [a-zA-Z];