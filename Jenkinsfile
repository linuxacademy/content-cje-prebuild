pipeline{
	agent { dockerfile true }
	stages {
		stage ('index'){
		     steps {
			docker.image('makeindex').inside("go build ./src/makeindex.go")
			docker.image('makeindex').inside("makeindex")
			docekr.image('makeindex').inside("go version")
		     }
		     post {
                        success {
                          archiveArtifacts 'index.html'            
                        }
                     }
                }
        }
}
             
