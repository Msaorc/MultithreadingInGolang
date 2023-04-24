package main

import "fmt"

type Conta struct {
	Numero int
	Nome   string
	Valor  float64
}

func main() {
	// ch := make(chan int)
	// go sendMessage(ch)
	// receiveMessage(ch)
	chConta := make(chan Conta)
	go createContasRecursive(chConta)
	confirmaContasCriadas(chConta)
}

func sendMessage(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveMessage(ch chan int) {
	for x := range ch {
		fmt.Printf("Message number %d received with sucess.\n", x)
	}
}

func newConta(conta int, nome string, valor float64) Conta {
	return Conta{Numero: conta, Nome: nome, Valor: valor}
}

func createContasRecursive(contaCanal chan Conta) {
	for i := 1; i <= 10; i++ {
		contaCanal <- newConta(i, fmt.Sprintf("Nome-%d", i), 250)
	}
	close(contaCanal)
}

func confirmaContasCriadas(contas chan Conta) {
	for conta := range contas {
		fmt.Printf("Conta: %d do cliente %s, com o valor %.2f foi criada com sucesso\n", conta.Numero, conta.Nome, conta.Valor)
	}
}
