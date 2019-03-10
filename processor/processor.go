package main

import (
  "os"
  "net/http"
  "net/url"
  "encoding/json"
  "time"
  "github.com/jeremywohl/flatten"
  "github.com/influxdata/influxdb1-client"
  log "github.com/sirupsen/logrus"
)

type json_response_struct struct {
  Struct map[string]interface{}
  List []map[string]interface{}
}

type processor_instruction struct {
  Protocol string
  Host string
  Port string
  Path string
}

type processor_data struct {
  Host string
  Path string
  Timestamp time.Time
  Data map[string]interface{}
}

func prepareJsonObject(instruction_path string) (json_response_struct){
  map_blob, list_blob := get_obj_for_path(instruction_path)
  var body json_response_struct
  if len(list_blob) != 0 {
    body.List = list_blob
  } else if len(map_blob) != 0 {
    body.Struct = map_blob
  }
  return body
}

func jsonRequestWrapper(netClient *http.Client, url string, target *json_response_struct) {
  r, err := netClient.Get(url)
  if err != nil {
      log.Error("Error creating http client: ",err)
  }
  defer r.Body.Close()
  
  if len(target.Struct) != 0 {
    err := json.NewDecoder(r.Body).Decode(&target.Struct)
    if err != nil {
      log.Error("Error when decoding to struct: ",err)
    }
  } else if len(target.List) != 0 {
    err := json.NewDecoder(r.Body).Decode(&target.List)
    if err != nil {
      log.Error("Error when decoding to list: ",err)
    }
  } else {
    err := json.NewDecoder(r.Body).Decode(target)
    if err != nil {
      log.Error("Error when decoding to overall target: ",err)
    }
  }
}

func flattenJson(jsonbody map[string]interface{}) (map[string]interface{}, error){
  flat, err := flatten.Flatten(jsonbody, "", flatten.DotStyle)
  if err != nil{
    return nil, err
  }
  return flat, err
}

func scheduler(gateway_protocol string, gateway_host string, gateway_port string, ticker_channel <-chan time.Time, instruction_feed chan<- processor_instruction){
  for {
    t := <-ticker_channel
    log.Debug("Scheduler triggered at: ",t)
    paths := get_all_paths()
    for _, path := range (paths){
      instruction_feed <- processor_instruction{gateway_protocol,gateway_host, gateway_port,path}
    }   
  }
}

func get_data(instruction_feed <-chan processor_instruction, data_feed chan<- processor_data){
  var netClient = &http.Client{
    Timeout: time.Second * 10,
  }
  var timestamp time.Time
  for {
    instruction := <-instruction_feed
    jsonbody := prepareJsonObject(instruction.Path)
    jsonRequestWrapper(netClient,(instruction.Protocol+"://"+instruction.Host+":"+instruction.Port+instruction.Path),&jsonbody)
    timestamp = time.Now()
    var flat map[string]interface{}
    if len(jsonbody.Struct) != 0 {
      flat, _ = flattenJson(jsonbody.Struct)
    } else if len(jsonbody.List) != 0 {
      flat, _ = flattenJson(jsonbody.List[0])
    }
    if flat != nil {
      log.Debug("data about to be sent to influx: ",processor_data{instruction.Host, instruction.Path, timestamp, flat})
      data_feed <- processor_data{instruction.Host, instruction.Path, timestamp, flat}
    } else {
      log.Debug("Flat was nil!")
    }
  }
}

func save_data_to_influx(influx_protocol string, influx_host string, influx_port string, data_feed <-chan processor_data, success_feed chan<- string){
  // Create Influx Client
  host, err := url.Parse(influx_protocol+"://"+influx_host+":"+influx_port)
  if err != nil {
    log.Error(err)
  } else {
    influxConn, err := client.NewClient(client.Config{
      URL:      *host,
      Username: os.Getenv("influx_username"),
      Password: os.Getenv("influx_password"),
    })
    if err != nil {
        log.Error("Error creating Client: ", err)
    } else {
      influx_database := os.Getenv("influx_database_name")
      for {
        // Grab data from data_feed channel
        data := <-data_feed

        // Instantiate datapoints array
        datapoints := make([]client.Point, 1)

        // Populate datapoints array
        datapoints[0] = client.Point{
          Measurement: data.Path,
          Tags: map[string]string{"host": data.Host},
          Fields: data.Data,
          Time: data.Timestamp,
        }
        
        // Create wrapper BatchPoints object
        batches := client.BatchPoints{
          Points: datapoints,
          Database: influx_database,
          Time: data.Timestamp,
        }

        _, err := influxConn.Write(batches)
        if err != nil {
          log.Error("Error in writing data to Influx: ", err)
        } else {
          success_feed <- ("Measurement saved: " + data.Path + " with tags: host: " + data.Host + " and timestamp: " + data.Timestamp.String())
        }
      }
    }
  }
}

func main() {
  // Get env variables
  gateway_protocol := os.Getenv("gateway_protocol")
  gateway_host := os.Getenv("gateway_host")
  gateway_port := os.Getenv("gateway_port")
  influx_protocol := os.Getenv("influx_protocol")
  influx_host := os.Getenv("influx_host")
  influx_port := os.Getenv("influx_port")

  log.SetLevel(log.InfoLevel)
  // Create Channels needed for go routines
  instruction_feed := make(chan processor_instruction, 10)
  data_feed := make(chan processor_data, 10)
  success_feed := make(chan string, 10)

  // Start go routines
  ticker := time.NewTicker(time.Second * 15)
  defer ticker.Stop()
  go scheduler(gateway_protocol, gateway_host, gateway_port, ticker.C, instruction_feed)
  go get_data(instruction_feed, data_feed)
  go save_data_to_influx(influx_protocol, influx_host, influx_port, data_feed, success_feed)
  for {
    success := <- success_feed
    log.Info("Success in saving the following data to Influx!: ",success)
  }

}
