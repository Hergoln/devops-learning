# Go CLI application to deal with things (not fully defined yet)

## Features
* gather statistics regarding usage of reusable workflows and actions within GitHub Actions workflows
  * repositories that script looks into are all repositories accessible using given PAT

## Notes
When running code using container and wants to use a volume, have to pass in absolute path instead of relative path.

### Wrong
This will run a container and create `/go` directory but it will be empty unless `/devops-learning/go` is a valid **ABSOLUTE** path with content \
`docker run -it -v /devops-learning/go:/go golang:1.24.3`


### Correct
Providing **ABSOLUTE** path will feed container proper content \
`docker run -it -v E:/.../go:/go golang:1.24.3`