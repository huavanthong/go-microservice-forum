FROM nginx AS base 
COPY ./ApiGateways/NginxProxy/nginx.conf /etc/nginx/conf.d/default.conf

FROM nginx AS build 

# RUN apt-get update
# RUN apt-get install sudo
# RUN sudo apt-get install nano

COPY ./ApiGateways/NginxProxy/nginx.conf /etc/nginx/conf.d/default.conf
