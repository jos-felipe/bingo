# binGO

Um programa simples em Go para realizar sorteios aleatórios a partir de uma lista de participantes em um arquivo CSV.

## Descrição

O binGO é uma ferramenta de linha de comando escrita em Go que permite sortear um número específico de ganhadores a partir de uma lista de participantes fornecida em um arquivo CSV. O programa garante um sorteio justo e aleatório, ideal para eventos, meetups, ou qualquer situação que requeira a seleção imparcial de ganhadores.

## Requisitos

- Go 1.16 ou superior
- Sistema operacional: Windows, macOS ou Linux

## Instalação

### Opção 1: Compilar a partir do código fonte

1. Clone este repositório ou baixe o arquivo `binGO.go`
2. Navegue até o diretório onde o arquivo está localizado
3. Compile o programa usando o comando:

```bash
go build -o binGO binGO.go
```

### Opção 2: Instalar via `go install`

```bash
go install github.com/seu-usuario/binGO@latest
```

## Uso

```bash
./binGO arquivo.csv numero_de_ganhadores
```

Onde:
- `arquivo.csv` é o caminho para o arquivo CSV contendo a lista de participantes (um nome por linha)
- `numero_de_ganhadores` é a quantidade de ganhadores que deseja sortear

## Exemplo

1. Crie um arquivo CSV chamado `meetup.csv` com o seguinte conteúdo:
```
Ana
Bob
Charlie
David
Elena
Fernando
Gabriela
```

2. Execute o programa:
```bash
./binGO meetup.csv 3
```

3. O resultado será algo como:
```
1 - Bob
2 - Charlie
3 - Ana
```

O resultado será diferente a cada execução, pois o sorteio é aleatório.

## Estrutura do arquivo CSV

O arquivo CSV deve:
- Ter um nome por linha
- Não ter cabeçalho (o programa considera que todas as linhas são nomes de participantes)
- Estar codificado em UTF-8 (recomendado)

## Como funciona

O binGO funciona seguindo estas etapas:
1. Lê a lista de participantes do arquivo CSV especificado
2. Verifica se há participantes suficientes para o número de ganhadores solicitado
3. Embaralha a lista de participantes usando um algoritmo aleatório
4. Seleciona os primeiros N participantes após o embaralhamento
5. Exibe os ganhadores numerados na tela

## Detalhes técnicos

- O programa utiliza a biblioteca padrão de Go, sem dependências externas
- A aleatoriedade é garantida usando o pacote `math/rand` com uma semente baseada no tempo atual
- Tratamento adequado de erros para entradas inválidas ou arquivos inexistentes

rm## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.
