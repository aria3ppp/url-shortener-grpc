# url-shortener

`url-shortener` is a backend implementation for a url shortener service written in golang advantage Hexagonal architecture (ports & adaptors) and DDD (domain-driven design).

[![Tests](https://github.com/aria3ppp/url-shortener/actions/workflows/tests.yml/badge.svg)](https://github.com/aria3ppp/url-shortener/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/aria3ppp/url-shortener/badge.svg?branch=master)](https://coveralls.io/github/aria3ppp/url-shortener?branch=master)

### To run the server:

```bash
cp .env.example .env && make server-testdeploy-up
```

Now rest api server is up and running at port `8080` on your `localhost`.
This also run the grpc server at port `8081` on your `localhost`.
