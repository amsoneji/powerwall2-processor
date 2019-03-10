network_exists=`docker network inspect powerwall2-test-network`
if [ "$network_exists" == "[]" ]
then 
  # Create powerwall2-test-network within docker
  docker network create --driver bridge --subnet=192.168.0.0/16 powerwall2-test-network
fi
docker stop powerwall2-processor
if [ -n "$1" ]
then
  if [ "$1" = "test" ]
  then
    if [ -n "$2" ]
    then
      docker run --rm --net powerwall2-test-network \
      --ip 192.168.0.4 \
      --name powerwall2-processor \
      --env gateway_protocol \
      --env gateway_host \
      --env gateway_port \
      --env influx_protocol \
      --env influx_host \
      --env influx_port \
      --env influx_database_name \
      --env influx_username \
      --env influx_password \
      amsoneji/powerwall2-processor-go:$2 /bin/bash -c "go test -v"
    else
      docker run --rm --net powerwall2-test-network \
      --ip 192.168.0.4 \
      --name powerwall2-processor \
      --env gateway_protocol \
      --env gateway_host \
      --env gateway_port \
      --env influx_protocol \
      --env influx_host \
      --env influx_port \
      --env influx_database_name \
      --env influx_username \
      --env influx_password \
      amsoneji/powerwall2-processor-go:latest /bin/bash -c "go test -v"
    fi
  else
    docker run --rm --net powerwall2-test-network \
      --ip 192.168.0.4 \
      --name powerwall2-processor \
      --env gateway_protocol \
      --env gateway_host \
      --env gateway_port \
      --env influx_protocol \
      --env influx_host \
      --env influx_port \
      --env influx_database_name \
      --env influx_username \
      --env influx_password \
      -d amsoneji/powerwall2-processor-go:$1
  fi
else
  docker run --rm --net powerwall2-test-network \
    --ip 192.168.0.4 \
    --name powerwall2-processor \
    --env gateway_protocol \
    --env gateway_host \
    --env gateway_port \
    --env influx_protocol \
    --env influx_host \
    --env influx_port \
    --env influx_database_name \
    --env influx_username \
    --env influx_password \
    -d amsoneji/powerwall2-processor-go:latest
fi