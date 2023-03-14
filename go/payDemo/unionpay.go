package payDemo

import "time"

const (
	version  = "1.0"
	encoding = "utf-8"
	merId    = ""
	frontUrl = ""
	backUrl  = ""
)

func getRequestMap(orderId, amt string) map[string]string {
	requestMap := make(map[string]string)
	requestMap["version"] = version
	requestMap["encoding"] = encoding
	requestMap["signMethod"] = ""
	requestMap["txnType"] = "01"
	requestMap["txnSubType"] = "01"
	requestMap["bizType"] = "000201"
	requestMap["channelType"] = "07"
	requestMap["merId"] = merId
	requestMap["accessType"] = "0"
	requestMap["orderId"] = orderId
	requestMap["txnTime"] = time.Now().Format("20060102150405")
	requestMap["currencyCode"] = "156"
	requestMap["txnAmt"] = amt
	requestMap["frontUrl"] = frontUrl
	requestMap["backUrl"] = backUrl
	return requestMap
}

func main() {

}
