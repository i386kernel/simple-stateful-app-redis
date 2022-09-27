pipeline {
    // install golang 1.19.1 on Jenkins node
    agent any
    tools {
        go '1.19.1'
    }
//     environment {
//         GO114MODULE = 'on'
//         CGO_ENABLED = 0
//         GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
//     }
//     stages {
//         stage("unit-test") {
//             steps {
//                 echo 'UNIT TEST EXECUTION STARTED'
//                 sh 'go test'
//             }
//         }
//         stage("functional-test") {
//             steps {
//                 echo 'FUNCTIONAL TEST EXECUTION STARTED'
//                 sh 'make functional-tests'
//             }
//         }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'GOOS=linux GOARCH=amd64 go build -o time-app .'
                sh 'chmod +x time-app'
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
// }
