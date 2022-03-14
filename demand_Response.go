package gorte

import (
	"net/http"
	"time"
)

type OperatorsResp struct {
	Operators []struct {
		StartDate                    time.Time `json:"start_date"`
		EndDate                      time.Time `json:"end_date"`
		UpdatedDate                  time.Time `json:"updated_date"`
		EicCode                      string    `json:"eic_code"`
		Name                         string    `json:"name"`
		TrialNebefRulesAgreement     bool      `json:"trial_nebef_rules_agreement"`
		TrialNebefRulesQualification bool      `json:"trial_nebef_rules_qualification"`
		NebefRulesRecognition        bool      `json:"nebef_rules_recognition"`
		NebefRulesQualification      bool      `json:"nebef_rules_qualification"`
	} `json:"operators"`
}

func (co *consumption) GetOperators(opt *Period) (*OperatorsResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/demand_response/v1/operators", opt)
	if err != nil {
		return nil, nil, err
	}
	sig := &OperatorsResp{}
	resp, err := c.Do(req, sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}

type VolumesResp struct {
	Volumes []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate        time.Time `json:"start_date"`
			EndDate          time.Time `json:"end_date"`
			ProgramsRetained int       `json:"programs_retained"`
			UpdatedDate      time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"volumes"`
}

func (co *consumption) GetVolumes(opt *Period) (*VolumesResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/demand_response/v1/volumes", opt)
	if err != nil {
		return nil, nil, err
	}
	sig := &VolumesResp{}
	resp, err := c.Do(req, sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}
