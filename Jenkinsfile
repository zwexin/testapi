#!/usr/bin/env groovy
pipeline {
    agent any

    environment {
        DOCKER_IMAGE_BASE = 'zwexin/testapi'
    }

    stages {
        stage('notify') {
            steps {
                sh '/usr/local/bin/dingdingnotify start'
            }
        }
        stage('build') {
            steps {
                sh 'go build server.go'
            }
        }
        stage('docker') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'master') {
                        def customImage = docker.build("${DOCKER_IMAGE_BASE}:${env.BUILD_NUMBER}")
                        customImage.push()

                        customImage.push('latest')
                    } else {
                        def customImage = docker.build("${DOCKER_IMAGE_BASE}:${env.BUILD_NUMBER}-${env.BRANCH_NAME}")
                        customImage.push()
                    }
                }
            }
        }
        stage('kubernetes') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'master') {
                        sh ("""
                              sed -i -E \"s#PROJECTVESRION#v${env.BUILD_NUMBER}#g\"  testapi-deployment.yml
                              sed -i -E \"s#DOCKERIMAGENAME#${DOCKER_IMAGE_BASE}:${env.BUILD_NUMBER}#g\"  testapi-deployment.yml 
                              kubectl apply -f testapi-deployment.yml
                        """)
                    } else {
                        sh ("""
                              sed -i -E \"s#PROJECTVESRION#v${env.BUILD_NUMBER}-${env.BRANCH_NAME}#g\"  testapi-deployment.yml
                              sed -i -E \"s#DOCKERIMAGENAME#${DOCKER_IMAGE_BASE}:${env.BUILD_NUMBER}-${env.BRANCH_NAME}#g\"  testapi-deployment.yml
                              kubectl apply -f testapi-deployment.yml
                        """)
                    }
                }
            }
        }
        stage('istio') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'master') {
                        sh ("""
                              sed -i -E \"s#PROJECTVESRION#v${env.BUILD_NUMBER}#g\"  testapi-istio.yml
                              kubectl apply -f testapi-istio.yml
                        """)
                    } else {
                        sh ("""
                              sed -i -E \"s#PROJECTVESRION#v${env.BUILD_NUMBER}-${env.BRANCH_NAME}#g\"  testapi-istio.yml
                              kubectl apply -f testapi-istio.yml
                        """)
                    }
                }
            }
        }
    }
    post {
        success {
            sh '/usr/local/bin/dingdingnotify success'
        }
        unsuccessful {
            sh '/usr/local/bin/dingdingnotify unsuccessful'
        }
    }
}
