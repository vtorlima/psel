package main //define o arquivo como executavel

/*
Para compilar:
	go build arquivo.go
	./arquivo
Para rodar:
	go run arquivo.go
*/

import (
	"fmt" //prints formatados
	"net" //conexoes TCP
	//    "strings" //manipular texto
	"log" //mensagens em log
)

func main() {
	//funcao principal do programa
	fmt.Println("hello, world!")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		// blocking call - programa espera nessa linha ate que alguem se conecte
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
		// Goroutine - comeca a rodar em paralelo as acoes determinadas a respeito
		//do usuario que conectou
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Fecha no final

	buf := make([]byte, 4096)
	// cria um buffer
	n, err := conn.Read(buf)
	// blocking call - le os dados da conexao e armazena os resultados dentro do buffer "buf"
	// n = numero de bytes lidos
	if err != nil {
		log.Println(err)
		return
	}

	requestText := string(buf[:n])
	// transforma os bites lidos em uma string (pegando apenas n, nao todos)
	fmt.Println(requestText)

	/*
		exemplo de um HTTP request

		GET /meuarquivo.txt HTTP/1.1
		Host: localhost:8080
		User-Agent: Mozilla/5.0...
		Accept: text/html,...
	*/

	/* TEMPORARELY DISABLED
	lines := strings.Split(requestText, "\r\n")
	// separa as linhas do request (\r\n troca a linha)
	firstLine := lines[0]
	parts := strings.Split(firstLine, " ")
	// separa os elementos + importantes do request
	method := parts[0]
	path := parts[1]
	*/

	// exemplo de resposta manual

	body := "<html><body>hello!</body></html>"
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
		"\r\n" +
		body

	conn.Write([]byte(response))

	/*resultado:

	HTTP/1.1 200 OK
	Content-Type: text/html
	Content-Length: 35

	<html><body>Hello!</body></html>
	*/

	// leitura, parsing, resposta aqui
}
