FROM kelseyhightower/kube-cert-manager:0.5.0

LABEL maintainer="Michael Laccetti (michael@laccetti.com)"

ADD dist/route33 /route53

ENTRYPOINT ["/kube-cert-manager"]
