#!/usr/bin/env groovy
pipeline {
    agent any

    stages {
        stage('Build app') {
            steps {
                sh "cd api/"
                sh "nohup go run api.go &"
            }
        }
    }
}
