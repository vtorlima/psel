package main

import (
	"fmt"
	"net" //pacote para trabalhar com rede (TCP, sockets, etc)
	"os"  // ler e escrever arquivos
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

	//aumenta o buffer de 1024 pra 4096
	//esse buffer vai armazenar os dados recebidos temporariamente
	buf := make([]byte, 4096)

	// lê os dados recebidos
	n, _ := conn.Read(buf)

	// converte os bytes recebidos em string
	raw := string(buf[:n])

	// chama o parser -> transforma a string bruta na struct organizada
	req := parseRequest(raw)

	// Imprime o método da requisição
	fmt.Println("Method:", req.Method)

	// Imprime o caminho solicitado
	fmt.Println("Path:", req.Path)

	// se o path não começar com /files/, responde 404 ( not found )
	if len(req.Path) < 7 || req.Path[:7] != "/files/" {
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		conn.Write([]byte(response))
		return
	}

	// pega apenas o nome do arquivo (removendo files)
	fileName := req.Path[7:]

	// se o nome do arquvo estiver vazio, responde 400 (bad request)
	if fileName == "" {
		response := "HTTP/1.1 400 Bad Request\r\n\r\n"
		conn.Write([]byte(response))
		return
	}

	// caminho do arquivo no disco
	filePath := "files/" + fileName

	// se for GET, le o arquivo e devolve o conteúdo
	if req.Method == "GET" {
		data, err := os.ReadFile(filePath)

		if err != nil {
			response := "HTTP/1.1 404 Not Found\r\n\r\n"
			conn.Write([]byte(response))
			return
		}

		response := "HTTP/1.1 200 OK\r\n\r\n" + string(data)
		conn.Write([]byte(response))
		return

	}
	// se for POST, escreve o body no arquivo
	if req.Method == "POST" {
		// garante que a pasta files exista
		os.MkdirAll("files", 0755)
		err := os.WriteFile(filePath, []byte(req.Body), 0644)

		// falha ao escrever o arquivo -> 500
		if err != nil {
			response := "HTTP/1.1 500 Internal Server Error\r\n\r\n"
			conn.Write([]byte(response))
			return
		}
		response := "HTTP/1.1 200 OK\r\n\r\n"
		conn.Write([]byte(response))
		return
	}
	// se não for GET nem POST, responde 405
	response := "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
	conn.Write([]byte(response))
}
