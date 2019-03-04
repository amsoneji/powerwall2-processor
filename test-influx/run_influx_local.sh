network_exists=`docker network inspect powerwall2-test-network`
if [ "$network_exists" == "[]" ]
then 
  # Create powerwall2-test-network within docker
  docker network create --driver bridge --subnet=192.168.0.0/16 powerwall2-test-network
fi
docker stop test-influx-local
docker run --rm -p 8086:8086 --net powerwall2-test-network --ip 192.168.0.2 --name test-influx-local -e "INFLUXDB_DB="$influx_database_name -e "INFLUXDB_HTTP_AUTH_ENABLED=true" -e "INFLUXDB_ADMIN_USER="$influx_username -e "INFLUXDB_ADMIN_PASSWORD="$influx_password -d influxdb
