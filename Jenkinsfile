pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
			sh 'go version'
		     }
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
