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

## Functionality
Functionality is not yet complete

### Additional version demo
Add later API versions to demonstrate that fields can be added but breaking changes need a full API version change.

### API and documentation versioning
If there is v1, v2, v3 of the API, make all versioned endpoints accessible as well as their corresponding docs.
Additionally, make `/api` and `/docs` point to the latest version.
