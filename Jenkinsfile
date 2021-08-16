pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''source /etc/profile
go version
ls
make'''
      }
    }

  }
}