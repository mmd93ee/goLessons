package main

/*

Make sure the libraries are imported:

   go get -u <name>
   gota/dataframe: github.com/kniren/gota/dataframe
   gonmum/plot: gonum.org/v1/plot
   gonum/gonum: gonum.org/v1/gonum

If using VSCode then update debugger - analysis tools can be updated via VSCode:

   delve: github.com/go-delve/delve/cmd/dlv
   godef: github.com/rogpeppe/godef

Following Analysis tools can be installed via Console:

   gopkgs
   go-outline
   gocode-gomod
   godef
   golint

*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	useGophernotes := false
	datapath := "/Users/matt/github/datasets/ML/bmi/500_Person_Gender_Height_Weight_Index.csv"

	data, err := ioutil.ReadFile(datapath)
	if err != nil {
		fmt.Println("Error loading datafile: ", err)
	}

	df := dataframe.ReadCSV(bytes.NewReader(data))

	// Dump out the dataframe to make sure it loaded properly
	fmt.Println("Dataframe Details:", df.Describe())

	HistogramData(SeriesToPlotValues(df, "Height"), "Height Distribution", useGophernotes)
	HistogramData(SeriesToPlotValues(df, "Weight"), "Weight Distribution", useGophernotes)

}

// SeriesToPlotValues takes a dataframe and a column name and return a plotter.Values slice to be used as a graph indices.  Panic if column does not exist.
func SeriesToPlotValues(df dataframe.DataFrame, col string) plotter.Values {

	fmt.Println(df)
	rows, _ := df.Dims()
	v := make(plotter.Values, rows)
	s := df.Col(col)

	// Iterate row index by adding a value as a float for each item in the
	for i := 0; i < rows; i++ {
		v[i] = s.Elem(i).Float()
	}

	return v
}

// HistogramData produces a byteslice of jpg data for a histogram of the column with the name 'col' in the database 'df'
func HistogramData(v plotter.Values, title string, useGophernotes bool) []byte {

	// Make a plot and set the title
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = title
	h, err := plotter.NewHist(v, 10)
	if err != nil {
		panic(err)
	}

	// h.Nomalise(1) // Uncomment to normalise the data uuder the graph

	p.Add(h)
	var b bytes.Buffer

	if useGophernotes {
		// Use below to integrate with Gophernotes
		w, err := p.WriterTo(5*vg.Inch, 4*vg.Inch, "jpg")
		if err != nil {
			panic(err)
		}

		writer := bufio.NewWriter(&b)
		w.WriteTo(writer)

	} else {
		// Use below to save to file
		p.Save(5*vg.Inch, 4*vg.Inch, title+".png")

	}

	return b.Bytes()
}
