pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''#!/bin/bash -ilex
go version
ls
git branch -a'''
      }
    }

  }
}