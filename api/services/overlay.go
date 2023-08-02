package service

import (
	"database/sql"
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/golang/freetype/truetype"
	"github.com/google/renameio"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type OverlayContext struct {
	bg         *image.RGBA
	height     int
	width      int
	fontFace   font.Face
	fontHeight float64
}

type OverlayObject struct {
	Id         string  `json:"id"`
	CoverId    string  `json:"coverId"`
	Type       string  `json:"type"`
	Text       string  `json:"text"`
	FontFamily string  `json:"fontFamily"`
	FontSize   int     `json:"fontSize"`
	LineHeight float64 `json:"lineHeight"`
	TextColor  string  `json:"textColor"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	PointX     int     `json:"pointX"`
	PointY     int     `json:"pointY"`
	Show       bool    `json:"show"`
	TextAlign  string  `json:"textAlign"`
}

func (s *MainServer) changeSongOverlay(send bool) {
	resWidth := 1280
	resHeight := 720

	bg := image.NewRGBA(image.Rect(0, 0, resWidth, resHeight))

	c := NewOverlay(bg)

	// draw text and cover on image.
	for i := 0; i < len(s.overlays); i++ {
		switch s.overlays[i].Id {
		case "cover":
			if s.overlays[i].Show {
				s.overlays[i].CoverId = s.currentSong.Cover
				err := c.DrawImage(s.overlays[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
		case "song_name":
			if s.overlays[i].Show {
				s.overlays[i].Text = s.currentSong.Name
				err := c.DrawText(s.overlays[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
		case "song_author":
			if s.overlays[i].Show {
				s.overlays[i].Text = s.currentSong.Author
				err := c.DrawText(s.overlays[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
		case "song_page":
			if s.overlays[i].Show {
				s.overlays[i].Text = s.currentSong.Page
				err := c.DrawText(s.overlays[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	b, err := json.Marshal(&DataChannelMsg{Type: "overlay", Message: s.overlays})
	if err != nil {
		log.Fatalln(err)
	}

	if send {
		SendChannelData <- string(b)
	}

	// Save the new image to a file
	output, err := os.Create("files/stream/next.png")
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer output.Close()

	if err := png.Encode(output, bg); err != nil {
		log.Fatalln("Error encoding file:", err)
	}

	// Atomically copy new image to stream.png
	overlay, err := os.ReadFile("files/stream/next.png")
	if err != nil {
		log.Fatalln(err)
	}

	renameio.WriteFile("files/stream/stream.png", overlay, 0644)

	output.Close()
}

func (s *MainServer) changeOverlay(object OverlayObject) {
	for i := 0; i < len(s.overlays); i++ {
		if s.overlays[i].Id == object.Id {
			s.overlays[i] = object
		}
	}

	s.changeSongOverlay(false)
}

func (s *MainServer) InitOverlay() error {
	if len(s.overlays) > 0 {
		s.overlays = nil
	}

	// Get cover overlay settings
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}

	var coverSettings OverlayObject
	err = db.QueryRow("SELECT id, type, width, height, point_x, point_y, show FROM overlays WHERE id='cover'").Scan(
		&coverSettings.Id,
		&coverSettings.Type,
		&coverSettings.Width,
		&coverSettings.Height,
		&coverSettings.PointX,
		&coverSettings.PointY,
		&coverSettings.Show,
	)
	if err != nil {
		return err
	}

	s.overlays = append(s.overlays, coverSettings)

	var songNameSettings OverlayObject
	err = db.QueryRow("SELECT id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align FROM overlays WHERE id='song_name'").Scan(
		&songNameSettings.Id,
		&songNameSettings.Type,
		&songNameSettings.Width,
		&songNameSettings.Height,
		&songNameSettings.PointX,
		&songNameSettings.PointY,
		&songNameSettings.Show,
		&songNameSettings.FontFamily,
		&songNameSettings.FontSize,
		&songNameSettings.LineHeight,
		&songNameSettings.TextColor,
		&songNameSettings.TextAlign,
	)
	if err != nil {
		return err
	}

	s.overlays = append(s.overlays, songNameSettings)

	var songAuthorSettings OverlayObject
	err = db.QueryRow("SELECT id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align FROM overlays WHERE id='song_author'").Scan(
		&songAuthorSettings.Id,
		&songAuthorSettings.Type,
		&songAuthorSettings.Width,
		&songAuthorSettings.Height,
		&songAuthorSettings.PointX,
		&songAuthorSettings.PointY,
		&songAuthorSettings.Show,
		&songAuthorSettings.FontFamily,
		&songAuthorSettings.FontSize,
		&songAuthorSettings.LineHeight,
		&songAuthorSettings.TextColor,
		&songAuthorSettings.TextAlign,
	)
	if err != nil {
		return err
	}

	s.overlays = append(s.overlays, songAuthorSettings)

	var songPageSettings OverlayObject
	err = db.QueryRow("SELECT id, type, width, height, point_x, point_y, show, font_family, font_size, line_height, text_color, text_align FROM overlays WHERE id='song_page'").Scan(
		&songPageSettings.Id,
		&songPageSettings.Type,
		&songPageSettings.Width,
		&songPageSettings.Height,
		&songPageSettings.PointX,
		&songPageSettings.PointY,
		&songPageSettings.Show,
		&songPageSettings.FontFamily,
		&songPageSettings.FontSize,
		&songPageSettings.LineHeight,
		&songPageSettings.TextColor,
		&songPageSettings.TextAlign,
	)
	if err != nil {
		return err
	}

	s.overlays = append(s.overlays, songPageSettings)

	return nil
}

func NewOverlay(im *image.RGBA) *OverlayContext {
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	return &OverlayContext{
		bg:         im,
		width:      w,
		height:     h,
		fontFace:   basicfont.Face7x13,
		fontHeight: 13,
	}
}

func (dc *OverlayContext) SetFontFace(fontFace font.Face) {
	dc.fontFace = fontFace
	dc.fontHeight = float64(fontFace.Metrics().Ascent.Ceil()) - float64(fontFace.Metrics().Descent.Ceil()) + 1
}

func (c *OverlayContext) DrawImage(overlay OverlayObject) error {
	// Draw cover
	// Open the original image
	file, err := os.Open("files/covers/" + overlay.CoverId)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// Resize image to specified width and height.
	coverResized := resize.Resize(uint(overlay.Width), uint(overlay.Height), img, resize.Lanczos2)
	draw.Draw(c.bg, c.bg.Bounds(), coverResized, image.Pt(-overlay.PointX, -overlay.PointY), draw.Src)

	return nil
}

func (c *OverlayContext) DrawText(overlay OverlayObject) error {
	fontBytes, err := os.ReadFile(overlay.FontFamily)
	if err != nil {
		log.Println("Error opening font file: ", err)
		return err
	}

	// Parse font
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println("Error parsing font: ", err)
		return err
	}

	// Set the font options
	fontOptions := &truetype.Options{
		Size: float64(overlay.FontSize),
	}

	c.SetFontFace(truetype.NewFace(f, fontOptions))

	lineSpacing := fontOptions.Size * overlay.LineHeight

	c.DrawStringWrapped(overlay.Text, overlay.TextAlign, overlay.TextColor, float64(overlay.PointX), float64(overlay.PointY), float64(overlay.Width), lineSpacing)

	return nil
}

func splitOnSpace(x string) []string {
	var result []string
	pi := 0
	ps := false
	for i, c := range x {
		s := unicode.IsSpace(c)
		if s != ps && i > 0 {
			result = append(result, x[pi:i])
			pi = i
		}
		ps = s
	}
	result = append(result, x[pi:])
	return result
}

// MeasureString returns the rendered width and height of the specified text
// given the current font face.
func (c *OverlayContext) measureString(s string) (w, h float64) {
	d := &font.Drawer{
		Face: c.fontFace,
	}
	a := d.MeasureString(s)
	return float64(a >> 6), c.fontHeight
}

func (c *OverlayContext) wordWrap(s string, width float64) []string {
	var result []string
	for _, line := range strings.Split(s, "\n") {
		fields := splitOnSpace(line)
		if len(fields)%2 == 1 {
			fields = append(fields, "")
		}
		x := ""
		for i := 0; i < len(fields); i += 2 {
			w, _ := c.measureString(x + fields[i])
			if w > width {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i] + fields[i+1]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func fixp(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{fix(x), fix(y)}
}

// DrawStringWrapped word-wraps the specified string to the given max width
// and then draws it at the specified anchor point using the given line
// spacing and text alignment.
func (c *OverlayContext) DrawStringWrapped(s, align, textColor string, x, y, width, lineSpacing float64) {
	lines := c.wordWrap(s, width)

	y += c.fontHeight

	for _, line := range lines {
		var textX float64

		w, _ := c.measureString(line)

		switch align {
		case "left":
			textX = x
		case "center":
			textX = (width / 2) + x
			textX -= 0.5 * w
		case "right":
			textX = x
			textX += width - w
		}

		RGB := strings.Split(textColor, " ")

		r, _ := strconv.Atoi(RGB[0])
		g, _ := strconv.Atoi(RGB[1])
		b, _ := strconv.Atoi(RGB[2])

		d := &font.Drawer{
			Dst:  c.bg,
			Src:  image.NewUniform(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}),
			Face: c.fontFace,
			Dot:  fixp(textX, y),
		}
		d.DrawString(line)
		y += c.fontHeight + lineSpacing
	}
}
