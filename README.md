# Odd Filter

`odd-filter` é uma aplicação de linha de comando (CLI) escrita em Go que lê uma lista de números fornecidos pelo usuário, separados por espaço, e exibe apenas os números ímpares. Entradas que não são números inteiros válidos são identificadas e listadas separadamente.

## Funcionalidades

*   Recebe uma string de números separados por espaço como entrada.
*   Identifica e exibe os números ímpares.
*   Identifica e exibe as entradas que não são números inteiros válidos.
*   Interface de linha de comando interativa.
*   Empacotado com Docker para fácil compilação e execução.

## Lógica do Código

O programa principal (`main.go`) consiste em duas funções principais:

1.  `main()`:
    *   Configura um leitor de `bufio` para capturar a entrada do usuário via `os.Stdin`.
    *   Solicita ao usuário que digite os números.
    *   Lê a linha de entrada do usuário.
    *   Trata erros de leitura e o caso de entrada vazia.
    *   Remove espaços em branco extras e divide a string de entrada em um slice de strings (cada string representando um potencial número).
    *   Chama a função `oddFilter` para processar os números.
    *   Imprime os números ímpares encontrados e/ou as entradas inválidas.

2.  `oddFilter(args []string) ([]int, []string)`:
    *   Recebe um slice de strings (`args`) como entrada.
    *   Inicializa dois slices vazios: `oddNumbers` (para inteiros ímpares) e `invalidNumbers` (para strings inválidas).
    *   Itera sobre cada string no slice `args`:
        *   Tenta converter a string para um inteiro usando `strconv.Atoi()`.
        *   Se a conversão falhar (retornar um erro), a string original é adicionada ao slice `invalidNumbers`.
        *   Se a conversão for bem-sucedida, verifica se o número é ímpar (`num % 2 != 0`).
        *   Se for ímpar, o número é adicionado ao slice `oddNumbers`.
    *   Retorna os slices `oddNumbers` e `invalidNumbers`.

## Como Usar

### Pré-requisitos

*   [Go](https://golang.org/dl/) (versão 1.21 ou superior, conforme especificado no `go.mod`) - para compilação local.
*   [Docker](https://www.docker.com/get-started) - para compilação e execução via container.

### Compilando e Executando Localmente (com Go instalado)

1.  Clone o repositório:
    ```bash
    git clone https://github.com/DaviRSouza/odd-filter.git
    cd odd-filter
    ```

2.  Execute a aplicação:
    ```bash
    go run main.go
    ```
    Ou compile e depois execute:
    ```bash
    go build -o oddfilter_app main.go
    ./oddfilter_app
    ```

### Compilando e Executando com Docker

O projeto inclui um `Dockerfile` que utiliza um build para criar uma imagem Docker leve.

1.  **Construa a Imagem Docker:**
    Navegue até o diretório raiz do projeto (onde o `Dockerfile` está localizado) e execute:
    ```bash
    docker build -t odd-filter-app .
    ```
    Isso irá construir uma imagem Docker com o nome `odd-filter-app`.

2.  **Execute a Aplicação em um Container Docker:**
    Após a imagem ser construída, você pode executar a aplicação dentro de um container Docker:
    ```bash
    docker run -it --rm odd-filter-app
    ```
    *   `-it`: Aloca um pseudo-TTY e mantém o STDIN aberto, necessário para a interação com a aplicação.
    *   `--rm`: Remove automaticamente o container quando ele é finalizado.
    *   `odd-filter-app`: O nome da imagem que você construiu.

    A aplicação irá iniciar no container, e você poderá interagir com ela no seu terminal:

    ```
    Digite os números separados por espaço (ex: 1 2 3 4) e pressione Enter:
    > 1 2 abc 3 4 xyz 5
    Números ímpares encontrados:
    1 3 5
    Valores invalidos que foram ignorados:
    abc xyz
    ```
