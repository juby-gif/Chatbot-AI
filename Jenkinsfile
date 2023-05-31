pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                // Perform build steps here
                sh 'npm install'  // Example build step for a Node.js project
            }
        }
        
        stage('Test') {
            steps {
                // Perform test steps here
                sh 'npm test'  // Example test step for a Node.js project
            }
        }
        
        stage('Deploy') {
            steps {
                // Perform deployment steps here
                sh 'npm run deploy'  // Example deployment step for a Node.js project
            }
        }
    }
}
