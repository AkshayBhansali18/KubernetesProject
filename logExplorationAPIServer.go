package main

import (
	//"context"
	//"log"

	//"bytes"
	"context"
	"encoding/json"
	//"log"
    "github.com/gin-gonic/gin"

//"reflect"

	//"time"

	//"context"
	//"io/ioutil"
	//"log"
	"net/http"


	//"context"
	//"context"
	"crypto/tls"
	//"reflect"

	//"crypto/x509"
	"fmt"
	//"io/ioutil"
	//"log"
	//"net/http"
	"github.com/olivere/elastic/v7"
	//"github.com/elastic/go-elasticsearch/v7"
)

//./openshift-install create cluster --dir my_cluster_dir --log-level=debug

func main() {
	cert, err := tls.LoadX509KeyPair("admin-cert", "admin-key")
	if err == nil {
		//fmt.Println(cert)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}


	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true,
				Certificates: []tls.Certificate{cert},},

		},
	}
	esClient, err := elastic.NewClient(
		elastic.SetHttpClient(client),
		elastic.SetURL("https://localhost:9200"),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if(err!=nil){
		fmt.Println(err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {


		ctx := context.Background()
		//termQuery := elastic.NewTermQuery("uuid", "PJWj0Q4ZRpOJuyWkoeW7qQ")
		searchResult, err := esClient.Search().
			Index("infra-000001").
			Pretty(true).
			Do(ctx)
		if err != nil {
			// Handle error
			panic("Get error occurred")
		}

		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			fmt.Println("Index is",hit.Index)
			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			//var t Tweet
			//fmt.Println(hit.Source)
			var mapResp map[string]interface{}
			err := json.Unmarshal(hit.Source, &mapResp)
			if err != nil {
				// Deserialization failed
				fmt.Println(err)
			} else {
				fmt.Println(mapResp,"\n")
				for k, v := range mapResp {
					fmt.Printf("%s : %s\n", k, v)
				}
				//keys := reflect.ValueOf(mapResp).MapKeys()
				//fmt.Println(keys)
				//break
			}
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	//ctx := context.Background()
	//out,_:=exec.Command("oc","-n","openshift-logging","get","pods").Output()
	//output := string(out[:])
	//fmt.Println(output)
	//cert , err:= ioutil.ReadFile("admin-cert")
	//if err!=nil{
	//	fmt.Println(err)
	//
	//}
	//key := /xMnm2AqlGEZ0/zipVt3KA5WrKXpoNZK6CpjXTd\n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\nthq7nGtXkYTAep/yy13gm/Yuxz6cQX0=\n")
	//fmt.Println(key)
	//key, err= ioutil.ReadFile("admin-key")

	//cfg := elasticsearch.Config{
	//	CACert: cert,
	//	//CAKey: key,
	//	Addresses:[]string{"https://localhost:8888"},
	//}
	//ctx := context.Background()

	//esclient, err :=  elasticsearch.NewClient(cfg)
	//if(err!=nil) {
	//	fmt.Println("ES client error ",err)
	//} else{
	//	fmt.Println("ES clients successful")
	//}
	//if(err==nil){
	//	fmt.Println("Error")
	//}
	//fmt.Println(esclient.Info())
	//fmt.Println(esclient)





	//fmt.Println(len(searchResult))


	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	//fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	//fmt.Printf("Found a total of %d logs\n", searchResult.TotalHits())

	//cfg := elasticsearch.Config{
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true,
	//			Certificates: []tls.Certificate{cert},},
	//
	//	},
	//}
	//es, err := elasticsearch.NewClient(cfg)
	//if(err==nil) {
	//	fmt.Println("Yayy No Error")
	//}else{
	//	fmt.Println(es, "yes error :(")
	//}
	//fmt.Println(es)
	////var mapResp map[string]interface{}
	////var buf bytes.Buffer
	//
	//res, _ := es.Search(
	//	es.Search.WithIndex("app", "infra", "audit"),
	//	es.Search.WithTrackTotalHits(true),
	//	es.Search.WithPretty(),
	//)
	//fmt.Println(res)
	//log.Println(es.Cat.Indices())
	//

	//names, err := esClient.IndexNames()
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//for _, name := range names {
	//	fmt.Printf("%s\n", name)
	//}

	//ctx := context.Background()
	//get1, err := esClient.Get().
	//	Index("app-000001").
	//	Type("log").
	//	Id("1").
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	fmt.Println("Error",err)
	//}
	//if get1.Found {
	//	fmt.Printf("Got document in version %d from index %s, type %s\n", get1.Version, get1.Index, get1.Type)
	//}

	//htmlData, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//fmt.Printf("%v\n", resp.Status)
	//fmt.Printf(string(htmlData))

	//newStudent := Student{
	//	Name:         "Gopher doe",
	//	Age:          10,
	//	AverageScore: 99.9,
	//}
	//	elastic.SetSniff(false),
	//	elastic.SetHealthcheck(false))
	//if err == nil {
	//	fmt.Println("No Error initializing : ")
	//	//panic("Client fail ")
	//}

	//htmlData, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//fmt.Printf("%v\n", resp.Status)
	//fmt.Printf(string(htmlData))

}
