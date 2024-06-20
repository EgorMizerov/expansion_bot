package fleet

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	entity2 "github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet/types"
	"github.com/pkg/errors"

	"github.com/google/uuid"

	//"github.com/EgorMizerov/expansion_bot/domain/shift/entity"
	requests "github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet/request"
	responses "github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet/response"
)

type Client struct {
	client *http.Client

	fleetHost string
	parkID    string
	clientID  string
	APIKey    string
}

func NewClient(fleetHost string, parkID string, clientID string, APIKey string) *Client {
	return &Client{
		client:    http.DefaultClient,
		fleetHost: fleetHost,
		parkID:    parkID,
		clientID:  clientID,
		APIKey:    APIKey,
	}
}

const (
	getOrdersURL            = "/v1/parks/orders/list"
	createDriverProfileURL  = "/v2/parks/contractors/driver-profile"
	getDriverProfileURL     = "https://fleet-api.taxi.yandex.net/v2/parks/contractors/driver-profile?contractor_profile_id=%s"
	putDriverProfileURL     = "https://fleet-api.taxi.yandex.net/v2/parks/contractors/driver-profile?contractor_profile_id=%s"
	createCar               = "https://fleet-api.taxi.yandex.net/v2/parks/vehicles/car"
	carBindingURL           = "https://fleet-api.taxi.yandex.net/v1/parks/driver-profiles/car-bindings?park_id=%s&car_id=%s&driver_profile_id=%s"
	getDriverIncomeURL      = "https://fleet.yandex.ru/api/v1/cards/driver/income"
	getContactorsURL        = "https://fleet.taxi.yandex.ru/api/fleet/contractor-profiles-manager/v1/active/list"
	createRegularChargesURL = "https://fleet.taxi.yandex.ru/api/v1/regular-charges/create"
	deleteRegularChargesURL = "https://fleet.taxi.yandex.ru/api/v1/regular-charges/terminate"
	getDriverSummaryURL     = "https://fleet.taxi.yandex.ru/api/reports-api/v2/summary/drivers/list"
)

func (self *Client) GetDriverIncome(ctx context.Context, requestBody requests.GetDriverIncomeRequest) (*GetDriverIncomeResponse, error) {
	return Internal[requests.GetDriverIncomeRequestBody, GetDriverIncomeResponse](ctx, http.MethodPost, getDriverIncomeURL, requestBody.ToBody())
}

func (self *Client) GetContractors(ctx context.Context, request requests.GetContractorsRequest) (*GetContractorsResponse, error) {
	return Internal[requests.GetContractorsRequestBody, GetContractorsResponse](ctx, http.MethodPost, getContactorsURL, request.ToBody())
}

func (self *Client) GetDriverProfile(ctx context.Context, driverID string) (*types.Contractor, error) {
	return Public[any, types.Contractor](ctx, http.MethodGet, fmt.Sprintf(getDriverProfileURL, driverID), nil)
}

func (self *Client) UpdateDriverProfile(ctx context.Context, driverID string, contractor *types.Contractor) error {
	_, err := Public[types.Contractor, any](ctx, http.MethodPut, fmt.Sprintf(putDriverProfileURL, driverID), contractor)
	return err
}

func (self *Client) CreateCar(ctx context.Context, car *types.CreateCarRequest, idempotencyToken string) (entity2.CarID, error) {
	id, err := Public[types.Car, types.VehicleID](ctx, http.MethodPost, createCar, car.ToBody(), idempotencyToken)
	if err != nil {
		return "", err
	}
	return entity2.CarID(id.VehicleID), nil
}

func (self *Client) CarBinding(ctx context.Context, carID entity2.CarID, driverID entity2.DriverID) error {
	_, err := Public[types.Car, any](ctx, http.MethodPut, fmt.Sprintf(carBindingURL, self.parkID, carID, driverID), nil)
	return err
}

type ID struct {
	ID string `json:"id"`
}

func (self *Client) CreateRegularCharges(ctx context.Context, request *requests.CreateRegularChargesRequest) (*ID, error) {
	return Internal[requests.CreateRegularCharges, ID](ctx, http.MethodPost, createRegularChargesURL, request.ToBody(), "v1")
}

