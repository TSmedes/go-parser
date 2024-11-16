# Lexical and Syntax analysis for 4Point


This project consists of the front end of a compiler. The lexer and parser are written in Go and are the front end of the compiler for 4Point, a made-up programming language. The lexer creates a token table that is then given to the parser to process for the syntax analysis. The parser uses a [recursive descent](https://en.wikipedia.org/wiki/Recursive_descent_parser) parsing method. The grammar for 4Point can be seen below.

## Grammar

```
START      --> STMT_LIST
STMT_LIST  --> STMT. |
               STMT; STMT_LIST
STMT       --> POINT_DEF |
               TEST
POINT_DEF  --> ID = point(NUM, NUM)
TEST       --> test(OPTION, POINT_LIST)
ID         --> LETTER+
NUM        --> DIGIT+
OPTION     --> triangle |
               square
POINT_LIST --> ID |
               ID, POINT_LIST
LETTER     --> a | b | c | d | e | f | g | ... | z
DIGIT      --> 0 | 1 | 2 | 3 | 4 | 5 | 6 | ... | 9

```

## Running the program

After cloning the repo, run the program with this command, inside the project directory:
   <pre>
      <b>go run .</b> [file] [-p | -s]</pre>
      
   **-p** &nbsp;&nbsp;&nbsp;*Prolog*
   
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Produces output code in prologwhich will have a series of queries about those four points.
   
   **-s** &nbsp;&nbsp;&nbsp;*Scheme*
   
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The program will output function calls in Scheme that is going to be called by a program in Scheme that will calculate properties of those four points.

   **file** &nbsp;&nbsp;&nbsp;*Input file*
   
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The file containing 4Point code to be parsed.

 ## Testing

 There are several files with test input code. These can be used as input files, some have errors which will be detected at compile time.
