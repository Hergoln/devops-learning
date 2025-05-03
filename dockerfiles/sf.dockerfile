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
RUN pip install requests
# RUN pip install -r requirements.txt

# install npm
RUN apt -y install nodejs
RUN apt -y install npm
# update npm to latest version
# after @ you can specify tag/version to install
# RUN npm install -g npm@latest

RUN npm install @salesforce/cli --global

ENTRYPOINT [ "/bin/bash", "sf.sh" ]