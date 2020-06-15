pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh '''
                   cd ~
                   pwd
                   cat /etc/passwd
                   whoami
                   ls
                   mkdir -p /var/jenkins_home/gocache
                   mkdir -p /var/jenkins_home/.config/go/env
                   export GOCACHE=/var/jenkins_home/gocache
                   export GOENV=/var/jenkins_home/.config/go/env

                   '''

             }
        }
    }
} 
