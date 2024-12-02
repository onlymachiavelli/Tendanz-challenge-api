pipeline {
    agent any

    environment {
        DOCKER_REGISTRY = 'my-docker-registry'
        APP_IMAGE = 'go-server:latest'
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    sh 'docker-compose build'
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Add your test commands here
                    sh 'docker-compose up -d'
                    sh 'docker-compose exec go-server go test ./...'
                    sh 'docker-compose down'
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh 'docker login -u $DOCKER_USER -p $DOCKER_PASSWORD $DOCKER_REGISTRY'
                    sh "docker tag go-server $DOCKER_REGISTRY/$APP_IMAGE"
                    sh "docker push $DOCKER_REGISTRY/$APP_IMAGE"
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    sh 'docker-compose up -d'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
