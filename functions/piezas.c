#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>


typedef struct _move {
     int start;
     int finish;
} move;

// Definicon de cada nodo de la linked list (utilizada para guardar los posibles movimientos).
typedef struct node {
     move move;
     struct node *next;
} node;


// Por ahora no lo uso, es una idea para en caso de tener la posicion de la pieza, poder saber que pieza es.
//typedef struct _hash{
//     int square;
//     int *ptrtopiece;
//} hash;
// hash hashtablero[64];

int tablero[64] = {-5, -3, -4, -9, -20, -4, -3, -5,
	-1, -1, -1, -1, -1, 0, -1, -1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, -1, 0, 0,
	0, 3, 0, 1, 0, 0, 9, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 0, 0, 1, 1, 1,
	5, 0, 0, 0, 20, 4, 3, 5};

move lastmove;
// Nos conviene saber en todo momento cual es la posicion del rey.
int wkingpos = 60;
// Defino el pointer a la lista de movimientos como varaible global.
node *listofmoves;


// En caso de que algun movimiento prohiba el enroque el el futuro, pasan a tomar el valor 0.
uint8_t enroquelargoblanco = 1;
uint8_t enroquecortoblanco = 1;
uint8_t enroquecortonegro = 1;
uint8_t enroquelargonegro = 1;


//funciones utilizadas
int ischeck(move movimiento);
void push(move movimiento);
void makemove(move);
void possiblemoves(void);
void bishopmoves(int i);
void rookmoves(int i);
void pawnmoves(int i);
void horsemoves(int i);
void kingmoves(int i);
void castling(void);


int main(void){
     // Ahora mismo no lo utilizamos, es util para rellenar los valores en caso de utilizar el hashmap.
     //for (int i = 0; i++; i<64){
     //     hashtablero[i].square = i;
     //     hashtablero[i].ptrtopiece = &tablero[i];
     //}

     listofmoves = malloc(sizeof(node));
     // Solo por precaucion, en caso de que malloc devuelva NULL abortamos el programa.
     if (listofmoves == NULL){
        return 1;
     }
     (*listofmoves).next = NULL;

     possiblemoves();
     node *ptr = listofmoves;
     while ((*ptr).next != NULL){
          printf("[%i, ", (*ptr).move.start);
          printf("%i]\n", (*ptr).move.finish);
          node *tmp = ptr;
          ptr = (*ptr).next;
          free(tmp);
     }
}

void push(move movimiento){
  // Se mete un nuevo nodo, en el primer lugar de la linked list.
     node *newelement = malloc(sizeof(node));
     (*newelement).move = movimiento;
     (*newelement).next = listofmoves;
     listofmoves = newelement;
}

//Funcion para realizar movimientos.
void makemove(move movement){
    // En passant a la izquierda.
    if (movement.finish == -1){
        tablero[movement.start - 9] = 1;
        tablero[movement.start - 1] = 0;
        tablero[movement.start] = 0;
        return;
    }
    // En passant izquierda inverso.
    if (movement.start == -1){
        tablero[movement.finish] = 1;
        tablero[movement.finish - 1] = -1;
        tablero[movement.finish - 9] = 0;
        return;
    }

    // En passant a la derecha.
    if (movement.finish == -2){
        tablero[movement.start - 7] = 1;
        tablero[movement.start + 1] = 0;
        tablero[movement.start] = 0;
        return;
    }
    if (movement.start == -2){
        tablero[movement.finish] = 1;
        tablero[movement.finish + 1] = -1;
        tablero[movement.finish - 7] = 0;
        return;
    }
    // Enroque corto.
    if (movement.finish == -3){
        tablero[60] = 0;
        tablero[61] = 5;
        tablero [62] = 20;
        tablero[63] = 0;
        enroquelargoblanco = 0;
        enroquecortoblanco = 0;
        return;
    }
    //Enroque corto inverso.
    if (movement.start == -3){
        tablero[60] = 20;
        tablero[61] = 0;
        tablero [62] = 0;
        tablero[63] = 5;
        enroquelargoblanco = 1;
        enroquecortoblanco = 1;
        return;
    }

    //Enroque largo.
    if (movement.finish == -4){
        tablero[60] = 0;
        tablero[58] = 5;
        tablero[57] = 20;
        tablero[56] = 0;
        enroquelargoblanco = 0;
        enroquecortoblanco = 0;
        return;
    }

    //Enroque largo inverso.
    if (movement.finish == -4){
        tablero[60] = 20;
        tablero[58] = 0;
        tablero[57] = 0;
        tablero[56] = 5;
        enroquelargoblanco = 1;
        enroquecortoblanco = 1;
        return;
    }

    //Movimientos normales.
     tablero[movement.finish] = tablero[movement.start];
     tablero[movement.start] = 0;
     lastmove = movement;

     if (movement.start == 59 ){
          enroquecortoblanco = 0;
          enroquelargoblanco = 0;
     }
     if (movement.start == 63 || movement.finish == 63){
          enroquecortoblanco = 0;
     }
     if (movement.start == 56 || movement.finish == 56){
          enroquelargoblanco = 0;
     }
     return;
}

