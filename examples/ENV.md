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

    # the database url in the format: provider://username:password@host:port/endpoint
    dsn: "postgres://foo:bar@127.0.0.1:5432"
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

