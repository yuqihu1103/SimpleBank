# Simple Bank
**Author:** [Yuqi Hu](https://yuqihu1103.github.io/)

**Project Description:**

The Simple Bank is a backend web service, which provide APIs to:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
   
2. Record all balance changes to each of the account. So that an account entry record will be created every time some money is added to or subtracted from the account.
   
3. Perform a money transfer between two accounts. This happens within a transaction: either both accounts’ balance are updated successfully or none of them are.

## Demo

## Getting Started For Local Devlopment
### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop) - [TablePlus](https://tableplus.com/) - [Golang](https://golang.org/) - [Homebrew](https://brew.sh/)

- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### Setup infrastructure

- Create the bank-network

    ``` bash
    make network
    ```

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```
    
### Generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

### Run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```

## Deploy to kubernetes cluster

- [Install nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/#aws):

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml
    ```

- [Install cert-manager](https://cert-manager.io/docs/installation/kubernetes/):

    ```bash
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
    ```

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

The development of this project was inspired the [Backend Master Class](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/) on Udemy.
