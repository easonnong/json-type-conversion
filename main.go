package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func main() {
	jsonStr := `"host":"http://localhost:8080","port":8080,"analytics_file":"","static_file_version":0,"static_dir":"","templates_dir":"","serTcpSocketHost":"","serTcpSocketPort":8081,"fruits":{}`

	// json str to map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==========json str to map==========")
		fmt.Println(dat)
	}

	// struct to josn str
	var config ConfigStruct
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("==========struct to json str==========")
		fmt.Println(string(b))
	}

	// map to json str
	fmt.Println("==========map to json str==========")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	// array to json str
	arr := []string{"hello", "golang", "python"}
	lang, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("==========array to json str==========")
		fmt.Println(string(lang))
	}

	// json to []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("==========json str to []string==========")
		fmt.Println(wo)
	}
}
