package space

type Planet = string

const EarthSeconds float64 = 31557600

var planetSeconds = map[string]float64{

	"Earth":   EarthSeconds,
	"Mercury": 0.2408467 * EarthSeconds,
	"Venus":   0.61519726 * EarthSeconds,
	"Mars":    1.8808158 * EarthSeconds,
	"Jupiter": 11.862615 * EarthSeconds,
	"Saturn":  29.447498 * EarthSeconds,
	"Uranus":  84.016846 * EarthSeconds,
	"Neptune": 164.79132 * EarthSeconds,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / planetSeconds[planet]

}
