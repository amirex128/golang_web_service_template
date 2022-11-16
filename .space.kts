job("build"){

    git("backend") {
        cloneDir = "backend"
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