func (self *Client) DeleteRegularCharges(ctx context.Context, chargingID string) error {
	_, err := Internal[map[string]interface{}, any](ctx, http.MethodPost, deleteRegularChargesURL, map[string]interface{}{"charging_id": chargingID}, "v1")
	return err
}

func (self *Client) GetDriversSummary(ctx context.Context, request *requests.GetDriversSummaryRequest) (*types.DriversSummary, error) {
	return Internal[requests.GetDriversSummaryRequestBody, types.DriversSummary](ctx, http.MethodPost, getDriverSummaryURL, request.ToBody(), "v1")
}

// =-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-
// =-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-
// =-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-=-=-=-=-==-=-=-==-=-=-=-=-=-=-==-

func (self *Client) CreateDriverProfile(ctx context.Context, requestBody requests.CreateDriverProfileRequest, idempotencyToken string) (*responses.ContractorProfileID, error) {
	body, err := convertBody(requestBody.ToBody())

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, self.getURL(createDriverProfileURL), body)
	request = self.withHeaders(request, header{key: "X-Idempotency-Token", value: idempotencyToken})

	resp, err := self.client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("status %s\nmessage: %s", resp.Status, string(msg)))
	}

	var response = new(responses.ContractorProfileID)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//func (self *Client) GetOrders(ctx context.Context, requestBody requests.GetOrdersRequest) (entity.Orders, error) {
//	body, _ := convertBody(requestBody.ToBody(self.parkID))
//
//	request, err := http.NewRequestWithContext(ctx, http.MethodPost, self.getURL(getOrdersURL), body)
//	request = self.withHeaders(request)
//
//	resp, err := self.client.Do(request)
//	if err != nil {
//		return nil, err
//	}
//
//	var response = new(responses.GetOrdersResponse)
//	err = json.NewDecoder(resp.Body).Decode(&response)
//	if err != nil {
//		return nil, err
//	}
//
//	return Reverse(response.ToEntities()), nil
//}

func (self *Client) getURL(path string) string {
	return fmt.Sprintf("https://%s%s", self.fleetHost, path)
}

type header struct {
	key   string
	value string
}

func (self *Client) withHeaders(request *http.Request, headers ...header) *http.Request {
	request.Header.Add("X-API-Key", self.APIKey)
	request.Header.Add("X-Client-ID", self.clientID)
	request.Header.Add("X-Park-ID", self.parkID)
	request.Header.Add("Accept-Language", "ru")

	for _, header := range headers {
		request.Header.Add(header.key, header.value)
	}

	return request
}

func convertBody(body any) (*strings.Reader, error) {
	if body == nil {
		return nil, nil
	}
	bytes, err := json.MarshalIndent(body, "", "	")
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(bytes))
	return reader, nil
}

func Reverse[T comparable](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	return append(Reverse(arr[1:]), arr[0])
}

