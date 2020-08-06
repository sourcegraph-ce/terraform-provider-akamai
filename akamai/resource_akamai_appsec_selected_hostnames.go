package akamai

import (
	"fmt"
	"strconv"

	appsec "github.com/akamai/AkamaiOPEN-edgegrid-golang/appsec-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html
func resourceSelectedHostnames() *schema.Resource {
	return &schema.Resource{
		Create: resourceSelectedHostnamesRead,
		Read:   resourceSelectedHostnamesRead,
		Update: resourceSelectedHostnamesUpdate,
		Delete: resourceSelectedHostnamesDelete,
		Schema: map[string]*schema.Schema{
			"configid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"version": &schema.Schema{
				Type:             schema.TypeInt,
				Required:         true,
				DiffSuppressFunc: suppressConfigurationCloneVersion,
			},
			"hostnames": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceSelectedHostnamesRead(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceSelectedHostnamesRead-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Read SelectedHostnames")

	selectedhostnames := appsec.NewSelectedHostnamesResponse()
	configid := d.Get("configid").(int)
	version := d.Get("version").(int)
	err := selectedhostnames.GetSelectedHostnames(configid, version, CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}

	d.SetId(strconv.Itoa(configid))
	return nil
}

func resourceSelectedHostnamesDelete(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceSelectedHostnamesDelete-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Deleting SelectedHostnames")
	return schema.Noop(d, meta)
}

func resourceSelectedHostnamesUpdate(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceSelectedHostnamesUpdate-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Updating SelectedHostnames")
	selectedhostnames := appsec.NewSelectedHostnamesResponse()
	configid := d.Get("configid").(int)
	version := d.Get("version").(int)
	/*
		v := d.Get("hostnames") //.([]interface{})
		//m := make(map[string]string)
		m := make([]appsec.Hostname, len(d.Get("hostnames").([]interface{})))
		//for k, val := range v.(map[string]interface{}) {
		//hn := &appsec.Hostname{}

		for k, val := range v.(*schema.Set).List() {
			hn := &appsec.Hostname{
				Hostname: val.(string),
			}
			m[k] = hn //val.(string)
		}*/
	//selectedhostnames.HostnameList = m
	//selectedhostnames = &appsec.SelectedHostnamesResponse{
	//		HostnameList: convertStringArrToInterface(d.Get("hostnames").([]interface{})),
	//	}

	selectedhostnames.HostnameList = d.Get("hostnames").([]appsec.Hostname)
	//hostnamelist := d.Get("hostnames").([]interface{})
	/*hostnamelist := d.Get("hostnames")

	hostnamearray := make([]*appsec.SelectedHostnamesResponse, len(hostnamelist).([]interface{}))

	codes := []string{}
	for _, code := range hostnamelist.(*schema.Set).List() {
		codes = append(codes, code.(string))
		hn := &appsec.Hostname{
			Hostname: code.(string),
		}
		//selectedhostnames = append(selectedhostnames, codes)
	}

	hna := []*appsec.Hostname{}
	//r.Objects = append(r.Objects, ms)
	selectedhostnames = &appsec.SelectedHostnamesResponse{
		HostnameList: append(hna, hna), //rrdata(d),
	}
	*/
	//hostnameList := hostnamelist.(appsec.SelectedHostnamesResponse)
	//hostnameList := make([]*appsec.SelectedHostnamesResponse, len(hostnamelist)) // create new object list
	//hostnameList := make([]string, len(d.Get("hostnames").([]interface{})))
	//ls := make([]string, len(hostnamelist.([]interface{})))
	//ls := make([]string, len(hostnameList))
	//	for i, sl := range hostnamelist.([]interface{}) {
	//		HostnameList := []appsec.SelectedHostnamesResponse{}
	//		ls[i] = sl.(string)
	//	}
	//	selectedhostnames.HostnameList = hostnameList

	//hostnamelist := d.Get("hostnames").(HostnameList)
	//selectedhostnames.HostnameList = hostnamelist.([]SelectedHostnamesResponse)
	err := selectedhostnames.UpdateSelectedHostnames(configid, version, CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}
	return resourceSelectedHostnamesRead(d, meta)

}

func rrdata(
	d *schema.ResourceData,
) []string {
	rrdatasCount := d.Get("hostnames.#").(int)
	data := make([]string, rrdatasCount)
	for i := 0; i < rrdatasCount; i++ {
		data[i] = d.Get(fmt.Sprintf("hostnames.%d", i)).(string)
	}
	return data
}

func convertStringArrToInterface(strs []string) []interface{} {
	arr := make([]interface{}, len(strs))
	for i, str := range strs {
		arr[i] = str
	}
	return arr
}

/*
func convertStringArr(ifaceArr []interface{}) []string {
	return convertAndMapStringArr(ifaceArr, func(s string) string { return s })
}

func convertStringArrHN(ifaceArr []interface{}) []appsec.Hostname {
	return convertAndMapStringArrHN(ifaceArr, func(s string) string { return s })
}

func convertAndMapStringArr(ifaceArr []interface{}, f func(string) string) []string {
	var arr []string
	for _, v := range ifaceArr {
		if v == nil {
			continue
		}
		arr = append(arr, f(v.(string)))
	}
	return arr
}

func convertAndMapStringArrHN(ifaceArr []interface{}, f func(string) string) []appsec.Hostname { //[]string {
	var arr []appsec.Hostname //[]string
	for _, v := range ifaceArr {
		if v == nil {
			continue
		}
		//arr = append(arr, f(v.(string)))
		arr = append(arr, f(v.(appsec.Hostname)))

	}
	return arr
}
*/
