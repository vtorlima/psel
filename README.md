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

TODO O MUNDO! Não importa de bixo a veterano de dentro ou fora da UFSCar, de dia noite, nas ferias ou durante as aulas, se você tem vontade de aprender e acha que o PATOS é um lugar onde você pode aprender, se desenvolver e contribuir, esse processo seletivo é pra você!

### O que é o PATOS e POMBO? São a mesma coisa? O processo seletivo é para qual dos dois?

PATOS é o grupo de Open Source da UFSCar, além de ser um grupo de estudos, um grupo de pesquisa, um grupo de desenvolvimento, um grupo de pessoas que se juntam para aprender e contribuir tanto com código tnato com conhecimento. Outra grande frente nossa são os Aulões, onde a gente ensina o que a gente sabe para a comunidade, seja ela interna ou externa.  
Já o POMBO é o grupo de Cibersegurança da UFSCar, sua existencia é anterior ao PATOS, mas hoje em dia é um subgrupo do PATOS, 'a frente' de segurança.  
Este processo seletivo é para ambos, e você pode escolher qualquer desafio para entrar no PATOS, que consequentemente te dá acesso ao POMBO, caso queira participar da frente de segurança.

#### Como funciona a hierarquia no PATOS/POMBO?

Não há, espero ter ajudado 👍!  
Simplesmente não achamos que faça sentido. O PATOS funciona como uma comunidade de pessoas que se juntam para aprender e contribuir (com código, conhecimento, recomendação e memes), não existe um 'nível' ou 'hierarquia' dentro do grupo, todos são iguais e tem voz ativa desde o dia um.
Além disso, não gostamos da forma engessada e ultrapassada que outros grupos tem, sem contar o tanto de confusão e discórdia que isso gera entre os membros. (Sério que vocês querem receber ordens de outro aluno sem nem receber um salário? Não, obrigado ~Marlon).
Se você entrar no PATOS saiba que você é membro e no máximo vamos ter um 'bixo' e um 'veterano', mas isso é só para diferenciar quem entrou recentemente de quem já tem mais tempo, não tem nada a ver com 'nível' ou 'hierarquia' e você pode levantar a mão puxar algo ou criticar algo também!

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

### 4. Open Source / Comunidade - Contribuição para um projeto open source

### Ajuda

