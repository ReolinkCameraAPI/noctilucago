<p align="center"><img src="https://github.com/ReolinkCameraAPI/noctiluca-go-server/blob/main/.assets/noctiluca.png" width=100px></p>
<h1 align="center">NoctiLuca Server</h1>

<p align="center">
<img alt="Reolink Approval" src="https://img.shields.io/badge/reolink-approved-blue?style=flat-square">
<img alt="GitHub" src="https://img.shields.io/github/license/ReolinkCameraApi/noctiluca-go-server?style=flat-square">
<img alt="Discord" src="https://img.shields.io/discord/773257004911034389?style=flat-square">
</p>

---

A Reolink Camera Management server written in Go built off of
the [reolinkapigo](https://github.com/ReolinkCameraAPI/reolinkapigo). This is the backend to the NoctiLuca frontend.

### Sponsorship

<a href="https://oleaintueri.com"><img src="https://oleaintueri.com/images/oliv.svg" width="60px"/><img width="200px" style="padding-bottom: 10px" src="https://oleaintueri.com/images/oleaintueri.svg"/></a>

[Oleaintueri](https://oleaintueri.com) is sponsoring the development and maintenance of these projects within their
organisation.


---

### This is still untested and heavily in development. :construction:

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

See the example in `examples` folder or go
[here](https://github.com/ReolinkCameraAPI/noctilucago/blob/main/examples/Configuration.md)

When running the server without any configurations passed, the defaults will kick in for development testing. It is
advised to read-up on the configurations to get it working for your use-case.

### API Documentation

This system uses [swagger](https://github.com/go-swagger/go-swagger) to generate the documentation needed from the
source code. For a complete matrix of the API, see below table.

| resource | GET | POST | PUT | DELETE | 
| :------------- | :----------: | :-----------: | :-----------: | :-----------: |
| camera | X | X | X | X |
| model | X | X | - | - |

### Supported Cameras

Any Reolink camera that has a web UI should work. The other's requiring special Reolink clients do not work and is not
supported here.

- RLC-411WS
- RLC-423
- RLC-420-5MP
- RLC-410-5MP
- RLC-520