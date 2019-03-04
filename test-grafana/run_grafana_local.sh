network_exists=`docker network inspect powerwall2-test-network`
if [ "$network_exists" == "[]" ]
then 
  # Create powerwall2-test-network within docker
  docker network create --driver bridge --subnet=192.168.0.0/16 powerwall2-test-network
fi
docker stop test-grafana-local
docker run --rm -p 3000:3000 --net powerwall2-test-network --ip 192.168.0.3 --name test-grafana-local -d grafana/grafana