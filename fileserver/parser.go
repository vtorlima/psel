package main

import (
	"strings" // utilizado para manipular as requisições
)

// struct para guardar a requisição depois de separar
// em vez de trabalhar com uma string imensa o tempo todo
type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

// recebe a requisição HTTP inteira bruta e devolve organizada

func parseRequest(raw string) HTTPRequest {
	// primeiro separamos a requisição em 2 blocos:
	// tudo antes de rnrn -> request line + headers
	// tudo depois -> body
	// o 2 limita o separador para no máximo 2 partes, caso o body tenha rnrn
	parts := strings.SplitN(raw, "\r\n\r\n", 2)

	// aqui pegamos a primeira parte (request line + headers)
	headerPart := parts[0]

	// como nem toda requisição tem body, começamos assumindo que está vazio
	bodyPart := ""

	// se existir uma segunda parte depois do rnrn, é o body
	if len(parts) > 1 {
		bodyPart = parts[1]
	}

	// agora quebramos a parte de cima por linhas
	// GET / HTTP/1.1\r\nHost: localhost:8080\r\nAccept: */*\r\n
	// vira:
	// lines[0] = "GET "
	// lines[1] = "Host: localhost:8080"
	// lines[2] = "Accept: */*"
	lines := strings.Split(headerPart, "\r\n")

	// a primeira linha é sempre a request line
	requestLine := lines[0]

	// a request line tem o formato: METHOD PATH VERSION
	requestLineParts := strings.Split(requestLine, " ")
	//  se por algum motivo a linha não tiver o formato padrão, ignora
	if len(requestLineParts) != 3 {
		return HTTPRequest{}
	}
	method := requestLineParts[0]
	path := requestLineParts[1]
	version := requestLineParts[2]

	// criação do mapa pra armazenar os headers
	// // sem o make, o map fica nil (não inicializado), em vez de vazio
	headers := make(map[string]string)

	// percorrer o resto das linhas, lines[0] já foi usada como request line
	for i := 1; i < len(lines); i++ {
		line := lines[i]

		// se a linha estiver vazia, ignora
		if line == "" {
			continue
		}

		// cada header tem formato key: value
		// split pra separar os dois
		headerParts := strings.SplitN(line, ": ", 2)

		// também checa o formato, se não for padrão, ignora
		if len(headerParts) < 2 {
			continue
		}
		// primeira parte é a chave
		key := headerParts[0]
		// segunda parte é o valor
		value := headerParts[1]
		// armazena no mapa
		headers[key] = value
	}

	return HTTPRequest{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Body:    bodyPart,
	}

}
