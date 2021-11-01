#!/usr/bin/env groovy
pipeline {
    agent any

    stages {
        stage('Build app') {
            steps {
                sh """
                chmod +x build.sh
                ./build.sh
                """
            }
        }
    }
}
