package gorte

import (
	"net/http"
	"time"
)

type AcceptedOffersResp struct {
	AcceptedOffers []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Type      string    `json:"type"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Direction   string    `json:"direction"`
			Value       int       `json:"value"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"accepted_offers"`
}

func (s *Market) GetAcceptedOffers(opt *Period) (*AcceptedOffersResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "/open_api/balancing_capacity/v4/accepted_offers", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *AcceptedOffersResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type ProcuredReservesResp struct {
	ProcuredReserves []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Type      string    `json:"type"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Direction   string    `json:"direction"`
			Value       int       `json:"value"`
			Price       float64   `json:"price"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"procured_reserves"`
}

func (s *Market) GetProcuredReservesResp(opt *Period) (*ProcuredReservesResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "/open_api/balancing_capacity/v4/procured_reserves", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *ProcuredReservesResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type PeakDailyMarginsResp struct {
	PeakDailyMargins []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate         time.Time `json:"start_date"`
			EndDate           time.Time `json:"end_date"`
			UpdatedDate       time.Time `json:"updated_date"`
			MorningPeakPeriod struct {
				DurationBeforeMarginPeriod string `json:"DURATION_BEFORE_MARGIN_PERIOD"`
				BeginMarginPeriod          string `json:"BEGIN_MARGIN_PERIOD"`
				EndMarginPeriod            string `json:"END_MARGIN_PERIOD"`
				RequiredMargin             string `json:"REQUIRED_MARGIN"`
				MaOffers                   string `json:"MA_OFFERS"`
				CompOffers                 string `json:"COMP_OFFERS"`
				RESCUE                     string `json:"RESCUE "`
				IMBALANCE                  string `json:"IMBALANCE "`
				SimulatedMaOffers          string `json:"SIMULATED_MA_OFFERS"`
				SimulatedCompOffers        string `json:"SIMULATED_COMP_OFFERS"`
				SimulatedRescue            string `json:"SIMULATED_RESCUE"`
				SIMULATEDREQUIREDMARGIN    string `json:"SIMULATED_REQUIRED_MARGIN "`
			} `json:"morning_peak_period"`
			EveningPeakPeriod struct {
				DURATIONBEFOREMARGINPERIOD string `json:"DURATION_BEFORE_MARGIN_PERIOD "`
				BEGINMARGINPERIOD          string `json:"BEGIN_MARGIN_PERIOD "`
				ENDMARGINPERIOD            string `json:"END_MARGIN_PERIOD "`
				RequiredMargin             string `json:"REQUIRED_MARGIN"`
				MaOffers                   string `json:"MA_OFFERS"`
				CompOffers                 string `json:"COMP_OFFERS"`
				Rescue                     string `json:"RESCUE"`
				Imbalance                  string `json:"IMBALANCE"`
				SimulatedMaOffers          string `json:"SIMULATED_MA_OFFERS"`
				SimulatedCompOffers        string `json:"SIMULATED_COMP_OFFERS"`
				SimulatedRescue            string `json:"SIMULATED_RESCUE"`
				SIMULATEDREQUIREDMARGIN    string `json:"SIMULATED_REQUIRED_MARGIN "`
			} `json:"evening_peak_period"`
		} `json:"values"`
	} `json:"peak_daily_margins"`
}

func (s *Market) GetPeakDailyMargins(opt *Period) (*PeakDailyMarginsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/peak_daily_margins", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *PeakDailyMarginsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type InsufficientsOffersResp struct {
	InsufficientsOffers struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			UpdatedDate time.Time `json:"updated_date"`
			Type        string    `json:"type"`
			Nature      string    `json:"nature"`
			Direction   string    `json:"direction"`
		} `json:"values"`
	} `json:"insufficients_offers"`
}

func (s *Market) GetInsufficientsOffers(opt *Period) (*InsufficientsOffersResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/insufficients_offers", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *InsufficientsOffersResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type ImbalanceResp struct {
	Imbalance []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       int       `json:"value"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"imbalance"`
}

func (s *Market) GetImbalance(opt *Period) (*ImbalanceResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/imbalance", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *ImbalanceResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type AggregatedoffersEnergybidsResp struct {
	AggregatedoffersEnergybids []struct {
		StartDate          time.Time `json:"start_date"`
		EndDate            time.Time `json:"end_date"`
		OfferedRRstdVolume []struct {
			StartDate time.Time `json:"start_date"`
			EndDate   time.Time `json:"end_date"`
			Values    struct {
				DownwardOfferedVolumeRRStdRestricted float64 `json:"DownwardOfferedVolumeRRStdRestricted"`
				UpwardOfferedVolumeRRStd             float64 `json:"UpwardOfferedVolumeRRStd"`
				DownwardOfferedVolumeRRStd           float64 `json:"DownwardOfferedVolumeRRStd"`
				UpwardOfferedVolumeRRStdRestricted   float64 `json:"UpwardOfferedVolumeRRStdRestricted"`
			} `json:"values"`
		} `json:"Offered_RRstd_Volume"`
		OfferedVolume []struct {
			StartDate time.Time `json:"start_date"`
			EndDate   time.Time `json:"end_date"`
			Values    struct {
				UpwardOfferedVolumemFRR         int `json:"UpwardOfferedVolumemFRR"`
				UpwardOfferedVolumeRRSpecific   int `json:"UpwardOfferedVolumeRRSpecific"`
				DownwardOfferedVolumeFCR        int `json:"DownwardOfferedVolumeFCR"`
				DownwardOfferedVolumeRRSpecific int `json:"DownwardOfferedVolumeRRSpecific"`
				UpwardOfferedVolumeFCR          int `json:"UpwardOfferedVolumeFCR"`
				UpwardOfferedVolumeaFRR         int `json:"UpwardOfferedVolumeaFRR"`
				DownwardOfferedVolumeaFRR       int `json:"DownwardOfferedVolumeaFRR"`
				DownwardOfferedVolumemFRR       int `json:"DownwardOfferedVolumemFRR"`
			} `json:"values"`
		} `json:"Offered_Volume"`
	} `json:"aggregatedoffers_energybids"`
}

