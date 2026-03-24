package main

import (
	"net" //pacote para trabalhar com rede (TCP, sockets, etc)
)

func main() {

	//sem tratamento de erro por enquanto
	//cria um servidor TCP escutando na porta 8080
	//retorna um listener (ln) que é responsável por aceitar conexões
	ln, _ := net.Listen("tcp", ":8080")

	//loop infinito: servidor nunca para de aceitar novas conexões enquanto ativo
	for {
		//espera aqui até que um cliente se conecte
		//quando alguém conecta, o Accept() pega essa conexão
		//e retorna conn (socket específico dessa conexão)
		conn, _ := ln.Accept()

		//cria uma goroutine pra lidar com o cliente
		//permite atender múltiplos clientes ao mesmo tempo
		go handleConn(conn)

	}
}

// função responsável por lidar com uma conexão específica
func handleConn(conn net.Conn) {

	//garante que a conexão feche quando a função terminar
	defer conn.Close()

	//cria um buffer de 1024 bytes
	//esse buffer vai armazenar os dados recebidos temporariamente
	buf := make([]byte, 1024)

	// loop infinito pra continuar lendo os dados dessa conexão
	for {

		//le dados vindos do cliente e coloca dentro do buffer
		//n é o número de bytes lidos
		n, _ := conn.Read(buf)

		println(string(buf[:n]))

		//retorna os mesmos bytes que foram recebidos
		//buf[:n] garante que só a parte válida seja enviada
		conn.Write(buf[:n])
	}
}
