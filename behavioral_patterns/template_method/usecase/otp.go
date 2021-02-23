package usecase

type IOtp interface {
    GenRandomOTP(int) string
    SaveOTPCache(string)
    GetMessage(string) string
    SendNotification(string) error
    PublishMetric()
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  iOtp.publishMetric()
//  return nil
// }

type Otp struct {
    IOtp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
    otp := o.IOtp.GenRandomOTP(otpLength)
    o.IOtp.SaveOTPCache(otp)
    message := o.IOtp.GetMessage(otp)
    err := o.IOtp.SendNotification(message)
    if err != nil {
        return err
    }
    o.IOtp.PublishMetric()
    return nil
}