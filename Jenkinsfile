pipeline {
    // install golang 1.19.1 on Jenkins node
    agent any
    tools {
        go '1.19.1'
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
                sh 'sudo chmod -R 777 ./*'
                sh 'GOOS=linux GOARCH=amd64 go build -o time-app .'
                sh 'chmod +x time-app'
            }
        }
    // Docker Build
        stage("docker build"){
            steps{
                echo 'BUILD DOCKER IMAGES'
                sh 'docker build .'
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


