/*
* Dadas las declaraciones de constantes y tipos.
* Desarrollar una funcion que reciba como parametros un vector de tipo T_Vector ya cargado y su maximo logico, y devuelva el arreglo invertido y la cantidad de intercambios realizados. Cumpliendo las siguientes condiciones:    
    - La inversion debe ser sobre si mismo, sin vector auxiliar.
    - Si las posiciones a invertir contienen el mismo valor, no deben ser intercambiadas
*/
#include <stdio.h>
#define MAX 100
typedef int T_Vector[MAX];

void invertirVector(T_Vector vec, int maxl, int *cantInter) {
    int numOrigen, numDestino, i;
    *cantInter = 0;
    i = 0;
    
    for (i=0; i<(maxl/2); i++) {
        numOrigen = vec[i];
        printf("el numero origen es %d\n", numOrigen);
        numDestino = vec[maxl -1 - i];
        if (numOrigen != numDestino) {
            vec[i] = numDestino;
            vec[maxl -1 - i] = numOrigen;
            *cantInter += 1;
        }
    }
}

int main() {
    T_Vector vec = {6, 2, 3, 4, 4, 5, 6};
    int maxl = 7;
    int cantInter;
    
    invertirVector(vec, maxl, &cantInter);
    printf("Intercambios realizados: %d\n", cantInter);
    printf("El vector invertido es: ");
    for (int i = 0; i < maxl; i++) {
        printf("%d ", vec[i]);
    }
    return 0;
}
