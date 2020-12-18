## Commands

There are only two commands
- serve
- migrate

At least one is required to start the application. 
Inspired by the [ory](https://github.com/ory) projects, this project does not automatically migrate
the database and is left up to the users' discretion.

## Environment variables

A Document containing all the environment variables with their purpose and explanation.
Environment variables are prefixed with NOC_

- NOC_PORT=8000
  - Set the server port. The default is 8000.
    
- NOC_HOST=0.0.0.0
  - Set the server binding address. Default binds to host.
    
- NOC_DSN="postgres://username:password@host:port/endpoint"
  - Set the database url. Default is none, which spawns an in memory database.
  
## Configuration file

We use yaml as the default supported configurations, however, one could also use json.

Example yaml configuration file:

    # the database url in the format: provider://username:password@host:port/endpoint?options
    dsn: "postgresql://foo:bar@127.0.0.1:5432/foobar?sslmode=disabled"
    serve:
      host: "0.0.0.0"
      port: 8000
    auth:
      jwt:
      enabled: false
      key: a-very-strong-key
      # timeout is in seconds
      timeout:  500
      # refresh is in seconds
      refresh:  500

### Database connection string

environment setting: `NOC_DSN="..."`

configuration setting: `dsn="..."`

#### Sqlite

**Leaving the dsn variable blank will create an in-memory sql database.**

Create a tmp.db file in the current directory

    sqlite://tmp.db

#### PostgreSQL

Connect to a database with username `foo`, password `bar` and host `127.0.0.1`.
In this case ssl is disabled, but can easily be enabled, see below for examples.

    postgresql://foo:bar@127.0.0.1:5432/foobar?sslmode=disabled


With ssl, server-ca, key and cert

    postgresql://foo:bar@127.0.0.1:5432/foobar?\
    sslmode=disabled\
    &sslrootcert=/path/to/server-ca.pem\
    &sslcert=/path/to/cert.pem
    &sslkey=/path/to/key.pem

#### MySQL

:construction: **In Progress**