package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
)

const halt = 1

var clear map[string]func()

func intro() {
	nome := os.Getenv("NOME")
	//fmt.Println("Digite seu nome:")
	//fmt.Scan(&nome)
	versao := 2.2
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
	fmt.Println("Digite seu comando: ")
	fmt.Scan(&comandoCarregado)
	fmt.Println("Comando Digitado ", comandoCarregado, "\n ")
	return comandoCarregado
}

func exibeOpcoes() {

	fmt.Println("****************************************")
	fmt.Println("Digite a opção que você deseja executar.")
	fmt.Println("1  - Script EPP")
	fmt.Println("2  - Script EPH")
	fmt.Println("3  - Script PRD")
	fmt.Println("4  - Script EPD")
	fmt.Println("5  - Senha SAP EPP")
	fmt.Println("6  - Senha SAP EPH")
	fmt.Println("7  - Senha SAP PRD")
	fmt.Println("8  - Senha SAP EPD")
	fmt.Println("9  - Script BMF")
	fmt.Println("10 - GSC")
	fmt.Println("0  - Sair do Programa")
	fmt.Println("****************************************\n ")
}

func ScriptEpp() {
	err := exec.Command("wscript", "loginEpp.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
	limpaTela()
}
func ScriptEph() {
	err := exec.Command("wscript", "loginEph.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
	limpaTela()
}
func ScriptPrd() {
	err := exec.Command("wscript", "loginPrd.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
	limpaTela()
}
func ScriptEpd() {
	err := exec.Command("wscript", "loginEpd.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
	limpaTela()
}
func epp() {
	epp := os.Getenv("EPP")
	fmt.Println("A senha atual do SAP EPP:", epp)
	clip(epp)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}
func eph() {
	eph := os.Getenv("EPH")
	fmt.Println("A senha atual do SAP EPP:", eph)
	clip(eph)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}
func prd() {
	prd := os.Getenv("PRD")
	fmt.Println("A senha atual do SAP PRD:", prd)
	clip(prd)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func epd() {
	epd := os.Getenv("EPD")
	fmt.Println("A senha atual do SAP EPD:", epd)
	clip(epd)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func ScriptBmf() {
	err := exec.Command("wscript", "loginBmf.vbs").Run()
	if err != nil {
		log.Fatal(err)
	}
	limpaTela()
}

func GSC() {
	gsc := os.Getenv("GSC")
	fmt.Println("A senha atual do SAP EPD:", gsc)
	clip(gsc)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()

}

/*func carregaComandoKill() string {
	var clear = make(map[string]func())
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

	var comandoKill string
	fmt.Println("Deseja Continuar? (S/N)")
	fmt.Scan(&comandoKill)
	fmt.Println("Comando Digitado:", comandoKill)
	return comandoKill
}

/*func coldWebPrd() {
	cold := ""
	fmt.Println("A senha atual do e-coldWeb PRD:", cold)
	clip(cold)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func emonitorPrd() {
	emonitor := ""
	fmt.Println("A senha atual do e-monitor PRD:", emonitor)
	clip(emonitor)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()
}

func ahgora() {
	senhaAhgora := ""
	fmt.Println("A senha atual do app Ahgora:", senhaAhgora)
	clip(senhaAhgora)
	msg()
	time.Sleep(time.Second * halt)
	limpaTela()

}*/

func limpaTela() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Sua platforma não suporta :(")
	}
}

func clip(pass string) {
	clipbd := pass
	err := clipboard.WriteAll(clipbd)
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	intro()
	for {
		exibeOpcoes()

		comando := carregaComando()
		switch comando {
		case 1:
			ScriptEpp()
		case 2:
			ScriptEph()
		case 3:
			ScriptPrd()
		case 4:
			ScriptEpd()
		case 5:
			epp()
		case 6:
			eph()
		case 7:
			prd()
		case 8:
			epd()
		case 9:
			ScriptBmf()
		case 10:
			GSC()
		case 0:
			fmt.Println("Saindo do Programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}
	}

}
