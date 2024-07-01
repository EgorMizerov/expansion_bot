package markup

import (
	"fmt"
	"regexp"
	"strconv"
)

type Regexp string

const (
	DriverPhoneRegexp                                  = Regexp(`\+7\d{10}`)
	RegistrationApplicationIDRegexp                    = Regexp(`^\d{5,7}$`)
	SetFixWorkRuleForApplicationRegexp                 = Regexp(`fwr:\d`)
	SetFixSelfEmployedWorkRuleForApplicationRegexp     = Regexp(`fsewr:\d`)
	SetPercentWorkRuleForApplicationRegexp             = Regexp(`pwr:\d`)
	SetPercentSelfEmployedWorkRuleForApplicationRegexp = Regexp(`psewr:\d`)
	SetPerDayWorkRuleForApplicationRegexp              = Regexp(`pdwr:\d`)
	ConfirmRegistrationApplicationRegexp               = Regexp(`cf_ra:\d`)
)

func (self Regexp) Endpoint() string {
	return fmt.Sprintf("rx:%s", string(self))
}

func (self Regexp) GetNumber() int {
	rx := regexp.MustCompile(`\d+`)
	submatch := rx.FindStringSubmatch(string(self))
	n, _ := strconv.Atoi(submatch[0])
	return n
}
