package models

import (
	"fmt"
	"time"
)

type Proceso struct {
	Id      uint64
	CData   chan int
	CStatus chan bool
}

func (p *Proceso) Init(id uint64) {
	p.Id = id
	p.CData = make(chan int)
	p.CStatus = make(chan bool)
	go p.start()
}

func (p *Proceso) start() {
	var cont int
	cont = 0
	for {
		p.CData <- cont
		cont++
	}
}

func (p *Proceso) Printer() {

	for {
		select {
		case <-p.CStatus:
			return
		default:
			cont := <-p.CData
			fmt.Printf("id %d: %d \n", p.Id, cont)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// func (p Proceso) Start(id uint64) {
// 	fmt.Println("Proceso iniciado")
// 	i := uint64(0)
// 	for {
// 		fmt.Printf("id %d: %d \n", id, i)
// 		i = i + 1
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func (p *Proceso) Stop() {
	p.CStatus <- true
	fmt.Println("Proceso detenido: ", p.Id)
}
