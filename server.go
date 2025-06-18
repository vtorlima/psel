package main

/*
Para compilar:
	go build arquivo.go
	./arquivo
Para rodar:
	go run arquivo.go
*/

import (
	"fmt" //prints formatados
	"log" //mensagens em log
	"net" //conexoes TCP
)

func main() {
	fmt.Println("Servidor iniciado")

	ln, err := net.Listen("tcp", ":8080")

	/*
		essa linha é a forma reduzida de:

		var ln net.Listener
		var err error
		ln, err = net.Listen("tcp",":8080")

		:= serve pra declarar e inicializar uma variavel ao mesmo tempo sem
		precisar declarar o tipo

		cria um listener (ln) TCP que escuta na porta 8080
		o programa fica aguardando conexoes nessa porta
	*/

	if err != nil {
		log.Fatal(err)
	}

	//se o erro nao for nulo, retorna o erro e termina o programa

	defer ln.Close()

	/* defer: adia a execucao ate o final da funcao

	func exemplo() {
	fmt.Println("1")
	defer fmt.Println("3")
	fmt.Println("2")
	}

	sem essa linha, a porta nao vai fechar e vai permanecer indisponivel pra uso futuro
	*/

	for {
		conn, err := ln.Accept()

		/* blocking call - o método accept funciona de modo que o programa pausa nessa
		linha e só continua rodando o código até que algum usuario se conecte

		após o usuario se conectar, a funcao retorna uma variavel conn que contém o canal
		de comunicacao bidirecional. A partir dele que podemos ler os dados recebidos,
		escrever respostas e fechar a conexao.
		*/

		if err != nil {
			log.Println(err)
			continue
			/* se tiver erro ao aceitar a conexao, printa o erro e o continue faz com que
			a execucao volte para o inicio do for diretamente, ignorando tudo abaixo
			*/
		}
		go handleConnection(conn)

		/*goroutine - comeca a rodar em paralelo as acoes determinadas a respeito do
		usuario que conectou enquanto simultaneamente volta ao inicio do for e aguarda
		novas conexoes
		*/
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	/* se nao fechar os dados vao permanecer na memoria sem motivo
	além de existir um limite de conexoes simultaneas
	*/

	buf := make([]byte, 4096)
	//cria um buffer pra armazenar dados recebidos

	n, err := conn.Read(buf)

	/* blocking call - le os dados da conexao e armazena os resultados dentro do buffer "buf"
	n = numero de bytes lidos
	*/

	if err != nil {
		log.Println(err)
		return
		//se der erro na leitura, registra e encerra
	}

	requestText := string(buf[:n])
	// transforma os bytes lidos em uma string (pegando apenas n, nao todos)

	fmt.Println(requestText)
	// printa o conteudo da requisicao http recebida
}
