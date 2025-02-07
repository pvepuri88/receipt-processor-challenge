package services

import (
	"math"
	"regexp"
	"receipt-processor/internal/models"
	"strconv"
	"strings"
	"time"
)

type PointsService struct{}

func NewPointsService() *PointsService {
	return &PointsService{}
}

func (s *PointsService) Calculate(r *models.Receipt) (int, error) {
	var total int
	total += countAlphaNum(r.Retailer)
	f, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return 0, err
	}
	if f == float64(int(f)) {
		total += 50
	}
	if int(math.Round(f*100))%25 == 0 {
		total += 25
	}
	total += (len(r.Items) / 2) * 5
	for _, it := range r.Items {
		desc := strings.TrimSpace(it.ShortDescription)
		if len(desc)%3 == 0 {
			p, err := strconv.ParseFloat(it.Price, 64)
			if err != nil {
				return 0, err
			}
			total += int(math.Ceil(p * 0.2))
		}
	}
	day, err := parseDay(r.PurchaseDate)
	if err != nil {
		return 0, err
	}
	if day%2 == 1 {
		total += 6
	}
	hour, min, err := parseTime(r.PurchaseTime)
	if err != nil {
		return 0, err
	}
	if (hour == 14 && min >= 0) || hour == 15 {
		total += 10
	}
	return total, nil
}

func countAlphaNum(s string) int {
	return len(regexp.MustCompile(`[A-Za-z0-9]`).FindAllString(s, -1))
}

func parseDay(d string) (int, error) {
	t, err := time.Parse("2006-01-02", d)
	return t.Day(), err
}

func parseTime(tm string) (int, int, error) {
	t, err := time.Parse("15:04", tm)
	return t.Hour(), t.Minute(), err
}

