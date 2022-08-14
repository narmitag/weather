FROM ubuntu:latest

RUN apt-get update
RUN apt-get -y install git python3 pip mysql