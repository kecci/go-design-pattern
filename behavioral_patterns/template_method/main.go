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
	err := o.GenAndSendOTP(4)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("")
	emailOTP := &usecase.Email{}
	o = usecase.Otp{
		IOtp: emailOTP,
	}
	err = o.GenAndSendOTP(4)
	if err != nil {
		fmt.Println(err.Error())
	}

}
