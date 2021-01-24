parser grammar Skele;

options { tokenVocab=SkeleLexer; }

start : spec+ EOF;

spec: file fun+;

file    : FILE WORD FS FILENAME NEWLINE;

fun     : FUN WORD ARG_NAME CP COMMENT NEWLINE+ pre* pos;

pre     : (PRE WORD COMMENT | PRE (NEWLINE ml)+) NEWLINE;
pos     : (POS WORD COMMENT | POS (NEWLINE ml)+) NEWLINE;

ml      : ML WORD COMMENT;