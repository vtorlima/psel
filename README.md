# PSEL PATOS 2.0 - Load Balancer

Load balancer HTTP em Go, escrito do zero para o PSEL PATOS/POMBO 2.0.

## Estrutura do Projeto

- **fileserver/** : servidor HTTP que armazena e serve arquivos via GET/POST em `/files/{nome}`
- **loadbalancer/** : proxy TCP na porta 9090 que distribui requisições entre fileservers usando least-connections

> A documentação do processo de aprendizado se encontra em [journal.md](journal.md)

## Pré-requisitos

- Go 1.21+
- `make` (opcional)

## Como Rodar

A aplicação é composta por 4 processos que ficam rodando simultaneamente:
- 3 fileservers (nas portas 9001, 9002, 9003)
- 1 load balancer (na porta 9090)

Cada um precisa do seu próprio terminal, porque cada processo bloqueia o terminal enquanto está ativo.

**Você vai precisar de:**
- 4 terminais para os processos em execução
- (+1 terminal se quiser testar com curl, também dá pra testar pelo navegador)

Todos os terminais devem estar abertos na raiz do projeto (`psel/`).

### Com Make

Abra 4 terminais na raiz do projeto. Em cada um, rode um comando:

**Terminal 1:**
```bash
make run-backend1
```

**Terminal 2:**
```bash
make run-backend2
```

**Terminal 3:**
```bash
make run-backend3
```

**Terminal 4:**
```bash
make run-lb
```

Os 4 processos agora estão rodando. Para parar qualquer um, aperte `Ctrl+C` no terminal correspondente.

### Sem Make

Você precisa primeiro compilar os binários (uma vez só) e depois rodar os processos (4 terminais simultâneos).

**Passo 1 — Compilar**

Em qualquer terminal, na raiz do projeto:

**Linux/macOS:**
```bash
mkdir -p bin
go build -o bin/fileserver ./fileserver
go build -o bin/loadbalancer ./loadbalancer
```

**Windows (PowerShell):**
```powershell
mkdir bin -Force
go build -o bin/fileserver.exe ./fileserver
go build -o bin/loadbalancer.exe ./loadbalancer
```

**Passo 2 — Rodar**

Abra 4 terminais e em cada um execute um dos comandos abaixo (um terminal por linha).

**Linux/macOS:**
```bash
# Terminal 1
PORT=9001 ./bin/fileserver

# Terminal 2
PORT=9002 ./bin/fileserver

# Terminal 3
PORT=9003 ./bin/fileserver

# Terminal 4
./bin/loadbalancer
```

**Windows (PowerShell):**
```powershell
# Terminal 1
$env:PORT="9001"; .\bin\fileserver.exe

# Terminal 2
$env:PORT="9002"; .\bin\fileserver.exe

# Terminal 3
$env:PORT="9003"; .\bin\fileserver.exe

# Terminal 4
.\bin\loadbalancer.exe
```

Cada terminal vai travar mostrando os logs do seu processo. Para parar, `Ctrl+C` em cada terminal.

## Endpoints

Todas as requisições devem ir para o load balancer (`localhost:9090`):

| Método | Rota | Efeito |
|--------|------|--------|
| GET | `/files/{nome}` | Retorna o conteúdo do arquivo |
| POST | `/files/{nome}` | Grava o body como arquivo |

## Como Fazer Requisições

Com os 4 processos rodando, abra um terminal extra para mandar requisições. Todas as requisições vão para o load balancer na porta 9090. 

Atenção (Windows/PowerShell): Para usar o curl no PowerShell, escreva curl.exe (com extensão) em todos os exemplos abaixo. No cmd, Git Bash, WSL, Linux ou macOS, basta curl.

### 1. Gravar um arquivo (POST)

O `-X POST` define o método e o `-d` define o conteúdo (body) que vai ser gravado como arquivo:

```bash
curl -X POST -d "conteúdo do arquivo" http://localhost:9090/files/teste.txt
```

Para enviar um arquivo do disco em vez de digitar o conteúdo, troque `-d "..."` por `--data-binary @nome-do-arquivo`. O `@` indica que é um caminho de arquivo:

```bash
curl -X POST --data-binary @local.txt http://localhost:9090/files/teste.txt
```

### 2. Ler um arquivo (GET)

```bash
curl http://localhost:9090/files/teste.txt
```

### 3. Ler pelo navegador

Abra qualquer navegador e cole na barra de endereço:

```
http://localhost:9090/files/teste.txt
```

DICA: tente visualizar os arquivos pombo.txt e patos.txt !!!
