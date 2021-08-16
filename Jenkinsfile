pipeline {
  agent any
  stages {
    stage('build') {
      parallel {
        stage('build') {
          steps {
            sh '''source /etc/profile
go version
git branch -a'''
          }
        }

        stage('b') {
          steps {
            sh '''source /etc/profile
go version
git branch -a'''
          }
        }

      }
    }

  }
}