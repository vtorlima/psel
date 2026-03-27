package main

import (
	"io"  // pra usar o copy e repassar os dados
	"net" // conexões TCP
)

func main() {
	// cria um listener na porta 9090
	ln, _ := net.Listen("tcp", ":9090")

	// loop infinito pra continuar aceitando clientes
	for {
		// aceita uma conexão
		clientConn, _ := ln.Accept()

		// cada cliente numa goroutine
		go forward(clientConn)
	}
}

// encaminha a requisição do cliente para o backend
func forward(clientConn net.Conn) {
	// garante que a conexão com o cliente feche ao final
	defer clientConn.Close()

	// buffer pra armazenar a requisição recebida
	buf := make([]byte, 4096)

	// lê a requisição recebida
	n, err := clientConn.Read(buf)
	if err != nil {
		return
	}

	// conecta no backend
	backendConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}

	// garante que a conexão com o backend feche no final
	defer backendConn.Close()

	// envia para o backend os bytes recebidos do cliente
	_, err = backendConn.Write(buf[:n])
	if err != nil {
		return
	}

	// copia a resposta do backend de volta para o cliente
	io.Copy(clientConn, backendConn)
}
