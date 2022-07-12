package scan

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/ns3777k/go-shodan/v4/shodan"
	"github.com/spf13/viper"
	"io"
	"os"
	"strconv"
	"strings"
	"test/internal"
	"test/sql"
)

type TargetData struct {
	Protocol string
	Domain   string
	Ip       string
	Port     int
	Banner   string
	BaseURL  string
	Title    string
}

func ShodanGet() {

	Targets := new(TargetData)
	// init shodan client
	ShodanApi := viper.Get("ShodanApi")
	SdApi := internal.Strval(ShodanApi)
	client := shodan.NewClient(nil, SdApi)

	filepath := "D:\\Study\\Programing\\Golang\\test\\query\\shodan.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	db := sql.ConnectMySQL()
	for {
		qy, errf := reader.ReadString('\n')

		qy = strings.Replace(qy, "\n", "", -1) // query target

		q := &shodan.HostQueryOptions{Query: qy}
		if q == nil {
			break
		}
		res, _ := client.GetHostsForQuery(context.Background(), q)

		if res == nil {
			break
		}

		if res.Total != 0 {
			fmt.Printf("==== Query result for \"%s\" ====\n", qy)
			for _, host := range res.Matches {
				if host.SSL != nil && len(host.Hostnames) > 0 {
					c := color.New(color.FgRed)
					c.Println("target using https ===> domain format")
					hostnames := strings.Join(host.Hostnames, "")
					baseurl := "https://" + hostnames + ":" + strconv.Itoa(host.Port)
					Targets.BaseURL = baseurl
					Targets.Banner = host.Banner
					Targets.Title = host.Title
				} else if host.SSL != nil && len(host.Hostnames) == 0 {
					c := color.New(color.FgRed)
					c.Println("target using https ===> ip format")
					baseurl := "https://" + host.IP.String() + ":" + strconv.Itoa(host.Port)
					Targets.BaseURL = baseurl
					Targets.Banner = host.Banner
					Targets.Title = host.Title
				}
				if host.SSL == nil && len(host.Hostnames) > 0 {
					c := color.New(color.FgGreen)
					c.Println("target using http ===> domain format")
					hostnames := strings.Join(host.Hostnames, "")
					baseurl := "http://" + hostnames + ":" + strconv.Itoa(host.Port)
					Targets.BaseURL = baseurl
					Targets.Banner = host.Banner
					Targets.Title = host.Title
				} else if host.SSL == nil && len(host.Hostnames) == 0 {
					c := color.New(color.FgGreen)
					c.Println("target using http ===> ip format")
					baseurl := "http://" + host.IP.String() + ":" + strconv.Itoa(host.Port)
					Targets.BaseURL = baseurl
					Targets.Banner = host.Banner
					Targets.Title = host.Title
				}
				fmt.Println(Targets.BaseURL)
				// Shodan api 暂时获取不到下面的字段, 先注释掉
				//fmt.Println(Targets.Banner)
				//fmt.Println(Targets.Title)
				if len(Targets.BaseURL) > 0 && len(Targets.BaseURL) < 255 {
					sql.InsertDB(db, Targets.BaseURL, Targets.Banner, Targets.Title)
				}
				if errf == io.EOF {
					break
				}
			}
		} else {
			fmt.Println("[-] No results found")
		}
	}
}
