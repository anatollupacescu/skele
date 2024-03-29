parser grammar Skele;

options { tokenVocab=SkeleLexer; }

start   : spec+ EOF;

spec    : (pkg fol? doc) | (pkg fol? doc? file+);

fsm     : FSM newln ln sts;

sts     : STS newln ln+;

fol     : FOL WORD newln;

pkg     : PKG WORD newln;

doc     : DOC newln ln+;

file    : FILE FILENAME newln fsm* fun+;

fun     : FUN newln ln pre? pos;

pre     : PRE newln ln+;
pos     : POS newln ln+;

ln      : LINE newln;
newln   : (comment? NEWLINE)+;

comment : COMMENT;
