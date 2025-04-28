FROM ubuntu:latest

COPY . .

RUN apt update
RUN apt -y upgrade
RUN apt -y install python3
RUN apt -y install python-is-python3

# install module globaly without pip
RUN apt -y install python3-requests

# install module with pip
# RUN apt -y install python3-pip
# RUN pip install pandas

ENTRYPOINT [ "/bin/bash", "sf.sh" ]