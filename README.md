<p align="center"><img src="https://github.com/ReolinkCameraAPI/noctiluca-go-server/blob/main/.assets/noctiluca.png" width=100px></p>
<h1 align="center">NoctiLuca Server</h1>

A Reolink Camera Management server written in Go built off of 
the [reolink-go-api](https://github.com/ReolinkCameraApi/reolink-go-api).
This is the backend to the NoctiLuca frontend.



### This is still untested and heavily in development.

### Join us on Discord

    https://discord.gg/8z3fdAmZJP
    
### Get started

#### Within Docker

    docker build . -t noctiluca
    // detatch with -d option
    docker run -p 8000:8000 noctiluca serve
    
#### From Binary

    tba
    
#### From source

    git clone git@github.com:ReolinkCameraAPI/noctiluca-go-server.git
    cd noctiluca-go-server
    go run . serve
    
    // Or build and run binary
    go build -o bin/noctiluca
    ./noctiluca serve
    
### Configuration

    Setup a yaml config file or use environment variables.
    Copy the default configs from conf/noctiluca.yaml
    
    Environment Variables start with NOC_<key>
    e.g.
    NOC_HOST=0.0.0.0
    NOC_PORT=8000
    NOC_DNS=""