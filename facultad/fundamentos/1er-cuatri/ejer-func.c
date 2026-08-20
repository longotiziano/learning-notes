#include <stdio.h>

int calcular_factorial(int x) {
    int fact_x, i;
    fact_x = x;

    if (x == 0) return 1;
    
    for (i=1; abs(x)>i; i++)
        fact_x=fact_x*i; 
    return fact_x;
}

void main() {
    int x = -8;
    int y = 0;
    int z = 1;
    int w = 5;

    int factorialx = calcular_factorial(x);
    int factorialy = calcular_factorial(y);
    int factorialz = calcular_factorial(z);
    int factorialw = calcular_factorial(w);

    printf("%d", factorialx);
    printf("%d", factorialy);
    printf("%d", factorialz);
    printf("%d", factorialw);
}