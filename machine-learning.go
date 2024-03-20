/*
  Nesse script iremos criar um modelo de aprendizado de máquina
  utilizando a linguagem GO.

  O nosso objetivo é gerar um modelo de classificação, para isso
  iremos utilizar o KNN para classificar os sentimentos de clientes.

*/

package main

import (
  "fmt"

  "github.com/sjwhitworth/golearn/base"
  "github.com/sjwhitworth/golearn/evaluation"
  "github.com/sjwhitworth/golearn/knn"
)

func main() {

  fmt.Println("Machine Learning with Go Lang")

  // carregando o conjunto de dados
  fmt.Println("\nData loading")
  rawData, err := base.ParseCSVToInstances("dataset.csv", true)
  if err != nil {
    panic(err)
  }

  // inicializando o modelo KNN
  fmt.Println("\nInitializing the KNN Model")
  cls := knn.NewKnnClassifier("euclidean", "linear", 2)

  // dividindo os dados em treino e teste - 70% para treino e 30% para teste 
  fmt.Println("\nDivision of Data into Training and Testing")
  trainData, testData := base.InstancesTrainTestSplit(rawData, 0.30)

  // realizando o treinando do modelo
  cls.Fit(trainData)

  // realizando as previsões com o modelo treinado
  fmt.Println("\nMaking Predictions with the Trained Model")
  predictions, err := cls.Predict(testData)
  if err != nil {
    panic(err)
  }

  // imprimindo as previsões
  fmt.Println(predictions)

  // imprimindo as métricas
  fmt.Println("\nModel Metrics")

  // matriz de confusão
  confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
  if err != nil {
    // imprimindo uma mensagem de erro se por acaso houver algum problema na geração da matriz de confusão - "Não foi possível gerar a matriz de confusão"
    panic(fmt.Sprintf("Unable to generate confusion matrix: %s", err.Error()))
  }

  // acurácia do modelo
  fmt.Println(evaluation.GetSummary(confusionMat))
}





