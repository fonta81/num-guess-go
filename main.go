package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ValorAleatorio := CrearNumero()
	for {
		ValorUsuario := PedirValor()
		resultado := Revisar(ValorAleatorio, ValorUsuario)
		fmt.Println(resultado)
		if resultado == "!!Ganaste!!" {
			break
		}
	}
}

// CrearNumero -> crea un nuero del 1-10 y lo devuelve como Int
func CrearNumero() int {
	return rand.IntN(11)
}

// PedirValor regresa un int del valor que escribio el usuario
// en caso de errores imprime el erro y regresa 0
func PedirValor() int {
	var valor int
	fmt.Print("Ingresa un valor: ")
	_, err := fmt.Scan(&valor)
	if err != nil {
		fmt.Println("Error: \n", err)
		return 0
	}
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
