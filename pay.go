package bcGolang

import (
	"fmt"
	"reflect"
	"math"
	"time"
)

type BCPay struct {
	bcApp 	BCApp
	//timeout	int
}

func (this *BCPay) RegisterApp(bcApp BCApp) {
	this.bcApp = bcApp
}

func (this *BCPay) Pay(payParam BCPayReqParams) BCResult {
	AttachAppSign(&payParam.BCReqParams, PAY, this.bcApp)
	fmt.Println(payParam.BCReqParams)


}

// private methods
func (this *BCPay) getBillPayUrl() string {
	if this.bcApp.IsTestMode {
		return GetRandomHost() + "rest/sandbox/bill"
	} else {
		return GetRandomHost() + "rest/bill"
	}
}
