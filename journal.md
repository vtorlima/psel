# Dia 1

Comecei estudando o que são load balancers e quais conceitos são necessários antes de chegar nessa parte. Para isso, revisei sockets, HTTP e o básico de TCP, focando principalmente na conexão do cliente, aceitação pelo servidor, fluxo de bytes e encerramento da conexão.

Como eu já tinha tido contato prévio com esses temas, não houve muita novidade, mas aproveitei para criar alguns flashcards a fim de reforçar a fixação de pontos importantes.

Além disso, pesquisei algumas funções da linguagem Go, especialmente `net.Listen`, `net.Conn`, `conn.Read` e `conn.Write`, analisando exemplos práticos de como elas funcionam.

Por fim, criei o fork do repositório e organizei a estrutura inicial do projeto, adicionando pastas separadas para o load balancer e para o backend.

# Dia 2

Hoje comecei relembrando `net.Listen` e goroutines (sem utilizá-las no `handleConn`, o servidor passa a tratar apenas uma conexão por vez).

Iniciei a estrutura do servidor a partir de um TCP echo server, ainda sem tratamento de erros e com uma implementação simplificada de `handleConn`, mas já consegui testá-lo localmente. Também comentei os pontos mais importantes do código para não esquecer detalhes nos próximos dias.

# Dia 3

Hoje comecei estudando a estrutura de uma requisição HTTP, focando principalmente em como identificar o fim dos headers usando `\r\n\r\n` e na composição da primeira linha (method, path e version).

Implementei o `parser.go`, criando a struct `HTTPRequest` e a função `parseRequest`, que separa a requisição em headers e body e extrai as principais informações.

Depois, integrei o parser ao servidor no `handleConn`, passando a interpretar a requisição recebida e imprimir o método e o path no terminal.

Também adicionei uma resposta HTTP mínima (`HTTP/1.1 200 OK\r\n\r\n`) para evitar que o browser ficasse esperando para sempre.

Por fim, testei o funcionamento com `curl`, validando tanto requisições GET quanto POST e confirmando que o parser está extraindo corretamente as informações da requisição.