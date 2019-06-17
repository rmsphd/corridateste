RESULTADO DA CORRIDA
========
[![Go Report Card](https://goreportcard.com/badge/github.com/rmsphd/corridateste)](https://goreportcard.com/report/github.com/rmsphd/corridateste)

Dado o seguinte log de uma corrida de Kart:

[corrida.log](corrida.log)

Resultado esperado
------------------
* A partir de um input de um arquivo de log do formato acima, montar o resultado da corrida com as seguintes informações: **Posição Chegada**, **Código Piloto**, **Nome Piloto**, **Qtde Voltas Completadas** e **Tempo Total de Prova**.

Observações
------------
* A primeira linha do arquivo pode ser desconsiderada (Hora, Piloto, Nº Volta, Tempo Volta, Velocidade média da volta).
* A corrida termina quando o primeiro colocado completa 4 voltas


Bônus
-----
Não obrigatório. Faça apenas caso se identifique com o problema ou se achar que há algo interessante a ser mostrado na solução

**********************************************
* Descobrir a melhor volta de cada piloto
***********************************************
* Descobrir a melhor volta da corrida
************************************************
* Calcular a velocidade média de cada piloto durante toda corrida
************************************************************************
* Descobrir quanto tempo cada piloto chegou após o vencedor



SOLUÇÃO
=======
* Utilizado Golang 1.12.x
* Testes unitários
* Utilitário de linha de comando, binário nátivo para linux, windows e mac
* Somente utilizado a standard library de Golang


Pré-requisitos
==============

- Instalar Go lang 1.12.x
  - https://golang.org/dl/

## Alternativa
- Instalar 
  - Docker https://docs.docker.com/install/
- Configurar para linux
  - https://docs.docker.com/install/linux/linux-postinstall/

BUILD da solução
======

* Dentro da pasta do projeto
    * Com docker
        * docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:alpine ./build.sh
    * Com Golang 1.12.x
        * ./build.sh
    * Com Golang 1.12.x alternativo
        * go build -o interview

TESTES
======

- Rodar os testes unitários 
  - go test ./...
- Rodar os testes unitários com o relatório do cobertura
  - go test -coverprofile=coverage.out ./... 
  - go tool cover -html=coverage.out

EXECUÇÃO
========

* Escolha o binário que se encaixa com o seu ambiente,
    * ./interview-\<OS>-\<ARCH>
    * Exemplo: ./interview-darwin-amd64 
* Execute na linha de comando apontando para um arquivo válido
    * ./interview-\<OS>-\<ARCH> <arquivoEntrada.log> [arquivoSaiida.log]
        * <arquivoEntrada.log> obrigatório, é o arquivo que contem os dados de cada volta da corrida
        * [arquivoSaiida.log] opcional, se não informado será retornado na tela, mostra os resultado da corrida
    * Exemplos
        * ./interview-darwin-amd64 corrida.log
        * ./interview-darwin-amd64 corrida.log resultado.log

