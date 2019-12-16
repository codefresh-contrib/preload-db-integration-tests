Preloading data in a database for integration tests with Codefresh
=============

This project is a simple Go application that is using a Postgres database for integration tests.
The project contains tests and also an SQL file to preload the DB with test data


## Local usage

You need to have docker, go and postgres-client installed on your workstation first.

```
$ docker run -p 5432:5432 postgres:11.5
```

Then open another terminal and load the test data after verifying that the database is up:

```
$ pg_isready -h localhost
$ psql -h localhost -U postgres < testdata/preload.sql
```

A Postgres instance is now running at localhost:5432 and you can run the tests with:

```
$ go test -v
```


## To use this project in Codefresh 

1. Create a [free Codefresh account](https://codefresh.io/docs/docs/getting-started/create-a-codefresh-account/)
1. Fork this project in your Github account
1. Create a new pipeline with [codefresh.yml](codefresh.yml) 

More information at the [documentation page](https://codefresh.io/docs/docs/yaml-examples/examples/populate-a-database-with-existing-data/)


Enjoy!





