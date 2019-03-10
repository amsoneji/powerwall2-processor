network_exists=`docker network inspect powerwall2-test-network`
if [ "$network_exists" == "[]" ]
then 
  # Create powerwall2-test-network within docker
  docker network create --driver bridge --subnet=192.168.0.0/16 powerwall2-test-network
fi
if [ -n "$1" ]
then
  docker stop test-gateway-local
  docker run --rm -p 127.0.0.1:8080:80 --net powerwall2-test-network --ip 192.168.0.5 --name test-gateway-local -d amsoneji/test-gateway-go:$1
else
  docker stop test-gateway-local
  docker run --rm -p 127.0.0.1:8080:80 --net powerwall2-test-network --ip 192.168.0.5 --name test-gateway-local -d amsoneji/test-gateway-go:latest
fi