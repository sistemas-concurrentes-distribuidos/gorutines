package main

import (
	"fmt"
	"gorutines/models"
)

func main() {

	var contador uint64
	var procesos []models.Proceso
	op := -1

	for op != 0 {
		fmt.Println("-----------------------------------")
		fmt.Printf("ADMINISTRADOR DE PROCESOS\n\n")
		fmt.Println("Procesos activos: ", len(procesos))
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Eliminar proceso")
		fmt.Println("3) Mostrar procesos")
		fmt.Println("0) SALIR")
		fmt.Print("Selecciona una opción: ")
		fmt.Scanln(&op)

		fmt.Println("OP: ", op)

		switch op {
		case 1:
			p := new(models.Proceso)
			p.Init(contador)
			procesos = append(procesos, *p)
			contador++
		case 2:
			fmt.Println("Elimnar proceso")
		case 3:
			for _, p := range procesos {
				go p.Printer()
			}
		case 0:
			return
		default:
			fmt.Println("Deteniendo procesos...")
			for _, p := range procesos {
				fmt.Println(p.Id)
				p.Stop()
			}
			procesos = nil
		}
		op = -1
	}
}
