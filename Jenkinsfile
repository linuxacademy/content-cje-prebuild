pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
			sh 'ls'
		     }
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
