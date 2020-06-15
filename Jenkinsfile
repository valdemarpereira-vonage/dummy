pipeline {
    agent { docker { image 'golang' } }

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