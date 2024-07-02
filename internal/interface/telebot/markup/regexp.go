package markup

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
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
	DriverInfoFromCarToDriverInfoRegexp                = Regexp(`di_f_c_t_di:+\d{11}`)
	DriverInfoShowDriverLicenseInfoRegexp              = Regexp(`di_s_dli:+\d{11}`)
	DriverInfoShowCarInfoRegexp                        = Regexp(`di_s_ci:+\d{11}`)
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

func (self Regexp) GetPhoneNumber() entity.PhoneNumber {
	rx := regexp.MustCompile(`\+\d{11}`)
	submatch := rx.FindStringSubmatch(string(self))
	return entity.PhoneNumber(submatch[0])
}
