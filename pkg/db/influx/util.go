package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/model"
	"github.com/shopspring/decimal"
)

func ConvertLoadProfileDataSetToPoints(data []model.LoadProfileRecord) []*write.Point {
	points := make([]*write.Point, 0)
	for _, v := range data {
		f, _ := v.LoadProfile.Float64()
		p := influxdb2.NewPoint("load-profile",
			map[string]string{"status": v.Status},
			map[string]interface{}{"value": f},
			v.Ts.Ts)
		points = append(points, p)
	}
	return points
}

func ConvertWeatherDataSetToPoints(data []model.WeatherRecord) []*write.Point {
	points := make([]*write.Point, 0)
	for _, v := range data {
		num, err := decimal.NewFromString(v.Value)
		if err == nil {
			f, _ := num.Float64()
			p := influxdb2.NewPoint("weather-data-num",
				map[string]string{
					"location":  v.Location,
					"parameter": v.Parameter,
					"unit":      v.Unit,
				},
				map[string]interface{}{"value": f},
				v.Ts.Ts)
			points = append(points, p)
		} else {
			p := influxdb2.NewPoint("weather-data-string",
				map[string]string{
					"location":  v.Location,
					"parameter": v.Parameter,
					"unit":      v.Unit,
				},
				map[string]interface{}{"value": v.Value},
				v.Ts.Ts)
			points = append(points, p)
		}
	}
	return points
}
