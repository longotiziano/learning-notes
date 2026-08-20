#include <stdio.h>

void calcular_seg_tiempo(int s_ingresados, int *days, int *hours, int *minutes, int *segs) {
    int cont_min=0, cont_hs=0, cont_day=0; 

    if (s_ingresados >= 0) {
        while (s_ingresados >= 60) {
            cont_min++;
            s_ingresados -= 60;
        };
        while (cont_min >= 60) {
            cont_hs++;
            cont_min -= 60;
        };
        while (cont_hs >= 24) {
            cont_day++;
            cont_hs -= 24;
        };
        *days = cont_day;
        *minutes = cont_min;
        *hours = cont_hs;
        *segs = s_ingresados;
        }
}

void main() {
    int seg_ingresados, days, hours, minutes, segs;
    printf("\nIngrese una cantidad de segundos: ");
    scanf("%d", &seg_ingresados);
    if (seg_ingresados >= 0) {
        calcular_seg_tiempo(seg_ingresados, &days, &hours, &minutes, &segs);
        printf("\nPara los segundos %d -> \nDias: %d\nHours: %d\nMinutes: %d\nSegundos: %d", seg_ingresados, days, hours, minutes, segs);
    }
    else prinft("\nValor ingresado inválido");
}