pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''#!/bin/bash -ilex
git config --global http.version HTTP/1.1
go version
ls
git branch -a'''
      }
    }

  }
}