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

# Dia 4

Hoje implementei as rotas do file server para leitura e escrita de arquivos utilizando HTTP.

Primeiro, modifiquei o `handleConn`, removendo o comportamento anterior de echo server e passando a tratar a requisição de forma estruturada usando o `parseRequest`.

A partir disso, adicionei lógica para diferenciar requisições `GET` e `POST` no formato `/files/{name}`, utilizando o `Method` e o `Path` extraídos da requisição.

No caso de `GET`, implementei a leitura do arquivo no disco e retorno do conteúdo na resposta. Já no `POST`, utilizei o `Body` da requisição para escrever no arquivo correspondente.

Também adicionei validações básicas, como verificação de path inválido, nome de arquivo vazio e métodos não suportados, retornando os códigos HTTP utilizados no servidor (`200`, `400`, `404`, `405` e `500`).

Além disso, ajustei o comportamento do servidor para garantir que a pasta `files` seja criada apenas quando necessário (no `POST`), mantendo o `GET` como operação apenas de leitura.

Por fim, testei o funcionamento com `curl`, validando as requisições ao servidor e o tratamento de rotas de leitura e escrita de arquivos.

Além disso, estudei permissões em octal e códigos HTTP, organizando essas informações em tabelas para facilitar a consulta e evitar confusão.
---

## Códigos HTTP usados no projeto

| Código | Nome | Quando usar |
|--------|------|------------|
| 200 | OK | Requisição deu certo (GET ou POST funcionou) |
| 400 | Bad Request | Requisição mal formada (ex: `/files/` vazio) |
| 404 | Not Found | Arquivo não existe ou rota inválida |
| 405 | Method Not Allowed | Método não suportado |
| 500 | Internal Server Error | Erro inesperado no servidor |

---

## Outros códigos HTTP importantes

| Código | Nome | Quando usar |
|--------|------|------------|
| 201 | Created | Recurso criado com sucesso (POST criando arquivo) |
| 204 | No Content | Sucesso sem conteúdo de resposta |
| 301 | Moved Permanently | Redirecionamento permanente |
| 302 | Found | Redirecionamento temporário |
| 401 | Unauthorized | Precisa de autenticação |
| 403 | Forbidden | Acesso negado |
| 409 | Conflict | Conflito (ex: arquivo já existe) |
| 413 | Payload Too Large | Body muito grande |
| 415 | Unsupported Media Type | Tipo não suportado |
| 429 | Too Many Requests | Rate limit |

---

## Permissões em octal (chmod)

### Valores individuais

| Valor | Permissão | Significado |
|------|----------|------------|
| 0 | --- | nenhuma permissão |
| 1 | --x | executar |
| 2 | -w- | escrever |
| 3 | -wx | escrever + executar |
| 4 | r-- | ler |
| 5 | r-x | ler + executar |
| 6 | rw- | ler + escrever |
| 7 | rwx | ler + escrever + executar |

---

### Estrutura

| Parte | Significado |
|------|------------|
| 0 | base octal |
| X | dono |
| Y | grupo |
| Z | outros |

---

### Combinações mais usadas

| Código | Dono | Grupo | Outros | Significado |
|--------|------|-------|--------|------------|
| 644 | rw- | r-- | r-- | dono lê/escreve, outros só leem |
| 600 | rw- | --- | --- | só o dono acessa |
| 666 | rw- | rw- | rw- | todos leem e escrevem |
| 755 | rwx | r-x | r-x | dono tudo, outros leem/executam |
| 700 | rwx | --- | --- | só o dono pode tudo |
