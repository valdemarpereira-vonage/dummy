pipeline {
    agent { docker { image 'golang' } }
    environment {
        GOPATH = ${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}
    }
    stages {
        stage('build') {
            steps {
                sh '''
                   ls
                   go version
                   go test -v
                   '''
             }
        }
    }
}