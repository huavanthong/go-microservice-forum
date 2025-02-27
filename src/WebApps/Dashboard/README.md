# Table of Contents
* [Getting Started](#getting-started)


# Question
* [How do you change the default port to another on ReacJS?](#change-port)

## Table of Contents
- [File Structure](#file-structure)
- [Getting Started](#getting-started)
- [Demo](#demo)
- [Issues knowledge](#issues-knowledge)


### File structure

Within the download you'll find the following directories and files:

```
Muse Ant Design Dashboard
    ├── muse-ant-design-dashboard
    │   ├── public
    │   │   ├── index.html
    │   │   ├── favicon.png
    │   │   ├── minifest.json
    │   │   └── robots.txt
    │   ├── src
    │   │   ├── assets
    │   │   │   ├── images
    │   │   │   └── styles
    │   │   ├── components
    │   │   │   ├── chart
    │   │   │   └── layout
    │   │   ├── pages
    │   │   │   ├── Billing.js
    │   │   │   ├── Home.js
    │   │   │   ├── Profile.js
    │   │   │   ├── Rtl.js
    │   │   │   ├── SignIn.js
    │   │   │   ├── SignUp.js
    │   │   │   └── Tables.js
    │   │   ├── App.js 
    │   │   └── index.js
    │   ├── CHANGELOG.md
    │   ├── LICENSE
    │   ├── package.json
    │   ├── README.md
    │   ├── .env
```


## Getting Started
After clone this project:
```
yarn install
```

Install react-scripts depedency
```
npm install react-scripts
```

Run dashboard
```
yarn start
```

## Running with Docker
Build
```
docker build -t dashboard .
```

Run with docker
```
docker run -p 80:4000 -it dashboard
```
## Demo

- [Dashboard](http://localhost:4000/dashboard)
- [Tables](http://localhost:4000/tables)
- [Billing](http://localhost:4000/billing)
- [RTL](http://localhost:4000/rtl)
- [Profile](http://localhost:4000/profile)
- [Sign In](http://localhost:4000/sign-in)
- [Sign Up](http://localhost:4000/sign-up)

## Issues knowledge
#### Change port
to change port, you can update at .env file
```
PORT=4000
```
More details: [here](https://scriptverse.academy/tutorials/reactjs-change-port-number.html)