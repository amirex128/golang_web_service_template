job("build dockerfile"){
	startOn {
        gitPush {
            tagFilter {
                +"release/*"
            }
        }
    }
    git("https://git.jetbrains.space/amirex128/selloora-backend/backend.git") {
        cloneDir = "backend"
    }
    git("https://git.jetbrains.space/amirex128/selloora-frontend/frontend.git") {
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