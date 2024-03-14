# The CR7 format

CR7 (Chess Representation 7) is one of the multiple ways of encoding chess positions. This is especially usefull when we talk about making petitions to different api, such as the ones we are using in our program. The format has three main parts:

- **Board:** The first characters of the CR7 string are the ones that have the informations of the positions of the pieces in the board. This information is represented with the starting letter of the piece name (ex. "q" would be the black queen and "K" would be the white king). In this part there are number that represent the blank spaces too.
- **Castling state:** A pair of two binary numbers that representa if white team (first number) and black team (second number) has castled already. 
- **Repetition state:** A number and two algebraic notated chess movements. The number represents the repetitions of the movements that have been made, and the chess movements (that would be 8 characters, the first four for white moves and the next ones for the black moves) are the movements we are tracking. The movements are set to zero if there is no repetition of it is the first turn in the game.