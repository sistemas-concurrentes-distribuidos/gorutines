package main

import (
	"fmt"
	"time"
)

var active bool = false
var deletedRutine uint64 = 0
var contador uint64 = 0

func Proceso(id uint64) {
	i := uint64(0)
	for {
		if active {
			fmt.Printf("id %d: %d \n", id, i)
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
		if deletedRutine == id {
			deletedRutine = 0
			break
		}
	}
}

func main() {
	var op int
	for {
		fmt.Println("-----------------------------------")
		fmt.Printf("ADMINISTRADOR DE PROCESOS\n\n")
		fmt.Println("Procesos activos: ", contador)
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar proceso")
		fmt.Println("3) Eliminar procesos")
		fmt.Println("0) SALIR")
		fmt.Scanln(&op)

		switch op {
		case 1:
			contador++
			go Proceso(contador)
		case 2:
			if contador > 0 {
				active = true
				fmt.Scanln()
				active = false
			} else {
				fmt.Printf("\nNo hay procesos registrados\n")
			}
		case 3:
			if contador > 0 {
				var error bool = true
				for error {
					fmt.Print("ID del proceso a leiminar: ")
					fmt.Scanln(&deletedRutine)
					if deletedRutine > contador || deletedRutine == 0 {
						fmt.Printf("\nEl proceso ingresado no existe\n")
						error = true
					} else {
						error = false
						contador--
					}
				}
			} else {
				fmt.Printf("\nNo hay procesos registrados\n")
			}

		case 0:
			return
		}
	}
}
