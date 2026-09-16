package maze

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dimStr := scanner.Text()
		filas, cols := ObtenerDimension(dimStr)

		// en caso de que falle alguna vez, seguira leyendo la entrada estándar hasta encontrar el formato válido
		if filas > 0 && cols > 0 {
			lab := CrearLaberinto(filas, cols, AlmacenarCasillas(scanner, filas, cols))
			direcs, cantPasos, err := ResolverLaberinto(lab)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("%d\n", cantPasos)
			for _, dir := range direcs {
				fmt.Printf("%s ", dir)
			}
			fmt.Printf("\n")
		}
	}
}
