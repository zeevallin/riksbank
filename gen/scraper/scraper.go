package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	seriesURL = "https://www.riksbank.se/sv/statistik/sok-rantor--valutakurser/oppet-api/serier-for-webbservices/"
)

// Sections is a slice of Section
type Sections []*Section

// Section represents a grouping of groups
type Section struct {
	Description string
	Groups      []*Group
}

// Group represents a group of series
type Group struct {
	Description string
	ID          int
	Series      []*Series
}

// Series represents a single series
type Series struct {
	GroupID     int
	ID          string
	Description string
}

// Scrape fetches the series from the riksbanks website
func Scrape() (sections Sections) {
	c := colly.NewCollector()

	var sectionsC = make(chan Sections, 1)

	c.OnHTML("div.page-base__main__body", func(e *colly.HTMLElement) {
		var (
			sections       = make(Sections, 0)
			currentSection *Section
		)

		e.ForEach(`.page-base__main__body > h2, .page-base__main__body > table`, func(n int, ce *colly.HTMLElement) {
			switch ce.Name {
			case "h2":
				if strings.TrimSpace(ce.Text) != "" {
					currentSection = &Section{
						Description: strings.TrimSpace(ce.Text),
						Groups:      make([]*Group, 0),
					}
					sections = append(sections, currentSection)
				}
			case "table":
				var currentGroup *Group

				ce.ForEach(`caption, tr`, func(n int, te *colly.HTMLElement) {
					switch te.Name {
					case "caption":
						currentGroup = &Group{
							Description: strings.TrimSpace(te.Text),
							Series:      make([]*Series, 0),
						}
						currentSection.Groups = append(currentSection.Groups, currentGroup)
					case "tr":
						if te.Index > 1 {
							s := &Series{}
							te.ForEach(`td`, func(n int, tde *colly.HTMLElement) {
								switch n {
								case 0:
									gid, _ := strconv.Atoi(strings.TrimSpace(tde.Text))
									s.GroupID = gid
									currentGroup.ID = gid
								case 1:
									s.ID = strings.TrimSpace(tde.Text)
								case 2:
									s.Description = strings.TrimSpace(tde.Text)
								}
							})
							currentGroup.Series = append(currentGroup.Series, s)
						}
					}
				})
			}
		})

		// printSections(sections)
		sectionsC <- sections

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping", r.URL)
	})

	c.Visit(seriesURL)

	return <-sectionsC
}

func printSections(sections Sections) {
	for _, section := range sections {
		fmt.Println(section.Description)
		for _, group := range section.Groups {
			fmt.Println("\t", group.Description)
			for _, series := range group.Series {
				fmt.Println("\t\t", series.GroupID, series.ID, series.Description)
			}
		}
	}
}
