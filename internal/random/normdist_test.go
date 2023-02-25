package random_test

import (
	"image/color"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/pprishchepa/whoami/internal/random"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

//goland:noinspection GoUnusedGlobalVariable
var res float64

func BenchmarkNormFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = random.NormFloat64(0.06, 2)
	}
}

func TestNormFloat64(t *testing.T) {
	rand.Seed(1)
	assert.Equal(t, 0.8578302881800055, random.NormFloat64(0.06, 2))
	assert.Equal(t, 0.1417047235875345, random.NormFloat64(0.06, 2))
}

func TestNormFloat64_MakeHist(t *testing.T) {
	t.Skip()
	rand.Seed(time.Now().UnixNano())
	var dist []float64
	for i := 0; i < 10000; i++ {
		dist = append(dist, random.NormFloat64(0.06, 2))
	}
	hist(dist, "normdist", "Normal Distribution")
}

func hist(dist []float64, name, title string) {
	n := len(dist)
	vals := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		vals[i] = dist[i]
	}

	plt := plot.New()
	plt.Title.Text = title
	hist, err := plotter.NewHist(vals, 25) // 25 bins
	if err != nil {
		log.Println("Cannot plot:", err)
	}
	hist.FillColor = color.RGBA{R: 255, G: 127, B: 80, A: 255} // coral color
	plt.Add(hist)

	err = plt.Save(400, 200, name+".png")
	if err != nil {
		log.Panic(err)
	}
}
