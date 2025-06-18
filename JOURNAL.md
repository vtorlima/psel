Dia 1 
No começo do semestre, eu tinha iniciado esse projeto, mas acabei priorizando outras atividades e o deixei de lado. Retomando agora, decidi começar do zero, mesmo ja tendo parte do código escrito anteriormente, com o objetivo de reorganizar tudo de forma mais clara, utilizando como referência o que já tinha feito antes. 
Minha intenção principal com esse repositório não é produzir um código limpo ou de facil leitura, mas sim um código bastante comentado, com explicações detalhadas, de forma que, além de facilitar o aprendizado, eu consiga revisitar qualquer parte no futuro e enetendr exatamente o que está acontecendo

No primeiro dia, construí a estutura básica de um servidor TCP em Go, que escuta na porta 8080 e aceita conexões de usuarios. Assim que a conexão é aceita, uma goroutine é iniciada para lidar com ela de forma paralela, o que permite que o servidor continue escutando novas conexões enquanto trata as já existentes. Dentro da função de tratamento, li os dados enviados pelo cliente e imprimi no terminal o conteúdo da requisição recebida. 

Agora, o próximo passo é ler e entender o conteúdo da requisição para conseguir retornar o arquivo que está sendo pedido. 