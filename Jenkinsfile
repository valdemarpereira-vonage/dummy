// pipeline {
//     agent {
//         docker { image 'golang' }
//     }
//
//     stages {
//             stage('Build Stage') {
//                 steps {
//                         script {
//
//                                     ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
//                                         withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
//                                         env.PATH="${GOPATH}/bin:$PATH"
//
//                                         sh "go version"
//                                         }
//                                     }
//                         }
//                 }
//             }
//     }
//
//
//              post {
//                     always {
//                         script {
//                             echo "Post Script"
//                         }
//                     }
//              }
// }

//      wrap([$class: 'BuildUser']) {
//            def user = env.BUILD_USER_ID
//         }
pipeline {
    agent { docker { image 'golang' } }
    environment {
        GOCACHE = "/tmp"

    }
    stages {
        stage('build') {
            steps {
            wrap([$class: 'BuildUser']) {
                  echo "${BUILD_USER}"
                  echo "${BUILD_USER_ID}"
                  echo "${BUILD_USER_EMAIL}"
                }

                sh '''
                    echo "${BUILD_USER}"
                          echo "${BUILD_USER_ID}"
                          echo "${BUILD_USER_EMAIL}"
                   ls
                   go version
                   go test -v
                   '''
             }
        }
    }
}