package main

import (
	"fmt"
	"net" //pacote para trabalhar com rede (TCP, sockets, etc)
	"os"  // ler e escrever arquivos
	"strconv"
)

func main() {

	// lê a porta pela variável PORT
	port := os.Getenv("PORT")

	// se nenhuma porta for informada, usa 8080
	if port == "" {
		port = "8080"
	}

	// cria um servidor tcp escutando na porta
	// returna um listener (ln) que é responsável por aceitar conexões
	ln, _ := net.Listen("tcp", ":"+port)

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

// monta uma resposta HTTP com os headers essenciais para compatibilidade com browser
func httpResponse(statusLine string, body string) string {
	headers := "Content-Length: " + strconv.Itoa(len(body)) + "\r\n"
	headers += "Content-Type: text/plain\r\n"
	headers += "Connection: close\r\n"
	return statusLine + "\r\n" + headers + "\r\n" + body
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
		conn.Write([]byte(httpResponse("HTTP/1.1 404 Not Found", "")))
		return
	}

	// pega apenas o nome do arquivo (removendo files)
	fileName := req.Path[7:]

	// se o nome do arquvo estiver vazio, responde 400 (bad request)
	if fileName == "" {
		conn.Write([]byte(httpResponse("HTTP/1.1 400 Bad Request", "")))
		return
	}

	// caminho do arquivo no disco
	filePath := "files/" + fileName

	// se for GET, le o arquivo e devolve o conteúdo
	if req.Method == "GET" {
		data, err := os.ReadFile(filePath)

		if err != nil {
			conn.Write([]byte(httpResponse("HTTP/1.1 404 Not Found", "")))
			return
		}

		conn.Write([]byte(httpResponse("HTTP/1.1 200 OK", string(data))))
		return

	}
	// se for POST, escreve o body no arquivo
	if req.Method == "POST" {
		// garante que a pasta files exista
		os.MkdirAll("files", 0755)
		err := os.WriteFile(filePath, []byte(req.Body), 0644)

		// falha ao escrever o arquivo -> 500
		if err != nil {
			conn.Write([]byte(httpResponse("HTTP/1.1 500 Internal Server Error", "")))
			return
		}
		conn.Write([]byte(httpResponse("HTTP/1.1 200 OK", "")))
		return
	}
	// se não for GET nem POST, responde 405
	conn.Write([]byte(httpResponse("HTTP/1.1 405 Method Not Allowed", "")))
}
