pipeline {
    environment{
        dockerimagename = "lnanjangud653/time-app"
        dockerImage = ""
    }
    // install golang 1.19.1 on Jenkins node
    agent any
    tools {
        go '1.19'
    }
    // Define different states for golang app
    stages {
    // Perform Unit Test
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
              //  sh 'go test'
            }
        }
    // Perform Functional Test
        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
            }
        }
    // Build Go Binary
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go env GOCACHE=/tmp/.cache'
                sh 'GOOS=linux GOARCH=amd64 go build -o time-app .'
                sh 'chmod +x time-app'
            }
        }
    //  Build Docker dockerImage
        stage("docker build"){
            steps{
               script{
                    dockerImage = docker.build dockerimagename
               }
            }
        }
    // Docker Push
        stage('Pushing Image'){
           environment {
                registryCredential = 'dockerhubcreds'
           }
           steps{
            script{
                docker.withRegistry('https://registry.hub.docker.com', registryCredential){
                    dockerImage.push("latest")
                }
            }
           }
        }
    // Test Kubernetes
        stage ("Test Kubernetes"){
        steps{
            sh 'kubectl get ns --kubeconfig=/home/bitnami/config'
        }
        }
    // Test Tanzu
       stage("Test Tanzu"){
       agent {
            label 'tanzu-mgmt'
       }
       steps{
            sh 'tanzu cluster list --include-management-cluster'
           }
       }
    }
}

