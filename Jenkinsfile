#!/usr/bin/env groovy
pipeline {
    agent any

    stages {
        stage('Build app') {
            steps {
                sh """
                cd api && go mod download && go mod tidy && go mod verify
                """
            }
        }
        stage ('Start app') {
            steps {
                withEnv(['JENKINS_NODE_COOKIE=dontkill']) {
                    sh "cd api && nohup go run api.go &"
                }
            }
        }
    }
}
