job("build dockerfile"){
	startOn {
        gitPush {
            tagFilter {
                +"release/*"
            }
        }
    }
    git("selloora-backend/backend") {
        cloneDir = "backend"
    }
    git("selloora-frontend/frontend") {
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