//Funcion que calcula todos los movimientos legales.
void possiblemoves(void){
     int i = 0;
     //Definimos el pointer del ultimo elemento de la lista a NULL, nos servira para saber cuando termina la linked list.
     move movlegal;

     //chequeamos primero posibles ernoques.
     castling();

        while (i < 64){
        //printf("%i\n", tablero[i]);
        if (tablero[i] == 0){
            i++;
        }
        // Tengo que inicializar las variables aca, no me deja dentro del switch.
        movlegal.start = i;
        switch(tablero[i])
        {
        //Caso peon blanco
        case 1:
        pawnmoves(i);
        break;
        // Caso caballo blanco.
        case 3:
        horsemoves(i);
        break;
        // Caso alfil.
        case 4:
        bishopmoves(i);
        break;
        // Caso torre.
        case 5:
             rookmoves(i);
             //El enroque se hace a aparte (al principio).
            break;
        // Caso reina.
        case 9:
              bishopmoves(i);
              rookmoves(i);
            break;
        // Caso rey.
        case 20:
              kingmoves(i);
              //El enroque se hace a aparte (al principio).
           break;

        default:
            break;
        }
        i++;
     }
}


void castling(void){
    move movlegal;
    movlegal.start = 60;
     if (enroquecortoblanco == 1){
        if (tablero[61] == 0 && tablero[62] == 0){
            movlegal.start = 60;
            movlegal.finish = 60;
            if (ischeck(movlegal)){
                movlegal.finish = 61;
                if (ischeck(movlegal)){
                    movlegal.finish = 62;
                    if (ischeck(movlegal)){
                        movlegal.finish = 63;
                        if (ischeck(movlegal)){
                            movlegal.finish = -3;
                            push(movlegal);
                        }
                    }
                }
            }
        }
     }
    if (enroquelargoblanco == 1){
        if (tablero[59] == 0 && tablero[58] == 0 && tablero[57] == 0){
            movlegal.start = 60;
            movlegal.finish = 60;
            if (ischeck(movlegal)){
                movlegal.finish = 59;
                if (ischeck(movlegal)){
                    movlegal.finish = 58;
                    if (ischeck(movlegal)){
                        movlegal.finish = 57;
                        if (ischeck(movlegal)){
                            movlegal.finish = 56;
                            if(ischeck(movlegal)){
                                movlegal.finish = -4;
                                push(movlegal);
                            }
                        }
                    }
                }
            }
        }
     }
}
void kingmoves(int i){
            move movlegal;
            movlegal.start = i;
            if ((i+1) % 8 !=0 && tablero[i+1] <= 0){
            movlegal.finish = i + 1;
            if (ischeck(movlegal)){
            push(movlegal);
            }
            }
            if ((i-1) % 8 != 7 && tablero[i+1] <= 0){
                movlegal.finish = i + 1;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i-8 >=0 && tablero[i-8] <= 0){
                movlegal.finish = i - 8;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i+8 < 64 && tablero[i+8] <= 0){
                movlegal.finish = i + 8;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i+8 < 64 && (i-1) % 8 != 0 && tablero[i+7] <= 0){
                movlegal.finish = i + 7;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i-8 >=0 && (i-1) % 8 != 0 && tablero[i-9] <= 0){
                movlegal.finish = i - 9;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i+8 < 64 && (i+1) % 8 != 7 && tablero[i+9] <= 0){
                movlegal.finish = i + 9;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i-8 >= 0 && (i+1) % 8 != 7 && tablero[i-7] <= 0){
                movlegal.finish = i - 7;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

}



void pawnmoves(int i){
            move movlegal;
            movlegal.start = i;
            if (tablero[i-8] == 0){
                movlegal.finish = i - 8;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i < 56 && i > 47 && tablero [i-16] == 0 && tablero[i-8] == 0){
                movlegal.finish = i - 16;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            //En passant
            if (23 < i && i < 32 && tablero[lastmove.finish] == -1){
                if (lastmove.finish == i - 1 && i % 8 != 0){
                    movlegal.finish = -1;
                    if (ischeck(movlegal)){
                    push(movlegal);
                }
                }
                if (lastmove.finish == i + 1 && i % 8 != 7){
                    movlegal.finish = -2;
                    if (ischeck(movlegal)){
                    push(movlegal);
                }
                }
            }

            if (i % 8 != 7 && tablero[i-7] < 0){
                movlegal.finish = i - 7;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }
            if (i % 8 != 0 && tablero[i-9] < 0){
                movlegal.finish = i - 7;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

}
void horsemoves(int i){
    move movlegal;
    movlegal.start = i;
    if (i > 15){
                if (tablero [i-15] <=0 && i % 8 != 7){
                        movlegal.finish = i - 15;
                        if (ischeck(movlegal)){
                        push(movlegal);
                        }
                }
                if (tablero [i-17] <=0 && i % 8 != 0){
                        movlegal.finish = i - 17;
                        if (ischeck(movlegal)){
                        push(movlegal);
                        }
                }
            }
            if (i > 7){
                if (tablero [i-6] <=0 && i % 8 < 6){
                        movlegal.finish = i - 6;
                        if (ischeck(movlegal)){
                         push(movlegal);
                         }
                }
                if (tablero [i-10] <=0 && i % 8 > 1){
                        movlegal.finish = i - 10;
                        if (ischeck(movlegal)){
                         push(movlegal);
                        }
                }
            }
            if (i < 48){
                if (tablero [i+15] <=0 && i % 8 != 0){
                        movlegal.finish = i + 15;
                        if (ischeck(movlegal)){
                         push(movlegal);
                         }
                }
                if (tablero [i+17] <=0 && i % 8 != 7){
                        movlegal.finish = i + 17;
                        if (ischeck(movlegal)){
                          push(movlegal);
                        }
                }
            }
            if (i < 56){
                if (tablero [i+6] <=0 && i % 8 > 1){
                        movlegal.finish = i + 6;
                        if (ischeck(movlegal)){
                         push(movlegal);
                        }
                }
                if (tablero [i+10] <=0 && i % 8 < 6){
                        movlegal.finish = i + 10;
                        if (ischeck(movlegal)){
                         push(movlegal);
                        }
                }
            }

}


//Funcion que chequea si el rey esta en jaque.
int ischeck(move movimiento){
     move inversemove;
     inversemove.start = movimiento.finish;
     inversemove.finish = movimiento.start;
     makemove(movimiento);
     int j = wkingpos - 9;
     while (j >= 9 && j % 8 != 0 && tablero [j] == 0){
          if (tablero[j - 9] == -4 || tablero[j - 9] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j - 9;
     }
     j = wkingpos + 9;
     while (j < 55 && j % 8 != 7 && tablero [j] == 0){
          if (tablero[j + 9] == -4 || tablero[j + 9] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j + 9;
     }


     j = wkingpos - 7;
     while (j > 7 && j % 8 != 7 && tablero [j] == 0){
          if (tablero[j - 9] == -4 || tablero[j - 9] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j - 7;
     }
     j = wkingpos + 7;
     while (j < 56 && j % 8 != 0 && tablero [j] == 0){
          if (tablero[j + 9] == -4 || tablero[j+9] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j + 7;
     }
     j = wkingpos + 1;
     while (j % 8 != 7 && tablero [j] == 0){
          if (tablero[j + 1] == -5 || tablero[j + 1] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j + 1;
     }
     j = wkingpos - 1;
     while (j % 8 != 0 && tablero [j] == 0){
          if (tablero[j - 1] == -5 || tablero[j - 1] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j - 1;
     }
     j = wkingpos + 8;
     while (j < 56 && tablero [j] == 0){
          if (tablero[j + 8] == -5 || tablero[j + 8] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j + 8;
     }
     j = wkingpos - 8;
     while (j > 7 && tablero [j] == 0){
          if (tablero[j - 8] == -5 || tablero[j - 8] == -9){
               makemove(inversemove);
               return 0;
          }
          j = j - 8;
     }
     if (wkingpos > 7)
     {
          if (wkingpos % 8 != 7 && tablero[wkingpos -7] == -1){
               makemove(inversemove);
               return 0;
          }
          if (wkingpos % 8 != 0 && tablero[wkingpos -9] == -1){
               makemove(inversemove);
               return 0;
          }
     }
     if (wkingpos > 15){
                if (tablero [wkingpos - 15] == -3 && wkingpos % 8 != 7){
                    makemove(inversemove);
                    return 0;
                }
                if (tablero [wkingpos - 17] == -3 && wkingpos % 8 != 0){
                    makemove(inversemove);
                    return 0;
                }
            }
     if (wkingpos > 7){
                if (tablero [wkingpos - 6] == -3 && wkingpos % 8 < 6){
                    makemove(inversemove);
                    return 0;
                }
                if (tablero [wkingpos-10] == -3 && wkingpos % 8 > 1){
                    makemove(inversemove);
                    return 0;
                }
            }
            if (wkingpos < 48){
                if (tablero [wkingpos + 15] == -3 && wkingpos % 8 != 0){
                    makemove(inversemove);
                    return 0;
                }
                if (tablero [wkingpos + 17] == -3 && wkingpos % 8 != 7){
                    makemove(inversemove);
                    return 0;
                }
            }
            if (wkingpos < 56){
                if (tablero [wkingpos + 6] == -3 && wkingpos % 8 > 1){
                    makemove(inversemove);
                    return 0;
                }
                if (tablero [wkingpos+10] == -3 && wkingpos % 8 < 6){
                    makemove(inversemove);
                    return 0;
                }
            }

makemove(inversemove);
return 1;
}

void bishopmoves(int i){
            move movlegal;
            movlegal.start = i;
            int j = i + 9;

            while (j >= 0 && j < 64 && j % 8 != 0 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }

           j = j + 9;
            }
            if (j >= 0 && j < 64 && j % 8 != 0 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i - 9;
            while (j >= 0 && j < 64 && j % 8 != 7 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j - 9;
            }
            if (j >= 0 && j < 64 && j % 8 != 7 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i + 7;
            while (j >= 0 && j < 64 && j % 8 != 0 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j + 7;
            }
            if (j >= 0 && j < 64 && j % 8 != 0 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i - 7;
            while (j >= 0 && j < 64 && j % 8 != 7 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j - 7;
            }
            if (j >= 0 && j < 64 && j % 8 != 7 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

}
void rookmoves(int i){
            move movlegal;
            movlegal.start = i;
            int j = i + 1;

            while (j % 8 != 0 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j++;
            }
            if (j % 8 != 0 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i - 1;
            while (j % 8 != 7 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j - 1;
            }
            if (j % 8 != 7 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i + 8;
            while (j < 64 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j + 8;
            }
            if (j < 64 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
            }

            j = i - 8;
            while (j >= 0 && tablero[j] == 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                push(movlegal);
                }
                j = j - 8;
            }
            if (j >= 0 && tablero[j] < 0){
                movlegal.finish = j;
                if (ischeck(movlegal)){
                 push(movlegal);
                }
            }

}


