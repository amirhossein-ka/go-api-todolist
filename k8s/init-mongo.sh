# actual script in app.yaml configmap

set -euo pipefail
MY_POD=$(hostname)
# shellcheck disable=SC2269
RS_NAME=${REPLICA_SET_NAME}
echo "Starting init container on pod: ${MY_POD}"

# Function to check DNS resolution of a pod hostname.
wait_for_dns() {
  local pod_host=$1
  echo "Waiting for DNS resolution of ${pod_host} ..."
  until getent hosts "${pod_host}" >/dev/null 2>&1; do
    sleep 2
  done
  echo "${pod_host} is resolvable."
}

# If this is the first pod, perform the replica set initialization.
if [[ "${MY_POD}" == "mongo-0" ]]; then
  echo "This is mongo-0. Waiting for all replica set members to be resolvable..."
  for i in $(seq 0 $((EXPECTED_MEMBERS - 1))); do
    POD_HOST="mongo-${i}.${SERVICE_NAME}.${NAMESPACE}.svc.cluster.local"
    wait_for_dns "${POD_HOST}"
  done

  # Give pods a bit more time to finish startup.
  echo "Waiting 10 additional seconds for all pods to be fully ready..."
  sleep 10

  echo "Initiating the replica set..."
  mongo --eval "rs.initiate({
		_id: '${RS_NAME}',
		members: [
		    { _id: 0, host: 'mongo-0.${SERVICE_NAME}.${NAMESPACE}.svc.cluster.local:27017' },
			  { _id: 1, host: 'mongo-1.${SERVICE_NAME}.${NAMESPACE}.svc.cluster.local:27017' },
      ]
    })" || {
    echo "Replica set initiation failed. It may already be initialized."
  }
else
  # For all other pods, wait until the replica set is initiated.
  echo "This is ${MY_POD}. Waiting for the replica set to be initiated by mongo-0..."
  RETRIES=30
  until mongo --eval "rs.status().ok" --quiet | grep -q 1; do
    RETRIES=$((RETRIES - 1))
    if [ $RETRIES -le 0 ]; then
      echo "Timeout waiting for replica set initialization."
      exit 1
    fi
    sleep 2
  done
  echo "Replica set is active. ${MY_POD} will now proceed."
fi

echo "Init container tasks completed on ${MY_POD}."
echo "container sleeping..."
sleep inf