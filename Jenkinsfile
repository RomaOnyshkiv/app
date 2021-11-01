#!/usr/bin/env groovy
pipeline {
    agent any

    stages {
        stage('Build app') {
            steps {
                sh """
                cd api
                chmod +x build.sh
                ./build.sh
                """
            }
        }
    }
}
