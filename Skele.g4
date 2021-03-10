parser grammar Skele;

options { tokenVocab=SkeleLexer; }

start   : spec+ EOF;

spec    : pkg fol? doc? file+;

fol     : FOL WORD NEWLINE;

pkg     : PKG WORD NEWLINE;

doc     : DOC NEWLINE ln+;

file    : FILE FILENAME NEWLINE fun+;

fun     : FUN NEWLINE ln pre? pos;

pre     : PRE NEWLINE ln+;
pos     : POS NEWLINE ln+;

ln      : LINE NEWLINE;
