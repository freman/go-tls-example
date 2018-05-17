# SSL odds and ends

## makecerts.go

Every so often I need a cert and key to test, or more recently I needed a set of certs and keys.
I can never be bothered setting up or maintaining even an easyca so I knocked this up.

```bash
$ go run makecerts.go
2018/05/18 09:36:23 written ssl/ca.cert
2018/05/18 09:36:23 written ssl/ca.key
2018/05/18 09:36:24 written ssl/bob.crt
2018/05/18 09:36:24 written ssl/bob.key
2018/05/18 09:36:24 written ssl/alice.crt
2018/05/18 09:36:24 written ssl/alice.key
```

This will generate a ca, and two certificates that can be used on localhost.

## mutualauth.go

An example of mutual auth, once you've run makecerts.go you can spin this up and fire curl at it to test things.

```bash
$ go run mutualauth.go
$ curl --key ssl/alice.key --cert ssl/alice.crt --cacert ssl/ca.crt https://localhost:8443/
Issued by Testing CA
Issued to alice cert
```