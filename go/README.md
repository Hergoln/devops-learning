Testing REST API usage for potential Hacker News implementation of my own.

## Notes
When running code using container and wants to use a volume, have to pass in absolute path instead of relative path.

### Wrong
This will run a container and create `/go` directory but it will be empty unless `/devops-learning/go` is a valid **ABSOLUTE** path with content \
`docker run -it /devops-learning/go:/go -it golang:1.24.3`


### Correct
Providing **ABSOLUTE** path will feed container proper content \
`docker run -it C:/codebase/devops-learning/go:/go -it golang:1.24.3`