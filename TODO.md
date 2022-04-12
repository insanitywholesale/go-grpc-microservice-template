# TODOs
Things left to be done before the project is in an okay state

## Make-type files
For every sufficiently popular language eventually there exists a Makefile replacement
- finish `magefile.go`
- bring `Makefile`, `Taskfile.yml`, `magefile.go` to parity when it comes to naming and functionality

## CI-type files
The more CI, the better
- add `Jenkinsfile`
- add `.travis.yml`
- add whatever Tekton has

## Deployment files
Need a way to run the application
- add `docker-compose.yml` for easier local testing
- add raw kubernetes yaml
- add kustomize yaml to demonstrate testing/prod

## Functionality
Functionality is not yet complete

### Additional version demo
Add later API versions to demonstrate that fields can be added but breaking changes need a full API version change.

### Rething REST api versioning
Might want to actually namespace this but it works somewhat okay

### Repo
Adjust repos to store and retrieve all info for HelloRequest/HelloResponse
