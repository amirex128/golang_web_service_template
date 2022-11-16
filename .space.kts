job("build"){
	startOn {
        gitPush {
            tagFilter {
                +"release/*"
            }
        }
    }
    git("backend") {
        cloneDir = "backend"
    }
    git("frontend") {
        cloneDir = "frontend"
    }
    
    container(displayName = "Show dirs", image = "ubuntu:latest") {
        shellScript {
            content = """
                echo Directory structure
                ls -R /mnt/space/work
                echo Working dir is
                pwd
            """
        }
    }
}