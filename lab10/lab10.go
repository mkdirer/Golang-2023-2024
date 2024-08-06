package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func GetLinks(url string) ([]string, error) {
	links := make([]string, 0)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return links, err
	}
	body := html.NewTokenizer(res.Body)
	for {
		nextBody := body.Next()
		switch nextBody {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken:
			token := body.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func getLinksToOtherPages(url string) ([]string, error) {
	links := make([]string, 0)
	n := 0
	res, _ := GetLinks(url) // ignore errors
	for _, v := range res {
		// limit number of pages to 10
		if strings.HasPrefix(v, "https://") && n < 10 {
			links = append(links, v)
			n += 1
		}
	}
	return links, nil
}

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

func BuildGraph(url string, n int) error {
	links := []string{url}
	var palette = []color.Color{color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}
	size := 100*(n+4) + 20
	delay := 20
	smallR := 2
	bigR := 100.0
	var images []*image.Paletted
	var delays []int
	x0 := size / 2
	y0 := size / 2
	rect := image.Rect(0, 0, size, size)
	img := image.NewPaletted(rect, palette)
	drawCircle(img, x0, y0, 10, palette[1])
	delays = append(delays, delay)
	images = append(images, img)

	for i := 0; i <= n; i++ {
		newImg := image.NewPaletted(image.Rect(0, 0, size, size), palette)
		fmt.Println("...Creating frame", i)
		// get data
		currentTotalLinks := make([]string, 0)
		for _, link := range links {
			currentLinks, err := getLinksToOtherPages(link)
			if err != nil {
				return err
			}
			currentTotalLinks = append(currentTotalLinks, currentLinks...)
		}
		currentNumber := len(currentTotalLinks)
		// create frame
		for j := 0; j < currentNumber; j++ {
			xj := x0 + int(math.Round(bigR*math.Cos(float64(2*j)*math.Pi/float64(currentNumber))))
			yj := y0 + int(math.Round(bigR*math.Sin(float64(2*j)*math.Pi/float64(currentNumber))))
			drawCircle(newImg, xj, yj, smallR, palette[2])
		}
		delays = append(delays, delay)
		images = append(images, newImg)
		bigR += 100
		links = currentTotalLinks
	}
	f, err := os.OpenFile("graph.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	gif.EncodeAll(f, &gif.GIF{Image: images, Delay: delays})
	return nil
}

func main() {
	// 1
	res, err := GetLinks("http://go.dev")
	if err != nil {
		fmt.Println("Error while getting links", err)
		return
	}
	fmt.Println("Found urls:")
	for _, v := range res {
		fmt.Printf("\t%v\n", v)
	}

	// 2 for example: go run lab10.go http://go.dev
	if len(os.Args) > 1 {
		fmt.Println("Creating gif graph.gif")
		url := os.Args[1]
		n := 2 // depth
		err := BuildGraph(url, n)
		if err != nil {
			fmt.Println("Error while creating graph", err)
			return
		}
		fmt.Println("Created gif graph.gif")
	}
}
