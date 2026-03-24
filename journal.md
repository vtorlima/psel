# Dia 1

Comecei estudando o que são load balancers e quais conceitos são necessários antes de chegar nessa parte. Para isso, revisei sockets, HTTP e o básico de TCP, focando principalmente na conexão do cliente, aceitação pelo servidor, fluxo de bytes e encerramento da conexão.

Como eu já tinha tido contato prévio com esses temas, não houve muita novidade, mas aproveitei para criar alguns flashcards a fim de reforçar a fixação de pontos importantes.

Além disso, pesquisei algumas funções da linguagem Go, especialmente `net.Listen`, `net.Conn`, `conn.Read` e `conn.Write`, analisando exemplos práticos de como elas funcionam.

Por fim, criei o fork do repositório e organizei a estrutura inicial do projeto, adicionando pastas separadas para o load balancer e para o backend.

# Dia 2

Hoje comecei relembrando `net.Listen` e goroutines (sem utilizá-las no `handleConn`, o servidor passa a tratar apenas uma conexão por vez).

Iniciei a estrutura do servidor a partir de um TCP echo server, ainda sem tratamento de erros e com uma implementação simplificada de `handleConn`, mas já consegui testá-lo localmente. Também comentei os pontos mais importantes do código para não esquecer detalhes nos próximos dias.