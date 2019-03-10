# powerwall2-processor
Processor framework to ingest and distribute data from a Tesla Powerwall 2

## Current features and implementation details

This implementation is heavily dependent on the data format and access defined in Vince Loschiavo's Tesla Powerwall 2 API [here](https://github.com/vloschiavo/powerwall2)

As of this writing, the processor is built to ingest from this API and send the data to influxdb. As such, a test harness providing an API stub, as well as containers for influxdb and grafana, have been provided to aid in local testing and development.

As this was developed as a strictly experimental/technical exercise, a decision was made to also use Go as the language for the processor. Not only did it serve as an opportunity to become more familiar with the language, but also test out its multiprocessing and concurrency programming models. Considering this processor is meant to connect to a battery system's API, having a strong base for the processor framework that allows for concurrent processing is critical.

## Getting started

This `Getting Started` tutorial assumes you already have certain basic commandline tools such as `curl` installed. It also expects a Docker daemon to be running as well. For more details on Docker and its installation, please visit [Docker's documentation here](https://docs.docker.com/install/).

### Running in local mode

To run the entire system in local mode for testing and development, please first export the necessary environment variables by running:
`source local_setup.sh`

Next, let's start influxdb. Please run `sh test-influx/run_influx_local.sh`. This will start influxdb as an ephemeral container on your system, which means that between container stops or restarts, all your data will disappear. Please use a separate influxdb installation in case you'd like to test persistence.
* Once you've run the above script, you can test that the admin user and database have been successfully created by running 
`curl -u $influx_username:$influx_password -G 'http://127.0.0.1:8086/query?db='$influx_database_name --data-urlencode 'q=SHOW DATABASES;'`
* You should get a response that looks like:
`{"results":[{"statement_id":0,"series":[{"name":"databases","columns":["name"],"values":[["testdb"],["_internal"]]}]}]}`

After influx is up and running, you can also choose to install a grafana instance, to make debugging locally painless and quick. This can be done by running `sh test-grafana/run_grafana_local.sh`

Also, in order to test out the processor's ability to fetch data appropriately, please start the test-gateway (also written in Go) by running `sh test-gateway/run_gateway_local.sh`. Please note, running that script without arguments will automatically pull and run the 'latest' tag of the image. If you'd like to run a different image tag, you can run `sh test-gateway/run_gateway_local.sh $TAG`.

Now that all local dependencies have been resolved and started, please feel free to run the processor itself by running `sh processor/run_processor.sh`. The moment this process finishes starting, you should see data start flowing into influxdb. Again, please note - running that script without arguments will automatically pull and run the 'latest' tag of the image. If you'd like to run a different image tag, you can run `sh processor/run_processor.sh $TAG`.

### Running in non-local mode

In case you'd like to have the processor run and connect to other services (say, if you already have influx running somewhere, or if you have access to an actual Tesla gateway), please instead adjust the values in setup.sh to reflect the IPs you'd like to use. Then, you can enable those values by running `source setup.sh` and start the processor using the same command as earlier: `sh processor/run_processor.sh`

## Building the application

If you are making adjustments to the code of either the test-gateway, or the processor itself, make sure to finish your changes by building the appropriate image and saving that image to dockerhub. Currently, these images are stored on amsoneji's personal dockerhub account, so you may need to be added as a collaborator before you have the correct permissions to do so. Please reach out in case you have any further questions.

Building the applications into images is easy:
* To build the test-gateway into a new image, please run `sh build_container.sh test-gateway`
* To build the processor into a new image, please run `sh build_container.sh processor`

## Running unit tests

Once you have built the application, you can run the tests that are inbuilt to the application using the same `run_processor.sh` script as before, just with the `test` modifier, as follows:
* To run the tests for the `:latest` image tag, you can run `sh processor/run_processor.sh test`
* To run the tests for a particular image tag, you can run `sh processor/run_processor.sh test $TAG` where `$TAG` is the image tag you want to run for.

## Further Development

Currently, the processor is quite tied to the vloschavio API, as well as InfluxDB. Long term, we do envision the application using a more standardized API from Tesla themselves. Since this is a PoC, we also don't have significant testing or benchmarking capabilities/metrics at the moment. These will be added at a later time since the code has been tested in integration using the influx, grafana and test-gateway code we have here.

If you do have improvements or capabilities in mind, please feel free to reach out or fork this repo to create a PR.

## Thanks!

Feel free to reach out to Aniket Soneji over github/linkedin with any questions!
