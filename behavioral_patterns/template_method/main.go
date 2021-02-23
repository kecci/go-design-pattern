package main

import (
	"fmt"

	"github.com/kecci/design-pattern-go/behavioral_patterns/template_method/usecase"
)

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &usecase.Sms{}
	o := usecase.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &usecase.Email{}
	o = usecase.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
