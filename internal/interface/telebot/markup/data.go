package markup

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strconv"
)

const (
	AdminMenuData                       = "admin"
	DriverInfoEditData                  = "admin:driver_info:edit"
	DriverInfoShowCarInfoBackData       = "admin:driver_info:car_info:back"
	DriverInfoShowCarInfoData           = "admin:driver_info:car_info"
	DriverInfoShowDriverLicenseInfoData = "admin:driver_info:driver_license_info"
)

const (
	MainMenuData = "main_menu"

	InformationMenuData       = "information_main_menu" // Информация о парке
	ProfileManagementMenuData = "profile_management"    // Управление профилем
	StatisticsMenuData        = "statistics"            // Статистика

	// -=-=-=-Information -=-=-=-
	Information_AboutUsData       = "information:about_us"       // Информация о нас
	Inforamtion_FAQData           = "information:faq"            // Часто задаваемые вопросы
	Information_BeginningWorkData = "information:beginning_work" // Начало работы

	Information_FAQ_WithdrawData             = "information:faq:withdraw"               // Вывод средств
	Information_FAQ_SelfEmployeedData        = "information:faq:self_employed"          // 	// Самозанятый
	Information_FAQ_LicensesAndRegistersData = "information:faq:licenses_and_registers" // Лицензии и реестр
	Information_FAQ_TrafficRulesData         = "information:faq:traffic_rules"          // Информация ПДД

	Information_FAQ_SelfEmployeed_HowToBeData     = "information:faq:se:htb"      // Как стать самозанятым
	Information_FAQ_SelfEmployeed_WhatIsItForData = "information:faq:se:wiif"     // Для чего это нужно
	Information_FAQ_SelfEmployeed_BenefitsData    = "information:faq:se:benefits" // Какие от этого плюсы

	Information_FAQ_LicensesAndRegisters_HowToJoinData            = "information:faq:lar:htj"   // Как встать в реестр перевозчика
	Information_FAQ_LicensesAndRegisters_WhatIsItForData          = "information:faq:lar:wiif"  // Для чего это нужно
	Information_FAQ_LicensesAndRegisters_HowToApplyForLicenseData = "information:faq:lar:htafl" // Как оформить лицензию

	Information_FAQ_TrafficRules_DialogueWithPassengerData = "information:faq:tr:dwp" // Диалог с пассажиром
	Information_FAQ_TrafficRules_TrafficViolationData      = "information:faq:tr:tv"  // Нарушения ПДД

	Information_BeginningWork_BotTrainingData       = "information:beginning_work:bt"   // Обученеи по боту
	Information_BeginningWork_TaximeterTrainingData = "information:beginning_work:tt"   // Обучени в таксометре
	Information_BeginningWork_HowToEarnMoreData     = "information:beginning_work:htem" // Как заработать больше

	Information_BeginningWork_HowToEarnMore_GBOData                = "information:beginning_work:htem:gbo"  // ГБО
	Information_BeginningWork_HowToEarnMore_ReferalBenefitsData    = "information:beginning_work:htem:rb"   // Плюсы реферальной системы
	Information_BeginningWork_HowToEarnMore_AboutPersonalLinksData = "information:beginning_work:htem:apl"  // Информация про персональные ссылки
	Information_BeginningWork_HowToEarnMore_HowToKeepRecordsData   = "information:beginning_work:htem:htkr" // Как правильно вести учет денег
	// -=-=-=-=-=-=-=-=--=-=-=-=-

	// -=-=-=- Profile Managment -=-=-=-
	ProfileManagement_ChangePhoneNumber     = "profile_management:change_phone_number"     // Изменить номер телефона
	ProfileManagement_ChangeCar             = "profile_management:change_car"              // Сменить автомобиль
	ProfileManagement_ChangePaymentMethod   = "profile_management:change_payment_method"   // Сменить способ оплаты
	ProfileManagement_BabyCarSeat           = "profile_management:baby_car_seat"           // Детские кресла
	ProfileManagement_ChangeTariff          = "profile_management:change_tariff"           // Сменить тариф
	ProfileManagement_ChangeTariff_Percent  = "profile_management:change_tariff:percent"   // Сменить тариф
	ProfileManagement_ChangeTariff_FixDay   = "profile_management:change_tariff:fix_day"   // Сменить тариф
	ProfileManagement_ChangeTariff_FixOrder = "profile_management:change_tariff:fix_order" // Сменить тариф
	// -=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	//-=-=-=-=-=- Statistics -=-=-=-=-=-
	Statistics_UnpaidData         = "statistics:unpaid"         // Неоплаченные заказы
	Statistics_ProfitData         = "statistics:profit"         // Заработок за период
	Statistics_IrregularitiesData = "statistics:irregularities" // Нарушения в таксометре

	Statistics_Profit_DayData        = "statistics:profit:day"         // Заработок за период
	Statistics_Profit_WeekData       = "statistics:profit:week"        // Заработок за период
	Statistics_Profit_MonthData      = "statistics:profit:month"       // Заработок за период
	Statistics_Profit_ThreeMonthData = "statistics:profit:three_month" // Заработок за период

	Statistics_Irregularities_DayData        = "statistics:irregularities:day"         // Заработок за период
	Statistics_Irregularities_WeekData       = "statistics:irregularities:week"        // Заработок за период
	Statistics_Irregularities_MonthData      = "statistics:irregularities:month"       // Заработок за период
	Statistics_Irregularities_ThreeMonthData = "statistics:irregularities:three_month" // Заработок за период

	Statistics_Unpaid_DayData        = "statistics:unpaid:day"         // Заработок за период
	Statistics_Unpaid_WeekData       = "statistics:unpaid:week"        // Заработок за период
	Statistics_Unpaid_MonthData      = "statistics:unpaid:month"       // Заработок за период
	Statistics_Unpaid_ThreeMonthData = "statistics:unpaid:three_month" // Заработок за период
	//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

	//-=-=-=-=-= Shift's =-=-=-=-=-
	ShiftMenuData    = "shift:menu"   // Начать смену
	Shift_StartData  = "shift:start"  // Начать смену
	Shift_FinishData = "shift:finish" // Закончить смену
	//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	//-=- Bonuses And Promotions -=-
	BaPMenuData        = "bap:menu"
	BaP_YourTaxiData   = "bap:your_taxi"
	BaP_BestDriverData = "bap:best_driver"
	BaP_50kData        = "bap:50k"
	BaP_CoffeeData     = "bap:coffee"
	BaP_GasData        = "bap:gas"

	BaP_YourTaxi_StatisticsData = "bap:your_taxi:statistics"
	//-=-=-=-=-=-=-=-=-=-=-=-=-=--=-

	//-=-=-=-=-=- Admin -=-=-=-=-=-
	Admin_UserRegistrationData = "admin:user_registration" // Регистрация пользователя

	Admin_CreateDriverData                              = "admin_create_driver"                          // Регистрация водителя
	Admin_CreateDriver_ConfirmData                      = "admin:create_driver:confirm"                  // Подтвердить регистрацию водителя
	Admin_CreateDriver_AbortData                        = "admin:create_driver:abort"                    // Прервать регистрацию водителя
	Admin_CreateDriver_InputIDData                      = "admin:create_driver:id"                       // Получить идентификатор пользователя
	Admin_CreateDriver_InputFullnameData                = "admin:create_driver:fullname"                 // Получить ФИО пользователя
	Admin_CreateDriver_InputPhoneNumberData             = "admin:create_driver:phone_number"             // Получить номер телефона пользователя
	Admin_CreateDriver_InputAddressData                 = "admin:create_driver:address"                  // Получить адрес проживания пользователя
	Admin_CreateDriver_InputIsSelfEmployedData          = "admin:create_driver:is_self_employed"         // Получить статус самозанятасти пользователя
	Admin_CreateDriver_InputTariffData                  = "admin:create_driver:tariff"                   // Получить тариф пользователя
	Admin_CreateDriver_InputCardNumberData              = "admin:create_driver:card_number"              // Получить номер карты пользователя
	Admin_CreateDriver_InputDrivingExperienceData       = "admin:create_driver:driving_experience"       // Получить водительский стаж с пользователя
	Admin_CreateDriver_InputRegistrationCerteficateData = "admin:create_driver:registration_certificate" // Получить серию и номер ВУ пользователя
	Admin_CreateDriver_InputLicenseCountry              = "admin:create_driver:license_country"          // Получить страну выдачи ВУ пользователя
	Admin_CreateDriver_InputLicenseIssueDate            = "admin:create_driver:license_issue_date"       // Получить дата получения ВУ пользователя
	Admin_CreateDriver_InputLicenseExpiryDate           = "admin:create_driver:license_expiry_date"      // Получить дата истечения действия ВУ пользователя
	Admin_CreateDriver_InputCarBrand                    = "admin:create_driver:car_brand"                // Получить марку автомобиля пользователя
	Admin_CreateDriver_InputCarModel                    = "admin:create_driver:car_model"                // Получить модель автомобиля пользователя
	Admin_CreateDriver_InputCarColor                    = "admin:create_driver:car_color"                // Получить цвет автомобиля пользователя
	Admin_CreateDriver_InputCarYear                     = "admin:create_driver:car_year"                 // Получить год выпуска автомобиля пользователя
	Admin_CreateDriver_InputCarVIN                      = "admin:create_driver:car_vin"                  // Получить VIN номер автомобиля пользователя
	Admin_CreateDriver_InputLicensePlateNumber          = "admin:create_driver:license_plate_number"     // Получить гос. номер автомобиля пользователя
	Admin_CreateDriver_InputRefferalKey                 = "admin:create_driver:referral_key"             // Получить реферальный ключ пользователя

	Admin_CardsInfo        = "admin:cards_info"        // Карты пользователей
	Admin_CardsInfo_Report = "admin:cards_info:report" // Сформировать отчёт

	Admin_CarRegistrationData         = "admin:car_registration"         // Регистрация автомобиля
	Admin_CarRegistration_CancelData  = "admin:car_registration:cancel"  // Отменить регистрацию автомобиля
	Admin_CarRegistration_ConfirmData = "admin:car_registration:confirm" // Подтвердить регистрацию автомобиля
	Admin_PromoData                   = "admin:promotions"               // ВКЛ/ВЫКЛ акций
	Admin_UserCardsData               = "admin:user_cards"               // У кого не Тинькофф
	Admin_InactiveUsersData           = "admin:inactive_users"           // Неактивные пользователи
	Admin_MetricsData                 = "admin:metrics"                  // Неактивные пользователи
	Admin_DeleteReferalData           = "admin:delete_referal"           // Неактивные пользователи
	Admin_RequestHistoryData          = "admin:request_history"          // Неактивные пользователи
	Admin_MessageForUsersData         = "admin:message_for_users"        // Неактивные пользователи

	Admin_UserRegistration_CancelData     = "admin:user_registration:cancel"     // Прервать регистрацию
	Admin_UserRegistration_RegistrateData = "admin:user_registration:registrate" // Подтвердить регистрацию
	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

	TODOData = "todo"
)

// Validate CallbackData length. Max 64 chars.
var _ = func() any {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "handlers/markup/data.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("markup", fset, []*ast.File{node}, info)
	if err != nil {
		panic(err) // type error
	}

	unqiue := map[string]struct{}{}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}
		for _, spec := range genDecl.Specs {
			cnst := spec.(*ast.ValueSpec)
			for _, name := range cnst.Names {
				value := info.ObjectOf(name).(*types.Const).Val().ExactString()
				_, ok := unqiue[value]
				if ok {
					panic("Duplicate value name - " + value)
				}
				unqiue[value] = struct{}{}

				if len([]rune(value)) > 64 {
					panic("Invalid CallbackData length! Value - " + value + "length - " + strconv.Itoa(len(value)))
				}
			}
		}
	}
	return nil
}
