#include <stdio.h> // libreria para manejar Standar Input/Output

void main()
{
    float temperatura, celsius, farenheit;
    char medida_elegida;

    int formula = celsius;

    printf("Ingrese una temperatura...");
    scanf("%f", &temperatura);
    printf("Ingrese una medida: ");
    scanf("%c", &medida_elegida);
    if (medida_elegida == 'F') {
        celsius = 5.0 / 9.0 * (temperatura - 32);
        printf("La temperatura en celsius es: %f", celsius);
    }
    if (medida_elegida == 'C') {
        farenheit = (celsius * 9.0 / 5.0) + 32;
        printf("La temperatura en farenheit es: %f", farenheit);
    }
    if (medida_elegida !='C' && medida_elegida !='F') {
        printf("No se ha encontrado una medida para el carácter ingresado.");
    }
}