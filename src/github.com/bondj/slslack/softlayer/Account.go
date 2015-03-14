package softlayer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseurl = "https://api.softlayer.com/rest/v3/SoftLayer_Account/"

func GetNextInvoiceTotalAmount(client *http.Client, user string, key string) string {

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

type Virtual_Guest struct {
	FullyQualifiedDomainName string
	PrimaryIpAddress         string
}
type Notification_Occurrence_Event struct {
	Subject               string
	SystemTicketId        int
	LastImpactedUserCount int
}

type Notification_Occurence_Event_Response struct {
	Collection []Notification_Occurrence_Event
}

func GetRecentEvents(client *http.Client, user string, key string) string {

	req, err := http.NewRequest("GET", baseurl+"getRecentEvents.json", nil)
	fmt.Println(err)
	req.SetBasicAuth(user, key)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	events := make([]Notification_Occurrence_Event, 0)
	err = decoder.Decode(&events)
	fmt.Println(err)
	for _, element := range events {
		fmt.Println(element.Subject, element.SystemTicketId, element.LastImpactedUserCount)
	}
	return "$"
}
func GetDownVirtualGuests(client *http.Client, user string, key string) []Virtual_Guest {

	req, err := http.NewRequest("GET", baseurl+"getNetworkMonitorDownVirtualGuests.json", nil)
	fmt.Println(err)
	req.SetBasicAuth(user, key)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	events := make([]Virtual_Guest, 0)
	err = decoder.Decode(&events)
	fmt.Println(err)
	for _, element := range events {
		fmt.Println(element.FullyQualifiedDomainName, element.PrimaryIpAddress)
	}
	return events
}
