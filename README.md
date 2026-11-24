
# 🔑 PassSanta

Aplicação desenvolvida em Go para facilitar o gerenciamento e a recuperação rápida de senhas utilizadas no dia a dia operacional usuário SAP.

O objetivo da ferramenta é permitir que o usuário selecione uma senha pré-configurada, não precise ficar digitando usuário e senha no SAP gui, e caso precise copie automaticamente a senha para a área de transferência e a utilize em processos de autenticação sem a necessidade de digitação manual.

## ✨ Funcionalidades

- 🔑 Exibição de senhas pré-configuradas.
- 📋 Cópia automática para a área de transferência.
- 💻 Interface simples via terminal.
- ⚡ Execução rápida e leve.
- 🐹 Desenvolvido em Golang.

## 🛠️ Requisitos

- Go 1.20 ou superior
- Windows 10 ou superior

## 🚀 Instalação

Clone o repositório:

```bash
git clone https://github.com/jclaudiotomasjr/passSanta.git
```

Acesse a pasta do projeto:

```bash
cd passSanta
```

Instale as dependências:

```bash
go mod tidy
```

Execute a aplicação:

```bash
go run .
```

## 📦 Build

Para gerar o executável:

```bash
go build -o passSanta.exe
```

O executável será criado na pasta atual.

## 🔒 Segurança

Por questões de segurança, recomenda-se:

- 🚫 Não armazenar senhas diretamente no código-fonte.
- ✅ Utilizar variáveis de ambiente.
- ✅ Utilizar arquivos `.env` não versionados.
- ✅ Adicionar o arquivo `.env` ao `.gitignore`.
- ✅ Nunca publicar credenciais em repositórios GitHub.

### 📝 Exemplo de .env

```env
SAP_EPD=senha_epd
SAP_EPH=senha_qas
SAP_PRD=senha_prd
USER=USUARIO01
```

## 📂 Estrutura do Projeto

```text
passSanta/
│
├── main.go
├── senhas.go
├── go.mod
├── go.sum
└── README.md
```

## 🧰 Tecnologias Utilizadas

- 🐹 Go (Golang)
- 📋 Clipboard Package
- 🖥️ Windows Console

## 👨‍💻 Autor

José Claudio Tomas Junior

- 🔗 GitHub: https://github.com/jclaudiotomasjr
- 🔗 LinkedIn: www.linkedin.com/in/claudiojtomassap

## 📄 Licença

Este projeto é destinado para uso interno e fins de automação operacional no uso do SAP GUI.

---

⭐ Se este projeto foi útil para você, considere deixar uma estrela no repositório!