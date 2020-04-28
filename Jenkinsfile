pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
			sh 'go build ./src/makeindex.go'
			sh './makeindex'
		     }
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
