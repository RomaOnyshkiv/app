#!/usr/bin/env groovy
pipeline {
    agent any

    stages {
        stage('Build app') {
            steps {
                sh "nohup go run api.go &"
            }
        }
    }
