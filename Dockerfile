FROM ubuntu:20.04
RUN apt update && apt upgrade
RUN apt install -y python3 mecab pip
RUN pip install mecab-python3 unidic-lite