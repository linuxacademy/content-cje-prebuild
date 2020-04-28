pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
		        sh label: '', script: '''VAR=$(docker create makeindex bash)
docker cp $VAR:/go/index.html .'''   
                     }
		    
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
