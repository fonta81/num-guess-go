package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// ValorAleatorio := CrearNumero()
}

// CrearNumero -> crea un nuero del 1-10 y lo devuelve como Int
func CrearNumero() int {
	return rand.IntN(11)
}

func PedirValor() int {
	var valor int
	_, _ := fmt.Scan(&valor)

	return valor
}

// Revisar Compara y devuelve que tan cerca esta del numero generando con el del usuario
// y devuelve un mensaje dependiendo de que tan cerca este
func Revisar(Numerogenerado int, NumeroUsuario int) string {
	if Numerogenerado > NumeroUsuario {
		if Numerogenerado-NumeroUsuario <= 2 {
			return "Cerca,Un poco mas alto"
		} else {
			return "Mas alto"
		}
	} else if Numerogenerado < NumeroUsuario {
		if NumeroUsuario-Numerogenerado <= 2 {
			return "Cerca, un poco mas abajo"
		} else {
			return "Mas abajo"
		}
	} else {
		return "!!Ganaste!!"
	}
}
