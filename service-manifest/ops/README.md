```kubectl --namespace=dev create secret docker-registry gcr-json-key \
          --docker-server=https://gcr.io \
          --docker-username=_json_key \
          --docker-password="$(cat ${GCR_JSON})" \
          --docker-email=youremail@example.com
secret "gcr-json-key" created```

```kubectl --namespace=ops patch serviceaccount default -p '{"imagePullSecrets": [{"name": "gcr-json-key"}]}'```
