package stats

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/go-prtg-sensor-api"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// stak Api data structs
type errorLog struct {
	Count    int    `json:"count"`
	LastSeen int    `json:"last_seen"`
	Text     string `json:"text"`
}
type hashrate struct {
	Threads [][]float64 `json:"threads"`
	Total   []float64   `json:"total"`
	Highest float64     `json:"highest"`
}
type results struct {
	DiffCurrent int        `json:"diff_current"`
	SharesGood  int        `json:"shares_good"`
	SharesTotal int        `json:"shares_total"`
	AvgTime     float64    `json:"avg_time"`
	HashesTotal int        `json:"hashes_total"`
	Best        []int      `json:"best"`
	ErrorLog    []errorLog `json:"error_log"`
}
type connection struct {
	Pool     string        `json:"pool"`
	Uptime   int           `json:"uptime"`
	Ping     int           `json:"ping"`
	ErrorLog []interface{} `json:"error_log"`
}
type stats struct {
	Version    string `json:"version"`
	hashrate   `json:"hashrate"`
	results    `json:"results"`
	connection `json:"connection"`
}

//
//
//// Map returns a map version of stats data for metrics.go
//// non concurrent usage
//func (stats *stats) Map() map[string]interface{} {
//	if len(stats.Total) == 0 {
//		stats.Total = []float64{0}
//	}
//	m := map[string]interface{}{
//		"DiffCurrent": stats.DiffCurrent,
//		"SharesGood":  stats.SharesGood,
//		"SharesTotal": stats.SharesTotal,
//		"AvgTime":     stats.AvgTime,
//		"HashesTotal": stats.HashesTotal,
//		"Pool":        stats.Pool,
//		"Uptime":      stats.Uptime,
//		"Ping":        stats.Ping,
//		"TotalHR":     stats.Total[0],
//	}
//
//	for k, v := range stats.Threads {
//		m[fmt.Sprintf("Thread.%v", k)] = v[0]
//	}
//	return m
//}
//
//// Map returns a map version of stats data for metrics.go
//// non concurrent usage
//func (stats *stats) threadMapSlice() []map[string]interface{} {
//	ms := make([]map[string]interface{}, 0, 10)
//	for k, v := range stats.Threads {
//		ms = append(ms, map[string]interface{}{"hashrate": v[0], "thread": k})
//
//	}
//	return ms
//}

func getStats(timeout time.Duration, url string) (stats, error) {
	client := http.Client{
		Timeout: timeout,
	}
	out := stats{}
	res, err := client.Get(url)
	if err != nil || res.StatusCode != 200 {
		return stats{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return stats{}, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return stats{}, err
	}
	err = res.Body.Close()
	if err != nil {
		log.Fatalf("failed to close body, stopping under the no leak policy %v", err)
	}

	return out, nil
}

func PrtgStats(timeout time.Duration, url string) {
	show := new(int)
	*show = 1
	// Create empty response and log start time
	r := &prtg.SensorResponse{}
	start := time.Now()
	s, err := getStats(timeout, url)
	if err != nil {
		r.SensorResult.Error = 1
		r.SensorResult.Text = fmt.Sprintf("error connecting to %v using a timeout of %v", url, timeout)
		fmt.Println(r.String())
		return
	}

	responseTime := time.Since(start)

	r.AddChannel(prtg.SensorChannel{
		Name:      "Hash Rate",
		ChannelID: show,
		Value:     s.Total[0],
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitCount,
	})
	// Response time channel
	r.AddChannel(prtg.SensorChannel{
		Name:      "Response time",
		Value:     responseTime.Seconds() * 1000,
		Unit:      prtg.UnitTimeResponse,
		Float:     1,
		ShowChart: show,
		ShowTable: show,
	})

	r.AddChannel(prtg.SensorChannel{
		Name:       "Pool",
		Value:      1,
		Unit:       prtg.UnitCustom,
		CustomUnit: s.connection.Pool,
		ShowChart:  show,
		ShowTable:  show,
	})

	r.AddChannel(prtg.SensorChannel{
		Name:      "Ping Pool",
		Value:     float64(s.connection.Ping),
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeResponse,
	})

	r.AddChannel(prtg.SensorChannel{
		Name:      "Connection Uptime",
		Value:     float64(s.connection.Uptime),
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeSeconds,
	})

	r.AddChannel(prtg.SensorChannel{
		Name:      "Share Time",
		Value:     s.AvgTime,
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeSeconds,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "SharesTotal",
		Value:     float64(s.SharesTotal),
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitCount,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "Good Shares",
		Value:     float64(s.SharesGood),
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitCustom,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "Bad Shares",
		Value:     float64(s.SharesTotal - s.SharesGood),
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitCustom,
	})

	r.AddChannel(prtg.SensorChannel{
		Name:      "Error Count",
		Value:     float64(len(s.results.ErrorLog)),
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitCount,
	})

	fmt.Println(r.String())
}