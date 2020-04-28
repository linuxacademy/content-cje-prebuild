pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
                         sh 'cat /go/index.html'
                     }
		    
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
