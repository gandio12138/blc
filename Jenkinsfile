pipeline {
  agent any
  stages {
    stage('test') {
      steps {
        sh '''chmod 755 ./blockchain
./blockchain printChain
./blockchain addBlock -data "Send 1 BTC to Ivan"
./blockchain addBlock -data "Pay 0.31337 BTC for a coffee"
./blockchain printChain'''
      }
    }

  }
}