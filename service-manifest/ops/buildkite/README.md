- create a secret to mount ssh key for github
  `kubectl --namespace=ops create secret generic buildkite-github-ssh --from-file=id_rsa_buildkite_git=${GITHUB_SSH}`
- create a secret for our buildkite token
  `kubectl --namespace=ops create secret generic buildkite-agent-token --from-literal=BUILDKITE_AGENT_TOKEN=${BUILDKITE_SSH}`
- create a configMap from hooks/environment 
  `kubectl --namespace=ops create configmap add-env-key-hook --from-file=ADD_ENV_KEY_HOOK=hooks/environment`