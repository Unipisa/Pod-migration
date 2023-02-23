pipeline {
    agent {
        docker {
            image 'golang:1.20'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    environment {
        GITHUB_TOKEN = credentials('github_token')
        DOCKER_REGISTRY = 'ghcr.io'
        DOCKER_IMAGE_PREFIX = 'unipi.it/'
        DOCKER_IMAGE_TAG = 'latest'
    }

    stages {
        stage('Build') {
            steps {
                dir('liqo') {
                    sh 'go build -o liqo'
                }

                dir('live-migration-operator') {
                    sh 'go build -o live-migration-operator'
                }
            }
        }

        stage('Publish') {
            steps {
                withCredentials([string(credentialsId: 'github-token', variable: 'GITHUB_TOKEN')]) {
                    sh 'docker login ${env.DOCKER_REGISTRY} -u ${env.GITHUB_ACTOR} -p ${env.GITHUB_TOKEN}'

                    dir('submodule-1') {
                        sh "docker build -t ${env.DOCKER_REGISTRY}/${env.DOCKER_IMAGE_PREFIX}-1:${env.DOCKER_IMAGE_TAG} ."
                        sh "docker push ${env.DOCKER_REGISTRY}/${env.DOCKER_IMAGE_PREFIX}-1:${env.DOCKER_IMAGE_TAG}"
                    }

                    dir('submodule-2') {
                        sh "docker build -t ${env.DOCKER_REGISTRY}/${env.DOCKER_IMAGE_PREFIX}-2:${env.DOCKER_IMAGE_TAG} ."
                        sh "docker push ${env.DOCKER_REGISTRY}/${env.DOCKER_IMAGE_PREFIX}-2:${env.DOCKER_IMAGE_TAG}"
                    }
                }
            }
        }
    }
}

