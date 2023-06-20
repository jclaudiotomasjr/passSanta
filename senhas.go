package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/atotto/clipboard"
)

const halt = 2

var clear map[string]func()

func intro() {

	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("****************************************")
	fmt.Println("*Bem-vindo!", nome)
	fmt.Println("*Estamos na versão do programa", versao)

}

func carregaComando() int {

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	var comandoCarregado int
	fmt.Print("Digite seu comando: ")
	fmt.Scan(&comandoCarregado)
	fmt.Println("Comando Digitado ", comandoCarregado, "\n ")
	return comandoCarregado
}

func exibeOpcoes() {

	fmt.Println("****************************************")
	fmt.Println("Digite a opção que você deseja executar.")
	fmt.Println("1 - Senha SAP EPP")
	fmt.Println("2 - Senha SAP PRD")
	fmt.Println("3 - Senha SAP EPD")
	fmt.Println("4 - Senha ColdWeb PRD")
	fmt.Println("5 - Senha Emonitor PRD")
	fmt.Println("6 - Senha Ahgora")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("****************************************\n ")
}

func epp() {
	epp := "1234"
	fmt.Println("A senha atual do SAP EPP:", epp)
	clip(epp)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}
func prd() {
	prd := "1234"
	fmt.Println("A senha atual do SAP PRD:", prd)
	clip(prd)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func epd() {
	epd := "1234"
	fmt.Println("A senha atual do SAP EPD:", epd)
	clip(epd)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}
func coldWebPrd() {
	cold := "1234"
	fmt.Println("A senha atual do e-coldWeb PRD:", cold)
	clip(cold)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func emonitorPrd() {
	emonitor := "1234"
	fmt.Println("A senha atual do e-monitor PRD:", emonitor)
	clip(emonitor)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func ahgora() {
	senhaAhgora := "1234"
	fmt.Println("A senha atual do app Ahgora:", senhaAhgora)
	clip(senhaAhgora)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()

}

func limpaTela() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Sua platforma não suporta :(")
	}
}

func clip(text string) {
	clpBd := text
	err := clipboard.WriteAll(clpBd)
	if err != nil {
		fmt.Println("Erro ao copiar senha", err)
		return
	}

}

func msg() {
	fmt.Println("Senha copiada com Sucesso!")
}

func main() {

	color1 := exec.Command("cmd", "/c", "color 2")
	color1.Stdout = os.Stdout
	color1.Run()

	intro()
	for {
		exibeOpcoes()

		comando := carregaComando()
		switch comando {
		case 1:
			epp()
		case 2:
			prd()
		case 3:
			epd()
		case 4:
			coldWebPrd()
		case 5:
			emonitorPrd()
		case 6:
			ahgora()
		case 0:
			fmt.Println("Saindo do Programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}
	}

}
