# Journal

## Dia 1

Honestamente, não tenho nenhuma experiência com redes de computadores.  
Conheci o PATOS pela internet e me despertou o interesse de aplicar, em parte por conta das participações que vi da equipe em hackathons, em parte pela bagagem que os integrantes têm de experiência no geral, e em parte pelo desafio de aprender algo novo, já que estou com bastante tempo livre neste primeiro semestre na UFSCar.

Por isso, reservei hoje para estudar o tema e entender de fato como isso tudo funciona.  
Esses dois links em especial foram bem úteis:

- [Aulão do protocolo HTTP + Flash Talks - Patos](https://www.youtube.com/watch?v=iuwSYRdxKjQ)
- [Beej's Guide to Network Programming](https://beej.us/guide/bgnet/)

Dei uma olhada em alguns servidores básicos em C e terminei o dia sem escrever nenhuma linha de código — fiquei apenas dissecando o guia.

## Dia 2

Diferentemente de ontem, resolvi aceitar a dica e utilizar Go.  
Abusei bastante da ajuda do ChatGPT para entender os processos-chave da inicialização de um servidor HTTP.  
Utilizei-o em conjunto com o [Beej's Guide](https://beej.us/guide/bgnet/) para encher o código de comentários conforme meu entendimento — ato que vai me auxiliar bastante sempre que eu esquecer algum aspecto, uma vez que poderei apenas retornar a esta versão e dar uma olhada.

Por enquanto, meu código começa imprimindo "hello, world!" no terminal, em seguida cria um servidor TCP ouvindo na porta 8080.  
Ele fica aguardando conexões de usuários e, quando um usuário se conecta, aceita a conexão e inicia uma goroutine para atendê-la.  
Dentro dessa goroutine, o servidor lê a requisição enviada pelo usuário, imprime o texto da requisição no terminal, constrói uma resposta HTTP simples com uma mensagem "hello!" dentro de uma página HTML e envia essa resposta de volta para o usuário, finalizando ao fechar a conexão.

Ainda não terminei por hoje, mas vou dar push para deixar essa versão primitiva do servidor focada em aprendizado salva.
