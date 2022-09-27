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
    // Perform Build
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go env GOCACHE=/tmp/.cache'
                sh 'GOOS=linux GOARCH=amd64 go build -o time-app .'
                sh 'chmod +x time-app'
            }
        }
    // Docker Build
        stage("docker build"){
            steps{
               script{
                    dockerImage = docker.build dockerimagename
               }
            }
        }
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
    }
}
//         stage('deliver') {
//             agent any
//             steps {
//                 withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
//                 sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
//                 sh 'docker push shadowshotx/product-go-micro'
//                 }
//             }
//         }
//     }


