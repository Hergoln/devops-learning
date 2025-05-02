FROM ubuntu:latest

COPY . .

# "source", used to activate venv is a bash
RUN apt update
RUN apt -y upgrade
RUN apt -y install python3 python-is-python3

# install module with pip
RUN apt -y install python3-pip python3-venv

# activate python venv
RUN python -m venv .venv
# activating venv by modifying PATH ENV variable
ENV PATH=".venv/bin:$PATH"
RUN pip install requests unittest
# RUN pip install -r requirements.txt

ENTRYPOINT [ "/bin/bash", "sf.sh" ]