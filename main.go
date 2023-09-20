package main

import "fmt"

func main() {
	// 2
	// fmt.Println(primoNoPrimo(13))

	// 3
	// arreglo := [5]int{2, 4, 6, 7, 9}
	// fmt.Println(busquedaSecuencial(5, arreglo[:]))

	// 4
	// fmt.Println(busquedaBinaria(3, arreglo[:]))

	// 5
	// matriz := [5][5]int{}
	// fmt.Println(matriz)
	// matriz = inicializarACero(matriz)
	// fmt.Println(matriz)

	// 6
	// matriz := [5][5]int{}
	// matriz = inicializarACeroSuperior(matriz)
	// for i := 0; i < len(matriz); i++ {
	// 	fmt.Println(matriz[i])
	// }

	// 7
	// matrizA := [2][2]int{{1, 2}, {3, 4}}
	// matrizB := [2][2]int{{2, 0}, {1, 2}}
	// fmt.Println(multiplicarMatriz(matrizA, matrizB))

	// 8
	// arreglo := []int{64, 34, 25, 12, 22, 11, 90}
	// arreglo := []int{11, 22, 12, 25, 34, 64, 90}
	// arreglo = burbujaBandera(arreglo)
	// fmt.Println(arreglo)

}

// problem 1
func parNoPar(x int) {
	if x%2 == 0 {
		fmt.Printf("%d es par", x)
	} else {
		fmt.Printf("%d no es par", x)
	}
}

// problem 2
func primoNoPrimo(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

// problem 3
func busquedaSecuencial(x int, m []int) bool {
	for i := 0; i < len(m); i++ {
		if m[i] == x {
			return true
		}
	}

	return false
}

// problem 4
func busquedaBinaria(x int, m []int) int {
	izquierda, derecha := 0, len(m)-1

	for izquierda <= derecha {
		medio := (izquierda + derecha) / 2

		if x == m[medio] {
			return medio // retorna la ubicación
		} else if x > m[medio] {
			izquierda = medio + 1
		} else {
			derecha = medio - 1
		}

	}
	return -1 // no se encontro el valor
}

// problem 5
func inicializarACero(matriz [5][5]int) [5][5]int {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			matriz[i][j] = 1
		}
	}

	return matriz
}

// problem 6
func inicializarACeroSuperior(matriz [5][5]int) [5][5]int {
	for i := 0; i < len(matriz); i++ {
		for j := 1; j < len(matriz); j++ {
			if j > i {
				matriz[i][j] = 1
			}
		}
	}

	return matriz
}

// problema 7
func multiplicarMatriz(matrizA [2][2]int, matrizB [2][2]int) [2][2]int {
	resultado := [2][2]int{}

	for i := 0; i < len(matrizA); i++ {
		for j := 0; j < len(matrizA[0]); j++ {
			for k := 0; k < len(matrizB); k++ {
				resultado[i][j] += matrizA[i][k] * matrizB[k][j]
			}
		}
	}

	return resultado
}

// problem 8
func burbujaBandera(array []int) []int {
	var aux int
	// para el conteo, no es necesario
	vueltas := 0
	// bandera boleana
	sinCambios := true

	for sinCambios {
		sinCambios = false
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				aux = array[j]
				array[j] = array[j+1]
				array[j+1] = aux
				vueltas += 1
				sinCambios = true
			}
		}
	}

	fmt.Printf("Cantidad de reordenamientos: %d \n", vueltas)
	return array
}
