# PATOS/POMBO PSEL 2.0

Tá sempre aberto, só enviar o PR

---

### Hey! Bem vindo(a) a segunda edição do PATOS/POMBO PSEL! **Clap Clap 👏👏**

![GIF](images/duck-dance.gif)

Durante os quase **Dois** primeiros anos do PATOS este processo de entrada de todos era um unico desafio, a criação de um **reverse proxy**, na prática o que queriamos era entregar algo "baixo nível" que deveria ser feito sem a ajuda de bibliotecas e sem ser nas **linguagens fácieis** como Python ou JavaScript.  
E o objetivo de vocês era justamente receber algo difícil e não tão comum e saber lidar com o "Ok, eu não sei nada disso" principalmente indo para o "Eu não sei nada disso, mas vou aprender, eu não tenho medo de errar e aprender algo novo".

Ficamos muito felizes com o resultado tanto na quantidade de pessoas, quanto na qualidade de cada uma delas. Descobrimos que sim, o processo seletivo funcionou bem e conseguimos trazer pessoas incríveis para o time, que de fato se encaixaram na cultura e no jeito PATOS de ser.  
Ficamos mais felizes ainda que não precisamos também de nenhuma bullshitagem de seleção como dinamicas de grupo, testes de personalidade, perguntas capciosas ou qualquer outra coisa que não fosse um bom desafio técnico.

Contudo, nós sabiamos que o processo poderia ser melhor. Antes sendo apenas um desafio específico deixavamos de lado muitas outras áreas de atuação do PATOS e seus integrandes. Um reverse proxy foca muito em redes que de fato é uma das nossas grandes áreas de atuação, mas e o resto? E o pessoal de Segurança?.. Hardware, SRE, Open Source?  
Além disso, acreditamos que chegamos em um platô (de fato elevado) de qualidade no antigo desafio. Os ultimos processos seletivos tiveram uma qualidade de entrega tão alta (diga-se de passagem, enviados todos por bixos) que não saberiamos mais como esperar menos ou aumentar o desafio para mais entregas.

Pensando nisso, resolvemos criar o **PATOS PSEL 2.0**. Um processo seletivo que abrange mais áreas do PATOS e que permite que você escolha o desafio que mais se encaixa com o seu perfil e com a sua vontade de aprender.

----

### Quem pode participar?

### O que é o PATOS/POMBO?

#### POMBO e PATOS é a mesma coisa?

#### Como funciona a hierarquia no PATOS/POMBO?

#### O que fazemos dentro do PATOS


### Como funciona o PATOS/POMBO PSEL 2.0?

Na segunda edição do PATOS PSEL, você poderá escolher entre 3 desafios diferentes, cada um focado em uma área de atuação do PATOS. Os desafios são:

- 1. Segurança & Redes - Firewall simples em UserSpace
- 2. Redes - Load Balancer
- 3. Segurança - Histórico de CTF
- 4. Open Source / Comunidade - Contribuição para um projeto open source

#### O que esperamos
- Pesquise, saiba lidar com problemas dificeis e busque aprender coisas novas
  


### 1. Segurança & Redes - Firewall simples em UserSpace

### 2. Redes - Load Balancer
Você deve fazer um load balancer do **zero**, lidando com as conexões e o redirecionamento na mão, sem usar qualquer lib que te auxilie. Além disso, vale ressaltar que você também não pode usar **nenhuma** lib para te ajudar no parsing das requests, ou seja, coisas desse gênero:
```python
from balancer import balance
import parser

server.balance()

parser.parse(request)
```
são completamente proíbidas. Queremos descobrir o quão longe vocês estão dispostos a aprenderem sozinhos, mesmo que tenham que reinventar a roda. O Load Balancer deve funcionar em cima de um servidor de arquivos que pode receber e armazenar arquivos de qualquer tipo, a sua escolha.

Para que você tenha uma ideia básica, o funcionamento de um Load Balancer, pode ser simplificado nesta imagem:

![IMAGE](images/loadbalancer.png)

Este processo seletivo é uma jornada de aprendizado, parte dele é descobrir mais detalhes sobre o que é e como funciona um Load Balancer, para que a sua entrega seja uma entrega com qualidade o suficiente para garantir seu ingresso no PATOS.

Lembrando que você só deve utilizar libs **extremamente** necessárias para que seu projeto rode, como por exemplo libs que conectam seu programa a um socket.

Por fim, como último empecilho, linguagens de extremo alto nível como python e javascript estão banidas neste processo seletivo. Recomendamos que você faça seu projeto em linguagens de baixo nível, como essas:

- `C/C++`
- `ASM`
- `Go`
- `Rust`
- `Zig`
- `Clojure`
- `Erlang`
- Qualquer outra desde que **não** tenha muita coisa pronta.

(Em especial, recomendamos C++/GO pois são linguagens mais parecidas com C mas com algumas regalias como strings, vetores e etc.)

Nós recomendamos fortemente que você não se limite a fazer *apenas* o que nós pedimos. Sinta-se livre para adicionar um diferencial - que pode ser qualquer coisa, desde que relevante - ao seu código.

As entregas serão individuais, mas sinta-se livre para discutir sobre o PSEL em grupo e para olhar as respostas de outros participantes. 

