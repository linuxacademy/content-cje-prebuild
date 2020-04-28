pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
		        sh label: '', script: '''docker create --name=once makeindex bash)
docker cp once:/go/index.html .'''   
                     }
		    
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
