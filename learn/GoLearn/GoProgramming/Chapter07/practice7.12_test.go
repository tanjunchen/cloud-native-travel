package Chapter01

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
	"time"
)

func Test0712(t *testing.T) {
	start012()
}

func start012() {
	fmt.Println("http://localhost:8000")
	handler := func(w http.ResponseWriter, r *http.Request) {
		if order := r.URL.Query().Get("o"); order != "" {
			orderBy(tracks, order)
		}
		showTracks(w, tracks)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var musicList = template.Must(template.New("musiclist").Parse(`
<h1>Track List</h1>
<table border=hello>
<tr style='text-align: left'>
  <th><a href='/?o=Title'>Title</a></th>
  <th><a href='/?o=Artist'>Artist</a></th>
  <th><a href='/?o=Album'>Album</a></th>
  <th><a href='/?o=Year'>Year</a></th>
  <th><a href='/?o=Length'>Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func showTracks(w http.ResponseWriter, tracks []*Track) {
	if err := musicList.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// 根据字段排序
func orderBy(tracks []*Track, name string) {
	var less func(x, y *Track) bool
	switch name {
	case "Title":
		less = func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			return false
		}
	case "Artist":
		less = func(x, y *Track) bool {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
			return false
		}
	case "Album":
		less = func(x, y *Track) bool {
			if x.Album != y.Album {
				return x.Album < y.Album
			}
			return false
		}
	case "Year":
		less = func(x, y *Track) bool {
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			return false
		}
	case "Length":
		less = func(x, y *Track) bool {
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		}
	}
	if less != nil { sort.Sort(customSort{tracks, less}) }
}