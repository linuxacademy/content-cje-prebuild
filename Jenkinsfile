pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
                         sh 'cat index.html'
                     }
		    
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