func Internal[Request any, Response any](ctx context.Context, method string, url string, requestBody Request, v ...string) (*Response, error) {
	body, err := json.Marshal(requestBody)
	fmt.Println(string(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal requests body")
	}

	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create requests")
	}

	if v != nil {
		request.Header.Add("Accept", "*/*")
		request.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
		request.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8,ru;q=0.7")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("Content-Type", "text/plain;charset=UTF-8")
		request.Header.Add("Cookie", "i=/XMmqDMPFXAs4o5PqXuQJuBWQuGb2OG5Y6lsb+Ngj3VUeTL4YVOG+fehUln/0vVD9fo4qCREHImsyRuArrrwxt7e8xk=; yandexuid=9626565111718023029; yashr=8488256641718023029; yuidss=9626565111718023029; ymex=2033383032.yrts.1718023032; receive-cookie-deprecation=1; gdpr=0; _ym_uid=1718023031378997644; _ym_d=1718023032; yabs-vdrf=A0; is_gdpr=0; is_gdpr_b=CP7cPxD/gAI=; park_id=c53efe4fccd149a19839dd6d5fcce124; Session_id=3:1718028392.5.1.1718023146166:a1MRvA:3f.1.2:1|576391523.-1.2.3:1718023146|1478872821.-1.2.2:3857.3:1718027003|3:10289705.78552.b2ij6D4DWaMdWda-lMeu6Ca3DzQ; sessar=1.1190.CiAsxB4eLeuu2s1LKCt1dpU_j8hZVEl6FONkrcm_q7Vymw.RgbWciI1NgsdHGc6ZXtRlz5EnoyNPgqdzpdiaf2Fl44; sessionid2=3:1718028392.5.1.1718023146166:a1MRvA:3f.1.2:1|576391523.-1.2.3:1718023146|1478872821.-1.2.2:3857.3:1718027003|3:10289705.78552.fakesign0000000000000000000; yp=2033383092.multib.1#2033388392.udn.cDpFbGlaWi1mcnZy; L=WypqQ2sNQHBgWQBAQmhdAnNgUHgFDHkIAz4fOTR+ERAeJA==.1718028392.15769.355263.1bb5f2e836238fa67181c217a4c7df74; yandex_login=EliZZ-frvr; ys=udn.cDpFbGlaWi1mcnZy#c_chck.4011611830; _ym_isad=2; device_id=ad5118603f029b51cd93d65ecc72222ae7f9ae927; active-browser-timestamp=1718185427727; bh=Ej8iR29vZ2xlIENocm9tZSI7dj0iMTI1IiwiQ2hyb21pdW0iO3Y9IjEyNSIsIk5vdC5BL0JyYW5kIjt2PSIyNCIaBSJ4ODYiIhAiMTI1LjAuNjQyMi4xNDIiKgI/MDoJIldpbmRvd3MiQggiMTUuMC4wIkoEIjY0IlJcIkdvb2dsZSBDaHJvbWUiO3Y9IjEyNS4wLjY0MjIuMTQyIiwiQ2hyb21pdW0iO3Y9IjEyNS4wLjY0MjIuMTQyIiwiTm90LkEvQnJhbmQiO3Y9IjI0LjAuMC4wIiI=; _yasc=Pw07EfhGMLbDPBoIN2RIvV6efl0K1Bu/Ek7UQ0RZzQ785X1UabTihZrDBXNQLHlfFwtKYzbS; _ym_visorc=b")
		request.Header.Add("Host", "fleet.taxi.yandex.ru")
		request.Header.Add("Language", "ru")
		request.Header.Add("Origin", "https://fleet.taxi.yandex.ru")
		request.Header.Add("Referer", "https://fleet.taxi.yandex.ru/regular-charges/create?park_id=c53efe4fccd149a19839dd6d5fcce124&lang=ru")
		request.Header.Add("Sec-Ch-Ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
		request.Header.Add("Sec-Ch-Ua-Mobile", "?0")
		request.Header.Add("Sec-Ch-Ua-Platform", "\"Windows\"")
		request.Header.Add("Sec-Fetch-Dest", "empty")
		request.Header.Add("Sec-Fetch-Mode", "cors")
		request.Header.Add("Sec-Fetch-Site", "same-origin")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
		request.Header.Add("X-Client-Version", "fleet/9538")
		request.Header.Add("X-Idempotency-Token", uuid.New().String())
		request.Header.Add("X-Park-Id", "c53efe4fccd149a19839dd6d5fcce124")
	} else {
		request.Header.Add("Cookie", "yuidss=6136643771706273856; yandexuid=6136643771706273856; yashr=6040384471706273856; ymex=2021633857.yrts.1706273857; amcuid=9807862091706276400; my=YwA=; _ym_uid=1707681340991089808; _ym_d=1707681341; gdpr=0; skid=8830466051709397310; _ga=GA1.2.2021932776.1709826316; receive-cookie-deprecation=1; park_id=c53efe4fccd149a19839dd6d5fcce124; L=CRdKB1NcD2lKDkENB1ZfD11FcgN7eXJidB8nEWlqKEsCQQ==.1716529984.15748.339062.f2ab5150f31fd3f094f85af4906ac76c; yandex_login=EliZZ-frvr; cycada=UiorVCZE+UW/MRisOOqXETsOFVrQE4+R2yTq//ahM1c=; yp=2031889984.udn.cDpFbGlaWi1mcnZy#1719622337.hdrc.0#1723067058.szm.1%3A1920x1080%3A1903x945#1717498935.ygu.1#1737810029.p_sw.1706274028#2032324848.pcs.1#1722148846.atds.1#1737810259.p_cl.1706274258#2029504256.multib.1#1748479937.stltp.serp_bk-map_1_1716943937; device_id=afcd9705784cabd22ec2b56dfa4488121ac0bfe0a; active-browser-timestamp=1716979591106; is_gdpr=0; is_gdpr_b=CMWYJRCn/wEoAg==; yabs-dsp=mts_banner.YzU0V2lSd05UQUtXbzhTel9uZTlEUQ==; yabs-vdrf=IULrcRG3Jfem1LY9c1W5MaVK01Y1c3W3ems01u1DcZm0iGrW0BGrcD03ucqy0DlXbF05wTd41cELbcG1Iwse1NjjbIW1w6LK00; Session_id=3:1717604059.5.1.1708105817032:CxpPJQ:cd.1.2:1|576391523.-1.2.3:1708105817|1478872821.6038439.2.2:6038439.3:1714144256|3:10289465.521042.xBSX_dxAvQDQqzsXb9vbYnGlpQA; sessar=1.1190.CiBqgvFyDEHdS43cCL8c9jrTUOBIwK3H306HliD0ecCk3w.xlvPp7vi0tis3p4FzFe0rju6FybJxIgTK2R0ElIx4tU; sessionid2=3:1717604059.5.1.1708105817032:CxpPJQ:cd.1.2:1|576391523.-1.2.3:1708105817|1478872821.6038439.2.2:6038439.3:1714144256|3:10289465.521042.fakesign0000000000000000000; i=6jqf2wd5Zgrwmyx+7FLqWI9ouOEPLV6A+Y9OKGQOP2RJFltk8lQaCaNrDHx4ozlnJ7If0r4TQwUf2WwJbxqrWq/d0a0=; ys=udn.cDpFbGlaWi1mcnZy#c_chck.3663968973; _ym_isad=2; _ym_visorc=b; bh=Ej8iR29vZ2xlIENocm9tZSI7dj0iMTI1IiwiQ2hyb21pdW0iO3Y9IjEyNSIsIk5vdC5BL0JyYW5kIjt2PSIyNCIaBSJ4ODYiIg8iMTI1LjAuNjQyMi43OCIqAj8wMgIiIjoJIldpbmRvd3MiQggiMTUuMC4wIkoEIjY0IlJaIkdvb2dsZSBDaHJvbWUiO3Y9IjEyNS4wLjY0MjIuNzgiLCJDaHJvbWl1bSI7dj0iMTI1LjAuNjQyMi43OCIsIk5vdC5BL0JyYW5kIjt2PSIyNC4wLjAuMCIiWgI/MA==; _yasc=B2p5blLPSixnWgrarDQtAqQsaLVnTfv03B5lbUSZi1ZPTCASzTsOngcV5ud9sTFSaiDsBYIiAR5A/cJbf8xOwPfhEAafyuidss=6136643771706273856; yandexuid=6136643771706273856; yashr=6040384471706273856; ymex=2021633857.yrts.1706273857; amcuid=9807862091706276400; my=YwA=; _ym_uid=1707681340991089808; _ym_d=1707681341; gdpr=0; skid=8830466051709397310; _ga=GA1.2.2021932776.1709826316; receive-cookie-deprecation=1; park_id=c53efe4fccd149a19839dd6d5fcce124; L=CRdKB1NcD2lKDkENB1ZfD11FcgN7eXJidB8nEWlqKEsCQQ==.1716529984.15748.339062.f2ab5150f31fd3f094f85af4906ac76c; yandex_login=EliZZ-frvr; cycada=UiorVCZE+UW/MRisOOqXETsOFVrQE4+R2yTq//ahM1c=; yp=2031889984.udn.cDpFbGlaWi1mcnZy#1719622337.hdrc.0#1723067058.szm.1%3A1920x1080%3A1903x945#1717498935.ygu.1#1737810029.p_sw.1706274028#2032324848.pcs.1#1722148846.atds.1#1737810259.p_cl.1706274258#2029504256.multib.1#1748479937.stltp.serp_bk-map_1_1716943937; device_id=afcd9705784cabd22ec2b56dfa4488121ac0bfe0a; active-browser-timestamp=1716979591106; is_gdpr=0; is_gdpr_b=CMWYJRCn/wEoAg==; yabs-dsp=mts_banner.YzU0V2lSd05UQUtXbzhTel9uZTlEUQ==; yabs-vdrf=IULrcRG3Jfem1LY9c1W5MaVK01Y1c3W3ems01u1DcZm0iGrW0BGrcD03ucqy0DlXbF05wTd41cELbcG1Iwse1NjjbIW1w6LK00; Session_id=3:1717604059.5.1.1708105817032:CxpPJQ:cd.1.2:1|576391523.-1.2.3:1708105817|1478872821.6038439.2.2:6038439.3:1714144256|3:10289465.521042.xBSX_dxAvQDQqzsXb9vbYnGlpQA; sessar=1.1190.CiBqgvFyDEHdS43cCL8c9jrTUOBIwK3H306HliD0ecCk3w.xlvPp7vi0tis3p4FzFe0rju6FybJxIgTK2R0ElIx4tU; sessionid2=3:1717604059.5.1.1708105817032:CxpPJQ:cd.1.2:1|576391523.-1.2.3:1708105817|1478872821.6038439.2.2:6038439.3:1714144256|3:10289465.521042.fakesign0000000000000000000; i=6jqf2wd5Zgrwmyx+7FLqWI9ouOEPLV6A+Y9OKGQOP2RJFltk8lQaCaNrDHx4ozlnJ7If0r4TQwUf2WwJbxqrWq/d0a0=; ys=udn.cDpFbGlaWi1mcnZy#c_chck.3663968973; _ym_isad=2; _ym_visorc=b; bh=Ej8iR29vZ2xlIENocm9tZSI7dj0iMTI1IiwiQ2hyb21pdW0iO3Y9IjEyNSIsIk5vdC5BL0JyYW5kIjt2PSIyNCIaBSJ4ODYiIg8iMTI1LjAuNjQyMi43OCIqAj8wMgIiIjoJIldpbmRvd3MiQggiMTUuMC4wIkoEIjY0IlJaIkdvb2dsZSBDaHJvbWUiO3Y9IjEyNS4wLjY0MjIuNzgiLCJDaHJvbWl1bSI7dj0iMTI1LjAuNjQyMi43OCIsIk5vdC5BL0JyYW5kIjt2PSIyNC4wLjAuMCIiWgI/MA==; _yasc=B2p5blLPSixnWgrarDQtAqQsaLVnTfv03B5lbUSZi1ZPTCASzTsOngcV5ud9sTFSaiDsBYIiAR5A/cJbf8xOwPfhEAaf")
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-Park-ID", "c53efe4fccd149a19839dd6d5fcce124")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do requests")
	}

	if response.Header.Get("Content-Encoding") == "gzip" {
		response.Body, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to un-gzip")
		}
	}

	//raw, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(raw))
	//response.Body = io.NopCloser(bytes.NewReader(raw))

	var responseBody *Response
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return responseBody, nil
}

func Public[Request any, Response any](ctx context.Context, method string, url string, requestBody *Request, idempotencyToken ...string) (*Response, error) {
	var body io.Reader

	if requestBody != nil {
		marshaled, err := json.Marshal(requestBody)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal requests body")
		}
		body = bytes.NewReader(marshaled)
	}

	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create requests")
	}

	request.Header.Add("X-API-Key", "lIXMtczJiCHaYMwVRUgSOHxxXtRxJnGvNol")
	request.Header.Add("X-Client-ID", "taxi/park/c53efe4fccd149a19839dd6d5fcce124")
	request.Header.Add("X-Park-ID", "c53efe4fccd149a19839dd6d5fcce124")
	request.Header.Add("Content-Type", "application/json")
	if len(idempotencyToken) != 0 {
		request.Header.Add("X-Idempotency-Token", idempotencyToken[0])
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do requests")
	}

	//raw, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(raw))
	//response.Body = io.NopCloser(bytes.NewReader(raw))

	var responseBody Response
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return &responseBody, nil
}