Sobre IA, nós desencorajamos o uso para a entrega deste processo seletivo. "A IA é apenas uma ferramenta. Ela é o **MEIO** e nunca o **FIM**". Se você decidir utilizar, utilize com sabedoria.

``O código deve ser entregue em um repositório do GITHub (que deve ser um fork deste repositório aqui), o qual deve conter um README.md``

Ao fim do processo, caso seu projeto seja aprovado, haverá uma entrevista individual.

#### Pontos de Avaliação:
- **HTTP Compilant** (conseguir acessar pelo navegador)
- **Documentação**
- **Colaboração** (documentar fontes de informações/código e informar sua jornada)
- **Organização e versionamento de código**
- **Experiência no Geral, não apenas um código**

> Vale ressaltar que os diferenciais que você adicionar no seu código também contam para a avaliação.

Finalmente, tenha em mente que:
- Você pode e deve contatar qualquer membro do PATOS em caso de dúvidas sobre o PSEL.
- Você pode deixar sua dúvida pública para outras pessoas que desejam fazer o PSEL mandando-a em [Issues](https://github.com/patos-ufscar/psel/issues)
- Não se sinta pressionado a fazer tudo, foque no que se sente confortável.
- Envie mesmo se não conseguir todas as partes essenciais, documente suas dificuldades.
- No seu README descreva como foi fazer o processo seletivo, o que você aprendeu, etc. **Documente sua jornada**.


``Boa sorte!``

> Lembrando que o processo é pra ser bem de boa, queremos ver até onde conseguem ir/se empurram, sem preocupação em fazer todos os essenciais.

### 3. Segurança - Histórico de CTF

Se você se interessa mais pela área de segurança e já tem experiência prévia, nós temos boas notícias para você! Você pode entrar no PATOS **sem ter que fazer o firewall ou o load balancer**: Basta que você tenha participado de algum CTF e consiga comprovar que você resolver, corretamente, pelo menos 1 challenge (mas sinta-se a vontade para mostrar mais challenges resolvidos!).

#### O que fazer para concorrer nesta categoria do PSEL?

Você ainda deve dar um fork neste github e depois mandar um pull request quando tudo estiver concluído. Você deve **obrigatóriamente** descrever os seguintes aspectos no seu README:
- Nome do CTF e data que ele ocorreu
- Tipo do challenge que você resolveu (RevEng, Web, Pwd, etc)
- Descrever o seu raciocínio para resolver o challenge (por onde você começou, quais vulnerabilidades você encontrou e como você as encontrou, etc.)
- Explicar as técnicas utilizadas para explorar as vulnerabilidades encontradas
- Caso você tenha utilizado um script para resolver o challenge, também explique como e por que você fez o script.

Além disso, o seu repositório deve conter uma cópia dos scripts utilizados para resolver o challenge (se possível, os arquivos do próprio challenge também), e prints comprovando que você resolveu o challenge corretamente.

> Esses serão os principais pontos de avaliação do seu processo seletivo.

Nós gostaríamos também que você nos contasse um pouco da sua história com a segurança da informação e as competições de CTF. Conte um pouco da sua jornada para nós!

Se você tiver feito tudo corretamente, nós o convocaremos para uma entrevista, assim como nos outros desafios!

``Boa Sorte!``

### 4. Open Source / Comunidade - Contribuição para um projeto open source

Para aqueles que já possuem experiência prática em projetos open source, o PATOS também decidiu facilitar sua entrada para o grupo! Você não precisa fazer o firewall nem o load balancer, também: basta que você comprove que você contribuiu efetivamente para algum projeto open source.

#### O que fazer para concorrer nesta categoria do PSEL?

O fork deste repositório continua sendo obrigatório, para que possamos analisar os envios de todos de uma forma mais padronizada. Quando tudo estiver pronto, mande um pull request e nós avaliaremos o conteúdo da sua aplicação para o PSEL.

Você deve **obrigatóriamente** fazer um README contendo os seguintes requisitos:

- Nome do projeto Open Source e linguagens utilizadas nele
- Link do repositório do projeto
- Link do seu pull request aceito
- Explicação detalhada da sua contribuição
- Detalhes de como você descobriu esse projeto e como você decidiu fazer sua contribuição para o mesmo, descrevendo os desafios que você enfrentou e sua experiência num geral durante processo de contribuição

> Esses são os principais pontos de avaliação do seu processo seletivo.

Se você tiver feito tudo corretamente, nós o convocaremos para uma entrevista, assim como nos outros desafios!

``Boa Sorte!``

### Ajuda

Achou alguma modalidade do PSEL muito difícil? Não fique desanimado, a dificuldade é proposital, mas garantimos que o processo seletivo **não é impossível**. Você sempre pode perguntar para algum membro do PATOS sobre dicas para fazer o PSEL e também discutir com outros candidatos sobre como cada um está fazendo o processo seletivo.

**Mas lembrem-se, a entrega do processo seletivo é INDIVIDUAL, ou seja: vocês podem se ajudar a fazer o PSEL, mas cada um deve entregar o seu!** 

> Vale ressaltar também que não é uma boa ideia plagiar o projeto do coleguinha, já que nós sempre comparamos os projetos.

Abaixo disponibilizamos alguns links para sites que podem ser úteis para o andamento do projeto de vocês.

#### Links de apoio