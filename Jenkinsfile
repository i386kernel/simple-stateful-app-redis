pipeline {
    environment{
        dockerimagename = "lnanjangud653/time-app"
        dockerImage = ""
        imageVer = "1.0.0"
    }
    // install golang 1.19.1 on Jenkins node
    agent any
    tools {
        go '1.19'
    }
    // Define different states for golang app
    stages {
    // Perform Unit Test
        stage("Unit-Test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                //sh 'go test'
            }
        }
    // Perform Functional Test
        stage("Functional-Test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
            }
        }
    // Build Go Binary
        stage("App Build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go env GOCACHE=/tmp/.cache'
                sh 'GOOS=linux GOARCH=amd64 go build -o time-app .'
                sh 'chmod +x time-app'
            }
        }
    //  Build Docker dockerImage
        stage("Docker Build"){
            steps{
               script{
                    dockerImage = docker.build dockerimagename
               }
            }
        }
    // Docker Push
        stage('Docker Push'){
           environment {
                registryCredential = 'dockerhubcreds'
           }
           steps{
            script{
                docker.withRegistry('https://registry.hub.docker.com', registryCredential){
                    dockerImage.push("2.0.0")
                }
            }
           }
        }
    // Test Kubernetes
        stage ("Deploy on Tanzu"){
        agent {
                    label 'tanzu-mgmt'
               }
        steps{
            sh 'kubectx tkg-wl-prod-admin@tkg-wl-prod'
            sh 'kubectl create ns time-app'
            sh 'kubens time-app'
            sh 'kubectl apply -f kubemanifest.yml'
        }
        }
       }
}


