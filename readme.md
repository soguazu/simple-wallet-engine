# Simple Wallet Engine

Wallet engine simulates wallet transactions.

## Built with:
- [Golang](https://go.dev/dl/)
- [Postgres](https://postgresapp.com) product database
- [Sqlite](https://www.sqlite.org/download.html) product testing
- [Docker](https://www.docker.com/products/docker-desktop/)

## Installation

Use the package manager [go modules](https://go.dev/blog/using-go-modules) to install all dependencies.

NOTE: Docker must be installed before running this application

```bash
git clone https://github.com/soguazu/simple-wallet-engine.git
```

```bash
cd simple-wallet-engine 
```

```bash
go mod download
```

```bash
touch .env
```
Copy the value inside .env.sample into the .env and fill the values for the necessary config


## Usage

```bash
# To build a postgres docker image
make postgres

# To create a postgres database 
make createdb

# To run build swagger docs and run service
make start

# To run both unit and integration test
make test
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
