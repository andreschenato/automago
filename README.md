# Automago - Analisador Léxico
O Automago é um app web extremamente simples construído para o estudo de autômatos finitos em linguagens formais. 

Uma demo do site pode ser acessada em https://automago.andreschenato.dev.br.

Ele funciona em duas partes:
1. A construção da linguagem com palavras inputadas pelo usuário;
2. O input de tokens para validar se estes são ou não aceitos pela linguagem.

## Como rodar o projeto
Existem duas opções para rodar este projeto localmente, a primeira é ter o Go/Golang instalado na máquina, você pode ver o tutorial de instalação [aqui](https://go.dev/doc/install). Outra opção é usar Docker, que não requer o Go instalado.

### Rodando na máquina
Para rodar direto na máquina, primeiro instale o Go, em seguida faça um clone deste repositório e rode os comandos abaixo:
```
go mod tidy
go run .
```

O `tidy` é, em tese, desnecessário neste projeto, já que ele não usa nenhuma lib de terceiros, entretanto, é sempre bom garantir.

Após rodar estes comandos é esperado que você veja uma mensagem no terminal com o endereço do app, neste caso: http://localhost:8080.

### Rodando com Docker
Usando o Docker, como dito anteriormente, não é necessário ter o Go instalado na máquina, basta rodar os seguintes comandos:
```
docker build -t automago .
docker run -d --name automago -p 8080:8080 automago
```

Com isso, você deve ter um container chamado `automago` e o projeto estará acessível na porta 8080, ou simplesmente http://localhost:8080.
