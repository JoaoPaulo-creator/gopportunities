Diretorios em go, sao entendidos como subpackages

### Exports

Faz se necessário criar funções onde a primeira letra de seu nome, deve ser maiusculo,
pois assim é indicado de que esssa função será exportada.

Funções com o nome todo em caixa baixa, são funções/variáveis locais

### Packages

Tudo que está no package, é acessível no package inteiro, ou seja, se dentro de um pacote houver vários arquivos .go,
suas funções serãos acessíveis pelos demais arquivos

### O que é um ponteiro

Um ponteiro é o endereço de algo que está na memória, como o endereço de uma variável.

### Body json

Se algum campo for passado no json, no momento da request, e este campo nao foi definidio nas funcoes de request,
o campo extra sera apenas ignorado, mantendo o app funcionando
