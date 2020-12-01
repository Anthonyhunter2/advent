package main

import (
	"fmt"
	"time"

	"codebox.galileosuite.com/gdi/galileo-go/galileo/client"
	_ "codebox.galileosuite.com/gdi/galileo-go/galileo/model"
)

type Result struct {
	Name    string
	RString string
	Error   error
}

type Result2 struct {
	Name    string
	RString string
	Error   error
}

func thing1(incoming int, c chan Result) {
	fmt.Println("Sleeping for 1 seconds")
	time.Sleep(3 * time.Second)
	thisJawn := Result{}
	if incoming == 1 || incoming == 2 {
		successString := fmt.Sprintf("Found %v", incoming)
		thisJawn.Name = "Result"
		thisJawn.RString = successString
		thisJawn.Error = nil
		c <- thisJawn
		return
	}
	thisJawn.Error = fmt.Errorf("Got an error Jay %v", incoming)
	c <- thisJawn
	return
}

func thing2(incoming int, c chan Result2) {
	thisJawn := Result2{}
	fmt.Println("Sleeping for 1 seconds")
	time.Sleep(2 * time.Second)
	if incoming == 3 || incoming == 5 {
		successString := fmt.Sprintf("Found %v", incoming)
		thisJawn.Name = "Result2"
		thisJawn.RString = successString
		thisJawn.Error = nil
		c <- thisJawn
		return
	}
	thisJawn.Error = fmt.Errorf("Got an error Jay %v", incoming)
	c <- thisJawn
	return
}
func main_old() {
	t := time.Now()
	resultChannel := make(chan Result, 1)
	results2 := make(chan Result2, 1)

	go thing1(1, resultChannel)
	go thing2(3, results2)
	i := <-resultChannel
	x := <-results2
	//for i := range resultChannel {
	//	fmt.Println(reflect.TypeOf(i))
	//}
	if i.Error == nil {
		if i.Name == "Result" {
			fmt.Println("Im a Result!")
		} else {
			fmt.Println("Im not a result :(")
		}
		fmt.Println("Success", i.RString)
	}
	if x.Error == nil {
		fmt.Println("Oh No Error:", x.Error)
	}
	fmt.Println(time.Since(t))

}
func main() {

	api, err := client.New(&client.Config{
		Username: "amazzatenta",
		Password: "BrightFuture1#!",
		URL:      "https://my.galileosuite.com",
	})
	if err != nil {
		fmt.Println(err)
	}

	api.SetSite("atsgroup")
	if err := api.GetAssets(); err != nil {
		fmt.Println(err)
	}

	host, err := api.GetHost("gvicaixtsm01")
	//host, err := api.GetHost("g01plapp04")
	if err != nil {
		fmt.Println(err)
	}

	memStats, err := api.GetMetrics(&client.Filter{
		Asset:    host,
		Metric:   "HostVirtualMemory",
		PastDays: 30,
		Summary:  300,
	})
	for _, v := range memStats {
		fmt.Println(v.Label)
	}
	//result, err := api.GetFSConfig(host)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for i, v := range memStats {
	//	fmt.Println("I:", i, "Value: ", v[1])
	//}
	//fmt.Println(result.Row[0][3])

}
