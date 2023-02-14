# go-module-small

A template for a small and simple go module.

## What Does This Contain?
- My standard [`.gitignore`](.gitignore) file. This is the gitignore I use, mostly generated from (gitignore.io)[https://gitignore.io], with some added miscellaneous files that I use locally.
- My standard [`Dockerfile`](Dockerfile) file. This is my super basic multi-stage Dockerfile, that uses the latest Go as the build-env and Google's Distroless as the final environment.
- My standard [`Makefile`](Makefile). I like to use a Makefile so that I don't have to always remember how to build, tag, and push new docker images to my registry. The Dockerfile doesn't actually contain anything other than build and push of docker images, and a build local binary step. Linting, testing, formatting is all taken care of outside of this.
- A [`main.go`](main.go) file with no logic.

## Using This As A Template

The way to use this as a template is with the following steps:
1. Replace all instances of github.com/probably-not/go-module-small with your module name
2. Replace all instances of go-module-small with your repo name

## Pushing to Docker Hub

If you want to push the Dockerfile to your Docker Hub account, you need to adjust the [`push-docker`](.github/workflows/push-docker) file to have the `.yaml` file extension (it doesn't so as to not trigger this action unless necessary.

Pushing to Docker Hub will also require setting the `DOCKER_HUB_USERNAME` and `DOCKER_HUB_ACCESS_TOKEN` secrets for your actions to use. See [this guide](https://docs.docker.com/ci-cd/github-actions/) for how to generate these values.

In addition to the above steps, you'll need to replace all instances of probablynot/go-module-small with <your dockerhub username>/<your module name>
