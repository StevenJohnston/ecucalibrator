- create a secret to mount ssh key for github
  `kubectl --namespace=ops create secret generic buildkite-github-ssh --from-file=id_rsa_buildkite_git=${GITHUB_SSH}`
- create a secret for our buildkite token
  `kubectl --namespace=ops create secret generic buildkite-agent-token --from-literal=BUILDKITE_AGENT_TOKEN=${BUILDKITE_SSH}`
- create a configMap from hooks/environment 
  `kubectl --namespace=ops create configmap add-env-key-hook --from-file=ADD_ENV_KEY_HOOK=hooks/environment`
- create a secret for buildkite to push images to gcr
  <!-- `kubectl --namespace=ops create secret docker-registry gcr-pusher-key --docker-server=https://gcr.io --docker-username=_json_key --docker-password="$(cat ${GCR_PUSHER_JSON})" --docker-email=steven@stevenjohnston.ca` -->
  `kubectl --namespace=ops create secret generic gcr-push-service-account --from-file=docker.json=${GCR_PUSH_SERVICE_ACCOUNT}`
