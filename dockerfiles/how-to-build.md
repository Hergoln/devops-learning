# Docker building cheat sheet
As I am using docker very rarely, I added this cheat sheet on instructions to build dockerfiles

## Build image with path to file and tag
```sh
docker build -t sf-t -f dockerfiles/sf.dockerfile dockerfiles
```
> -t : tag \
> -f : filepath \
> argument without flag "dockerfiles" provides a context for -f to be able to find file given as -f

## Run image in a container in interactive mode
```sh
docker run -it sf-t
```
> -it : interactive


## Build python image with requirements file in different location then dockerfile
```sh
```