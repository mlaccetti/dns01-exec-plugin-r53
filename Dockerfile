FROM kelseyhightower/kube-cert-manager:0.2.0

LABEL maintainer="Michael Laccetti (michael@laccetti.com)"

ADD route53 /route53

ENTRYPOINT ["/kube-cert-manager"]
