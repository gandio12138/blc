pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''source /etc/profile
go build -o blockchain'''
      }
    }

    stage('deps') {
      steps {
        sh '''source /etc/profile
go get -u github.com/boltdb/bolt'''
      }
    }

    stage('test') {
      steps {
        sh '''./blockchain printChain
./blockchain addBlock -data "Send 1 BTC to Ivan"
./blockchain addBlock -data "Pay 0.31337 BTC for a coffee"
./blockchain printChain'''
      }
    }

  }
}