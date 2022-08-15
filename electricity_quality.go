package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type QualityDataResp struct {
	QualityData []struct {
		QualityMeterPointID string `json:"quality_meter_point_id"`
		DataList            []struct {
			DataType  string    `json:"data_type"`
			Phase     string    `json:"phase"`
			Unit      string    `json:"unit"`
			StartDate time.Time `json:"start_date,omitempty"`
			EndDate   time.Time `json:"end_date,omitempty"`
			Data      []struct {
				Timestamp time.Time `json:"timestamp"`
				Value     float64   `json:"value"`
			} `json:"data"`
			ValueType         string      `json:"value_type,omitempty"`
			NominalDataPeriod int         `json:"nominal_data_period,omitempty"`
			HarmonicNumber    interface{} `json:"harmonic_number,omitempty"`
		} `json:"data_list"`
	} `json:"quality_data"`
}

type GetQualityDataOptions struct {
	ID        int        `url:"quality_meter_point_id"`
	StartDate utils.Time `url:"start_date"`
	EndDate   utils.Time `url:"end_date"`
	Type      string     `url:"data_type"`
}

func (co *consumption) GetQualityData(opt *GetQualityDataOptions) (*QualityDataResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "private_api/electricity_quality/v1/quality_data", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	var sig *QualityDataResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
