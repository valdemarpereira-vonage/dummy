pipeline {
    agent {
        docker { image 'golang' }
    }

    stages {
            stage('Build Stage') {
                steps {
                        script {

                                    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
                                        withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
                                        env.PATH="${GOPATH}/bin:$PATH"
                                        }
                                    }
                        }
                }
            }


             post {
                    always {
                        script {
                            echo "Post Script"
                        }
                    }
                }
}


// pipeline {
//     agent { docker { image 'golang' } }
//     environment {
//         GOPATH = ${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}
//     }
//     stages {
//         stage('build') {
//             steps {
//                 sh '''
//                    ls
//                    go version
//                    go test -v
//                    '''
//              }
//         }
//     }
// }