FROM kelseyhightower/kube-cert-manager:0.2.0

LABEL maintainer="Michael Laccetti (michael@laccetti.com)"

ADD dist/dns01-exec-plugin-r53 /route53

ENTRYPOINT ["/kube-cert-manager"]
