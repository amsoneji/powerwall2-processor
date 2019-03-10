if [ -n "$1" ]
then
  case "$1" in 
    test-gateway)
      app_name="test-gateway"
      image_name="test-gateway-go"
      ;;
    processor)
      app_name="processor"
      image_name="powerwall2-processor-go"
      ;;
    *)
      echo "ERROR: First input must match the directory name of the different applications in this repo"
      exit 1
  esac
  echo "Setting app_name var to "$app_name" and image_name var to "$image_name
  if [ -n "$2" ]
  then
    echo "Starting build"
    docker build --build-arg app_name=${app_name} -t amsoneji/$image_name:$2 .
    docker tag amsoneji/$image_name:$2 amsoneji/$image_name:latest
    if [ -n "$3" ]
    then
      if [ "$3" = "push" ]
      then
        docker push amsoneji/$image_name:$2
        docker push amsoneji/$image_name:latest
      else
        echo "ERROR: Invalid third argument. Only supply 'push' as third argument in case you'd like to automatically push to dockerhub"
        exit 1
      fi
    fi
  else
    echo "ERROR: You must provide tag name as the second input to this script! Exiting without building."
    exit 1
  fi
else
  echo "ERROR: You must provide the app_name as the first input to this script, so that the build script knows what app to build an image for"
  exit 1
fi