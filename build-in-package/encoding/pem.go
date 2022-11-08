package encoding

import (
	"strings"
)

const (
	AppPrivateKeyZfb = `MIIEpQIBAAKCAQEAohvxiThYobThyu3kZbqufT2hJO3xKNZfPDMU+CHylph0U6dSvOI2UDUN1CbbWtyMI0n+uRmjzRehRhQJmqvB6SKFTjiyCIlzPOyzgepQX81RfyhpL360pBeJPHKfipEpU2iCSmq3mcBJf0VILfiHalPCyzu5fh2nL+nlBrN34mPtH8c84RjpzMFxvYZr9oS9eOYGfC32xxfwNolaReBAjsHrApyqgRSQB7thMNO/7FRo1IqQl6fvOsmMFo/bSuZ8ZZdjuAp7iQ3hk/hVZJXXn8ZwT/6tPfS/k3tTwWa/MucWuz6IWbp3Z3DQFvaf1sGHo4j9MqlTefeonKN/IgdT3wIDAQABAoIBAQCZ5P0/z4YUvas9Auo4ySitLLy2Wkz+/8QZWkUl6tRpswF8CPS66+Wz7ynZpanIcGd5eN2gy359TVmFNbCIbVYBBQyYXapmAFauF+eyBceXq0sa9DonJJpIzS8ze32KBpS84hIOo5D5HVqowBTtz2p+vlXGqkRGPdBh5CHQYsAE4mUg+EMdpPdSqODX5Dhmsgw3aNhfhG+v07XQC7aQIlHttI8jeF4rdCUXfQZ9w9UPSNlOJ+DZnJ6y28L3AYYs5sjZT0pm5z4qZoNjERNRl1i/4agw3zNP3e06kd+veAaHVVvRc0d77i8Ob/WS5GzS05yQ/jEr4WkeAMxrLSc8bwqBAoGBANU2fmT7PXfZcw3oRqbpVc/mE2jwokQu6GYZ8w/HSdtYh+P1xVUkU+PzVkqHZKfonStTSKXwi9naXprr/qWkQWCI0gZgwFzGNFqERjN7FPZdRfcCePGwDI6JD+7tr4/cJztIuCfshdjuFLiM0l3HKfqfEA+frL1kpBdEiSaINKSXAoGBAMKkEdI+orUrbgrSGeFhrDNJdiacxlO9y+hfRZUzmPzDpZUOdMV/LMVxtR6VYKoOllH71nHfPpzeFkELvOcduGhlNresmul198m2gPXTE1p/lkPVksAzFRH25zAFjdDVy5iH5GKrEThpGj2Dx8LFjv6ptQ89haJuCAt8JCVGnUv5AoGBALm5vmA4elLsGE0FCIZcu2NeB4piEvdR5R8Le69C6hMoCzeH311LR8hJL+G8DvI0rrQO2Dm0UB51GfPZnvirHEf65vinTumBvhkbIAu4K4pvtYJ1pOTjdgyzWC3I3iVyoLoDnmcooW9V9LbN9HG8C9VEubjiXpacFFQwX0gxXaoBAoGBALPTq9doISFrB5cdt+WAPP2BYJSaAa46y1pcM0h+zFizVZcaLQ6Oycl9nY2tCwrywTTrjJvWWt3JXhwogRWyYr6ozF30LYeGgzEs2YbExbu15xmzB4tlZpEbYUr/xp9r91dsLbhvPlC04hZ1Wyj7J3rnRi3XdFC2gd5Cio6Z6BfZAoGAeSinXs6CgYzmTfg70bJYsYNfcwt02dbGQ2LDOLopMr7TGKbtj5BSduYJg8d8oPedsAXAYxPWEeOY1Hs5PKGzt/GX1lSfHqmi36wcnN9NKPSuqYEY7HYzu3bivSCHmlWObUKJAu4/EQ4R6+Pd61nAIH8ObxcT5mEwmVg9Eyl5RcE=`
)

// pem , base64 ASCII, 用于x509 certificates
// pem, 常用后缀.crt, .pem, .cer, and .key
func transKey2Pem(key string, head, tail string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString(head)
	rawLen := 64
	keyLen := len(key)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(key[start:])
		} else {
			buffer.WriteString(key[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString(tail)
	pKey = buffer.String()
	return
}