func (s *Market) GetAggregatedoffersEnergybids(opt *Period) (*AggregatedoffersEnergybidsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/aggregatedoffers_energybids", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *AggregatedoffersEnergybidsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type IndividualoffersEnergybidsResp struct {
	IndividualoffersEnergybids []struct {
		StartDate        time.Time `json:"start_date"`
		EndDate          time.Time `json:"end_date"`
		TotalPagesNumber int       `json:"total_pages_number"`
		TotalItemsNumber int       `json:"total_items_number"`
		CurrentPage      int       `json:"current_page"`
		Mesure           []struct {
			StartDate time.Time `json:"start_date"`
			EndDate   time.Time `json:"end_date"`
			Values    []struct {
				Quantity     int     `json:"Quantity"`
				PriceAmount  float64 `json:"PriceAmount"`
				OfferType    string  `json:"OfferType"`
				Dmo          int     `json:"DMO"`
				ReserveType  string  `json:"ReserveType"`
				RowDirection string  `json:"RowDirection"`
				Status       string  `json:"Status"`
			} `json:"values"`
		} `json:"mesure"`
	} `json:"individualoffers_energybids"`
}

func (s *Market) GetIndividualoffersEnergybids(opt *Period) (*IndividualoffersEnergybidsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/individualoffers_energybids", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *IndividualoffersEnergybidsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type GetMarginsDataOptions struct {
	Date Time   `url:"date"`
	Sens string `url:"sens"`
	Type string `url:"type"`
}

type MarginsDataResp []struct {
	StartDate       string `json:"startDate"`
	ComputationTime string `json:"computationTime"`
	ProjectionTime  string `json:"projectionTime"`
	MarginsData     []struct {
		StudiedTime                       string `json:"studiedTime"`
		ForecastedImbalanceRs             int    `json:"forecastedImbalanceRs"`
		AvailableMarginsNormalUp          int    `json:"availableMarginsNormalUp"`
		AvailableMarginsNormalDown        int    `json:"availableMarginsNormalDown"`
		AvailableMarginsComplementaryUp   int    `json:"availableMarginsComplementaryUp"`
		AvailableMarginsComplementaryDown int    `json:"availableMarginsComplementaryDown"`
		AvailableMarginsReliefUp          int    `json:"availableMarginsReliefUp"`
		AvailableMarginsReliefDown        int    `json:"availableMarginsReliefDown"`
		RequiredMarginsOperationnalUp     int    `json:"requiredMarginsOperationnalUp"`
		RequiredMarginsOperationnalDown   int    `json:"requiredMarginsOperationnalDown"`
	} `json:"marginsData"`
}

func (s *Market) GetMarginsData(opt *GetMarginsDataOptions) (*MarginsDataResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/individualoffers_energybids", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *MarginsDataResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type DailyProcuredReservesResp struct {
	DailyProcuredReserves []struct {
		StartDate time.Time     `json:"start_date"`
		EndDate   time.Time     `json:"end_date"`
		Type      string        `json:"type"`
		Values    []interface{} `json:"values"`
	} `json:"daily_procured_reserves"`
}

func (s *Market) GetDailyProcuredReserves(opt *Period) (*DailyProcuredReservesResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/daily_procured_reserves", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *DailyProcuredReservesResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type NeedsResp struct {
	Needs []struct {
		StartDate     time.Time `json:"start_date"`
		EndDate       time.Time `json:"end_date"`
		Product       string    `json:"product"`
		FlowDirection string    `json:"flow_direction"`
		Volume        int       `json:"volume"`
	} `json:"needs"`
}

func (s *Market) GetNeeds(opt *Period) (*NeedsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/tso_need_for_procured_reserves", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *NeedsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type AggregatedoffersAFRREnergybidsResp struct {
	AggregatedoffersAFRREnergybids []struct {
		StartDate  time.Time `json:"start_date"`
		EndDate    time.Time `json:"end_date"`
		Chronicles []struct {
			StartDate time.Time `json:"start_date"`
			EndDate   time.Time `json:"end_date"`
			Values    struct {
				DownwardOfferedVolumeaFRR int     `json:"DownwardOfferedVolumeaFRR"`
				UpwardOfferedVolumeaFRR   int     `json:"UpwardOfferedVolumeaFRR"`
				DownwardSharedVolumeaFRR  int     `json:"DownwardSharedVolumeaFRR"`
				UpwardSharedVolumeaFRR    int     `json:"UpwardSharedVolumeaFRR"`
				DownwardAveragePriceaFRR  float64 `json:"DownwardAveragePriceaFRR"`
				UpwardAveragePriceaFRR    float64 `json:"UpwardAveragePriceaFRR"`
			} `json:"values"`
		} `json:"chronicles"`
	} `json:"aggregatedoffers_aFRR_energybids"`
}

func (s *Market) GetAggregatedoffersAFRREnergybids(opt *Period) (*AggregatedoffersAFRREnergybidsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/balancing_capacity/v4/aggregatedoffers_afrr_energybids", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *AggregatedoffersAFRREnergybidsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}
