package softlayer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseurl = "https://api.softlayer.com/rest/v3/SoftLayer_Account/"

func GetNextInvoiceTotalAmount(client *http.Client, user string, key string) string {

	fmt.Println(client)

	req, err := http.NewRequest("GET", baseurl+"getNextInvoiceTotalAmount.json", nil)
	fmt.Println(err)
	req.SetBasicAuth(user, key)
	resp, err := client.Do(req)
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(err)

	return "$" + string(body)
}
