# route53 DNS-01 Exec Plugin

**Credit**

Pulled from [Matthias Loibl's PR on GitHub](https://github.com/kelseyhightower/dns01-exec-plugins/pull/1)

### Building

```
docker build -t mlaccetti/kube-cert-manager-r53:build -f Dockerfile.build .
docker run -it --rm -v $(pwd):/usr/src/dns01-exec-plugin-r53 mlaccetti/kube-cert-manager-r53:build

docker build -t mlaccetti/kube-cert-manager-r53:0.5.0 .
```

### Usage

##### Configuration

The `route53` plugin looks up the zone ID in Route53 directly, no need to add any additional environment variables.

We re-use the `dns01.json` approach, mapping it as follows:
```
{
  "username": "AWS_ACCESS_ACCESS_KEY_ID",
  "token": "AWS_SECRET_ACCESS_KEY"
}
```

##### Creating DNS-01 TXT Records

```
cat dns01.json | \
  APIVERSION="v1" \
  COMMAND="CREATE" \
  DOMAIN="example.com" \
  FQDN="_acme-challenge.example.com." \
  TOKEN="8bGFl9SNhZzukcwdR7e52gFwq6HaEHB43LbimZQwnLg" \
  route53
```

```
echo $?
```
```
0
```

##### Deleting DNS-01 TXT Records

```
cat dns01.json | \
  APIVERSION="v1" \
  COMMAND="DELETE" \
  DOMAIN="example.com" \
  FQDN="_acme-challenge.example.com." \
  TOKEN="8bGFl9SNhZzukcwdR7e52gFwq6HaEHB43LbimZQwnLg" \
  route53
```

```
echo $?
```
```
0
```

### API Version Conflict

```
cat dns01.json | \
  APIVERSION="v2" \
  COMMAND="DELETE" \
  DOMAIN="example.com" \
  FQDN="_acme-challenge.example.com" \
  TOKEN="8bGFl9SNhZzukcwdR7e52gFwq6HaEHB43LbimZQwnLg" \
  route53
```

```
echo $?
```
```
3
```

##### Bad Configuration Data

```
cat dns01-bad.json | \
  APIVERSION="v1" \
  COMMAND="DELETE" \
  DOMAIN="example.com" \
  FQDN="_acme-challenge.example.com" \
  TOKEN="8bGFl9SNhZzukcwdR7e52gFwq6HaEHB43LbimZQwnLg" \
  route53
```

A single error message is printed to stderr.

```
invalid character 'B' looking for beginning of value
```

```
echo $?
```
```
2
```
