parser grammar Skele;

options { tokenVocab=SkeleLexer; }

start   : spec+ EOF;

spec    : pkg fol? doc? file+;

fol     : FOL WORD newln;

pkg     : PKG WORD newln;

doc     : DOC newln ln+;

file    : FILE FILENAME newln fun+;

fun     : FUN newln ln pre? pos;

pre     : PRE newln ln+;
pos     : POS newln ln+;

ln      : LINE newln;
newln   : (comment? NEWLINE)+;

comment : COMMENT;
