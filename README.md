# PATOS/POMBO PSEL 2.0 - LOAD BALANCER

---

## Como rodar

### Pré-requisitos

Antes de tudo, você precisa ter instalado:

- Go 1.21 ou superior — https://go.dev/dl/

Verifique com:

```bash
go version
```

### Clonar o repositório

```bash
git clone https://github.com/<seu-usuario>/psel.git
cd psel
```

### Rodar o projeto

**Opção 1 — Usando `make` (se tiver instalado)**

```bash
make run-all
```

**Opção 2 — Sem `make` (rodar manualmente)**

Se você não tiver `make`, rode os comandos abaixo diretamente no terminal:

```bash
PORT=9001 go run ./fileserver/ &
PORT=9002 go run ./fileserver/ &
PORT=9003 go run ./fileserver/ &
go run ./loadbalancer/
```

Isso irá subir:

- 3 file servers (portas 9001, 9002, 9003)
- 1 load balancer (porta 9090)

---

## Como funciona

O load balancer roda em:

```
http://localhost:9090
```

Ele distribui as requisições entre três servidores backend:

| Backend | Porta |
|---------|-------|
| localhost:9001 | 9001 |
| localhost:9002 | 9002 |
| localhost:9003 | 9003 |

---

## Testes

### Enviar um arquivo (POST)

```bash
curl -X POST http://localhost:9090/files/[nomeDoArquivo] --data "[conteúdo]"
```

### Baixar um arquivo (GET)

```bash
curl http://localhost:9090/files/[nomeDoArquivo]
```

### Navegador

Abra no browser:

```
http://localhost:9090/files/[nomeDoArquivo]
```

### Dica

Experimente com esses arquivos:

patos.txt
pombo.txt
