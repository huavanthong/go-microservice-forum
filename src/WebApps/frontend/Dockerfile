## We specify the base image we need for our frontend application
FROM node:16-alpine

## We create /frontend directory within our image.
## This directory will hold our frontend source files.
WORKDIR /frontend

## We copy package configuration to ./frontend
COPY package*.json ./frontend

## We copy everything in the root directory 
## into our /frontend direcotry
COPY . /frontend 

## Install npm
RUN npm install

## Because yarn is avalable on npm source, we can ignore yarn
# RUN npm install --global yarn

## Install react-script for nodejs
RUN npm install react-scripts

## Start yarn
CMD yarn run start