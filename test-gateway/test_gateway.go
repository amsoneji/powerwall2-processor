package main

import (
  "net/http"
  "log"
  "encoding/json"
  "math/rand"
  "time"
)

func main() {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  // Serve up /api/meters/aggregates data
  http.HandleFunc("/api/meters/aggregates", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(Meters_aggregates_response{
      Site:       map[string]interface{}{"last_communication_time":"2018-04-02T16:11:41.885377469-07:00","instant_power":((r1.Float64()*5000)-2500),"instant_reactive_power":-138.8300018310547,"instant_apparent_power":140.47729986545957,"frequency":60.060001373291016,"energy_exported":1136916.6875890202,"energy_imported":3276432.6625890196,"instant_average_voltage":239.81999969482422,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Battery:    map[string]interface{}{"last_communication_time":"2018-04-02T16:11:41.89022247-07:00","instant_power":((r1.Float64()*5000)-2500),"instant_reactive_power":0,"instant_apparent_power":2350,"frequency":60.033,"energy_exported":1169030,"energy_imported":1638140,"instant_average_voltage":239.10000000000002,"instant_total_current":45.8,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Load:       map[string]interface{}{"last_communication_time":"2018-04-02T16:11:41.885377469-07:00","instant_power":(r1.Float64()*2500),"instant_reactive_power":-71.43153973801415,"instant_apparent_power":1547.920305979569,"frequency":60.060001373291016,"energy_exported":0,"energy_imported":7191016.994444443,"instant_average_voltage":239.81999969482422,"instant_total_current":6.44763264839839,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Solar:      map[string]interface{}{"last_communication_time":"2018-04-02T16:11:41.885541803-07:00","instant_power":(r1.Float64()*2500),"instant_reactive_power":53.26999855041504,"instant_apparent_power":3906.533259164868,"frequency":60.060001373291016,"energy_exported":5534272.949724403,"energy_imported":13661.930279959455,"instant_average_voltage":239.8699951171875,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Busway:     map[string]interface{}{"last_communication_time":"0001-01-01T00:00:00Z","instant_power":0,"instant_reactive_power":0,"instant_apparent_power":0,"frequency":0,"energy_exported":0,"energy_imported":0,"instant_average_voltage":0,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Frequency:  map[string]interface{}{"last_communication_time":"0001-01-01T00:00:00Z","instant_power":0,"instant_reactive_power":0,"instant_apparent_power":0,"frequency":0,"energy_exported":0,"energy_imported":0,"instant_average_voltage":0,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0},
      Generator:  map[string]interface{}{"last_communication_time":"0001-01-01T00:00:00Z","instant_power":0,"instant_reactive_power":0,"instant_apparent_power":0,"frequency":0,"energy_exported":0,"energy_imported":0,"instant_average_voltage":0,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0},
    })
  })
  // Serve up /api/meters/site data
  http.HandleFunc("/api/meters/site", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode([1]Meters_site_solar_response{ 
      Meters_site_solar_response{
        Id:               0,
        Location:         "site",
        Type:             "neurio_tcp",
        Cts:              [4]bool{true,true,false,false},
        Inverted:         [4]bool{false,false,false,false},
        Connection:       map[string]interface{}{"ip_address":"Neurio-39546","port":443,"short_id":"39546","device_serial":"OBB3364102752","neurio_connected":true,"https_conf":map[string]interface{}{"client_cert":"/etc/site/certs/neurio/neurio.crt","client_key":"/etc/site/certs/neurio/neurio.key","server_ca_cert":"/etc/site/certs/neurio/neurio-ca-chain.cert.pem","max_idle_conns_per_host":1}},
        Cached_Readings:  map[string]interface{}{"last_communication_time":"2018-06-10T16:51:46.187715089+01:00","instant_power":((r1.Float64()*50)-25),"instant_reactive_power":((r1.Float64()*50)-25),"instant_apparent_power":19.80627466405224,"frequency":49.95000076293945,"energy_exported":3724.253888912031,"energy_imported":26003.843888912033,"instant_average_voltage":247.52999755740166,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0,"v_l1n":247.3300018310547,"v_l2n":0.2199999988079071,"serial_number":"0x000004714B008720","version":"Tesla-0.0.7"},
      },
    })
  })
  // Serve up /api/meters/solar data
  http.HandleFunc("/api/meters/solar", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode([1]Meters_site_solar_response{
      Meters_site_solar_response{
        Id:               0,
        Location:         "solar",
        Type:             "neurio_tcp",
        Cts:              [4]bool{false,false,false,true},
        Inverted:         [4]bool{false,false,false,false},
        Connection:       map[string]interface{}{"ip_address":"Neurio-39546","port":443,"short_id":"39546","device_serial":"OBB3364102752","neurio_connected":true,"https_conf":map[string]interface{}{"client_cert":"/etc/site/certs/neurio/neurio.crt","client_key":"/etc/site/certs/neurio/neurio.key","server_ca_cert":"/etc/site/certs/neurio/neurio-ca-chain.cert.pem","max_idle_conns_per_host":1}},
        Cached_Readings:  map[string]interface{}{"last_communication_time":"2018-06-10T16:52:57.788560639+01:00","instant_power":((r1.Float64()*500)-250),"instant_reactive_power":((r1.Float64()*500)-250),"instant_apparent_power":344.3197561756678,"frequency":49.95000076293945,"energy_exported":3.8174999999938235,"energy_imported":125317.00444444444,"instant_average_voltage":246.82000732421875,"instant_total_current":0,"i_a_current":0,"i_b_current":0,"i_c_current":0,"v_l1n":246.8800048828125,"serial_number":"0x000004714B008720","version":"Tesla-0.0.7"},
      },
    })
  })
  log.Fatal(http.ListenAndServe(":80", nil))
}
