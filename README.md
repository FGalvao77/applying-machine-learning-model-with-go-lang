## Projeto prático de utilização da `Go Lang` em Aprendizado de Máquina.
---
---

Nesse projeto iremos utilizar um modelo supervisionado, nesse caso o modelo escolhido foi o **KNN**. Onde para cada observação, temos três notas como atributo do conjunto de dados e, respectivamente para cada observação é rotulado (sentimento) como _positivo_ ou _negativo_. 

### Segue as etapas para executar os _scripts_ 
---

- 1a etapa - Arquivo "data-loading.go":

    1. abra o seu terminal ou _shell_ de comando.

    2. já no terminal ou _shell_, execute o comando:
       
            go mod init data-loading.go

    4. e em seguida, execute o comando:
       
            go run leitura_dados.go

    OBS: note que nada aconteceu!!!
         Isso se deve que, precisamos instalar o pacote e a dependência.
         Para solução do problema, execute o comando sugerido de aviso que consta no terminal.

            go get github.com/go-gota/gota/dataframe

   Agora podemos executar o comando.
  
            go run leitura_dados.go

Se tudo deu certo, você verá a todas as informações do conjunto de dados.


- 2a etapa - Arquivo "machine-learning.go":

    1. novamente vá até o seu terminal ou _shell_

    2. e execute o comando:
       
            go mod machine_learning.go

    4. e por fim, execute o comando:
       
            go run machine_learning.go

    OBS: note novamente que nada aconteceu!!!
         Com anteriormente, precisamos instalar os pacotes e as dependências necessárias.
         Então execute os comandos sugeridos de aviso que consta no terminal.

            go get github.com/sjwhitworth/golearn/base
            go get github.com/sjwhitworth/golearn/evaluation
            go get github.com/sjwhitworth/golearn/knn

  Agora podemos executar o comando
  
            go run leitura_dados.go

Se tudo deu certo, você terá a todas as informações no terminal! ;)











