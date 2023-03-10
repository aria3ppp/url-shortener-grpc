# url-shortener-grpc

`url-shortener-grpc` is a grpc server for a url shortener service written in golang advantage Hexagonal architecture (ports & adaptors) and DDD (domain-driven design).

[![Tests](https://github.com/aria3ppp/url-shortener-grpc/actions/workflows/tests.yml/badge.svg)](https://github.com/aria3ppp/url-shortener-grpc/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/aria3ppp/url-shortener-grpc/badge.svg?branch=master)](https://coveralls.io/github/aria3ppp/url-shortener-grpc?branch=master)

### To run servers:

```bash
cp .env.example .env && make server-testdeploy-up
```

Now both rest api server at port `8080` and grpc server at port `8081` are up and running on your `localhost`.
