package try_fbs
import _ "bitbucket.paloaltonetworks.local/tzhao/compare-rtdns/handler/RealtimeDNS"

var testParameterReqList = []*ParameterReq{}
var testParameterReqJSONList = [][]byte{
	[]byte(`{"type":"ParametersReq","trans-id":230,"sn":"007257000062790","protocol":1,"device-data":{"panos-sw":"9.0.1","content":"8138-5378","av":"2936-3446"},"roundtrip-avg":0,"telemetry":"{   \"roundtrip_bucket\" : {     \"50\" : 0,    \"100\" : 0,    \"200\" : 0,    \"400\" : 0,    \"2000\" : 0,    \"timeout\" : 0  }}"}`),
	[]byte(`{"type":"","trans-id":0,"sn":"","protocol":0,"device-data":{"panos-sw":"","content":"","av":""},"roundtrip-avg":0,"telemetry":""}`),
}
var testParameterReqFbsList = [][]byte{

}

var testParameterRespList = []*ParametersResp{}
var testParameterRespJSONList = [][]byte{
	[]byte(`{"type":"","trans-id":0,"protocol":0,"parameter":{"keep-alive-interval":0,"MaxNumEntries":0,"whitelist-url":"","whitelist-update-interval":0}}`),
	[]byte(`{"type":"","trans-id":0,"protocol":0,"parameter":{"keep-alive-interval":0,"MaxNumEntries":0,"whitelist-url":"123455678123245475","whitelist-update-interval":0}}`),
	[]byte(`{"type":"","trans-id":0,"parameter":{"keep-alive-interval":0,"MaxNumEntries":0,"whitelist-url":"123455678123245475","whitelist-update-interval":0},"protocol":0}`),
	[]byte(`{"type":"ParametersResp","trans-id":230,"protocol":1,"parameter":{"keep-alive-interval":1800,"MaxNumEntries":10,"whitelist-url":"/dns/whitelist/whitelist.tar.gz","whitelist-update-interval":86400}}`),
}
var testParameterRespFbsList = [][]byte{

}

var testQueryReqList = []*QueryReq{}
var testQueryReqJSONList = [][]byte{
	[]byte(`{"type":"","trans-id":0,"protocol":0,"query-data":[],"sn":""}`),
	[]byte(`{"type":"","trans-id":0,"protocol":0,"query-data":[{"dst-ip":"","fqdn":"","dns-flag":"","query-type":0}],"sn":""}`),
	[]byte(`{"trans-id":123,"type":"","protocol":123,"query-data":[{"dst-ip":"","fqdn":"","dns-flag":"","query-type":123}],"sn":""}`),
	[]byte(`{"trans-id":123,"protocol":123,"type":"","query-data":[{"dst-ip":"","fqdn":"","dns-flag":"","query-type":123}, {"dst-ip":"","fqdn":"","dns-flag":"","query-type":123}, {"dst-ip":"","fqdn":"","dns-flag":"","query-type":123}],"sn":""}`),
	[]byte(`{"type":"QueryReq","trans-id":116119,"protocol":1,"query-data":[{"dst-ip":"1.0.0.1","fqdn":"c.betrad.com","dns-flag":"0x1","query-type":256}, {"dst-ip":"1.0.0.1","fqdn":"a.betrad.com","dns-flag":"0x1","query-type":256}, {"dst-ip":"1.0.0.1","fqdn":"b.betrad.com","dns-flag":"0x1","query-type":256}, {"dst-ip":"1.0.0.1","fqdn":"c.betrad.com","dns-flag":"0x1","query-type":256}, {"dst-ip":"1.0.0.1","fqdn":"d.betrad.com","dns-flag":"0x1","query-type":256}],"sn":"001801047827"}`),
	[]byte(`{"type":"QueryReq","trans-id":116119,"protocol":1,"query-data":[{"dst-ip":"1.0.0.1","fqdn":"c.betrad.com","dns-flag":"0x1","query-type":256}],"sn":"001801047827"}`),
}
var testQueryReqFbsList = [][]byte{

}

var testQueryRespList = []*QueryResp{}
var testQueryRespJSONList = [][]byte{
	[]byte(`{"type":"","trans-id":0,"protocol":0,"sig-data":[]}`),
	[]byte(`{"type":"","trans-id":0,"protocol":0,"sig-data":[{"fqdn":"","verdict":0,"gtid":0,"ttl":0}]}`),
	[]byte(`{"type":"QueryResp","trans-id":116119,"protocol":1,"sig-data":[{"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}]}`),
	[]byte(`{"type":"QueryResp","trans-id":1234567,"protocol":1,"sig-data":[{"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}, {"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}, {"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}, {"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}, {"fqdn":"c.betrad.com","verdict":0,"gtid":0,"ttl":3600}]}`),
	[]byte(`{"type":"","trans-id":10000000,"protocol":100,"sig-data":[{"fqdn":"123456789012345678901234567890","verdict":111,"gtid":111,"ttl":1111111}, {"fqdn":"1234567890123456789012345678901234567890","verdict":222,"gtid":222,"ttl":2222222}]}`),
}
var testQueryRespFbsList = [][]byte{

}