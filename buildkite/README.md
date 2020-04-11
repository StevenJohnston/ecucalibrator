- create a secret to mount ssh key for github
  `kubectl create secret generic buildkite-github-ssh --from-file=id_rsa_buildkite_git=${GITHUB_SSH}`
- create a configMap from hooks/environment 
  `kubectl create configmap add-env-key-hook --from-file=ADD_ENV_KEY_HOOK=hooks/environment`