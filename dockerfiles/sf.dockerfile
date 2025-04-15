FROM ubuntu:latest

RUN apt update
RUN apt -y upgrade
RUN apt -y install python3

CMD ["python3", "--version"]