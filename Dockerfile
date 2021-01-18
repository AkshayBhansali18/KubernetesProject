FROM python:3.6.1-alpine
MAINTAINER Akshay Bhansali abhansal@redhat.com
WORKDIR /project
ADD . /project
RUN pip3 install flask
EXPOSE 5100
CMD ["python3","-d","app.py"]


