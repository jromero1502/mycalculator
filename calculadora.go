package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		fmt.Println("Operador no soportado")
		return 0
	}
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
