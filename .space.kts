job("Build and push Docker") {
    container("openjdk:11") {
    kotlinScript { api ->
        api.space().projects.automation.deployments.start(
            project = api.projectIdentifier(),
            targetIdentifier = TargetIdentifier.Key("backend"),
            version = "1.0.0",
            // automatically update deployment status based on a status of a job
            syncWithAutomationJob = true
        )
    }
}
    host("Build artifacts and a Docker image") {
        // generate artifacts required for the image
        shellScript {
            content = """
                ./generateArtifacts.sh
            """
        }

        dockerBuildPush {
            // Docker context, by default, project root
            context = "docker"
            // path to Dockerfile relative to project root
            // if 'file' is not specified, Docker will look for it in 'context'/Dockerfile
            file = "docker/config/Dockerfile"
            labels["vendor"] = "mycompany"
            args["HTTP_PROXY"] = "http://10.20.30.1:123"

            val spaceRepo = "mycompany.registry.jetbrains.space/p/prjkey/mydocker/myimage"
            // image tags for 'docker push'
            tags {
                +"$spaceRepo:1.0.${"$"}JB_SPACE_EXECUTION_NUMBER"
                +"$spaceRepo:latest"
            }
        }
    }
}