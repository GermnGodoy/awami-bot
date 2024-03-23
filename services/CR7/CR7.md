# The CR7 format

CR7 (Chess Representation 7) is one of the multiple ways of encoding chess positions. This is especially usefull when we talk about making petitions to different api, such as the ones we are using in our program. The format has three main parts:

- **Board:** The first characters of the CR7 string are the ones that have the informations of the positions of the pieces in the board. This information is represented with the starting letter of the piece name (ex. "q" would be the black queen and "K" would be the white king). In this part there are number that represent the blank spaces too.
- **En pasant:** Two digits number that contain the position of a pawn we can eat with en pasant (if none it is 00).
- **Castling state:** A pair of four binary numbers that representa if white team (first two numbers) and black team (second two numbers) has castled already. 
- **Movement location seed:** The seed that has been set for the match in the server, for checking repetition information.