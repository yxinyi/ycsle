// Copyright 2016 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twenty48

import (
	"client/user"
	"csmsg"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/name5566/leaf/log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"image/color"

	"math/rand"
	"time"
)

var (
	uiFont        font.Face
	uiFontMHeight int
)

func init() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err.Error())
	}
	uiFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	b, _, _ := uiFont.GlyphBounds('M')
	uiFontMHeight = (b.Max.Y - b.Min.Y).Ceil()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

// Game represents a game state.
type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
	Count      uint32
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
	}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	detailStr := fmt.Sprintf("%d", g.Count)
	text.Draw(screen, detailStr, uiFont, 100, 100, color.White)

	if user.U != nil {
		detailStr = fmt.Sprintf("%d", user.U.LoginState)
		text.Draw(screen, detailStr, uiFont, 100, 200, color.White)
	}

	/*	if g.boardImage == nil {
			w, h := g.board.Size()
			g.boardImage = ebiten.NewImage(w, h)
		}
		screen.Fill(backgroundColor)
		g.board.Draw(g.boardImage)
		op := &ebiten.DrawImageOptions{}
		sw, sh := screen.Size()
		bw, bh := g.boardImage.Size()
		x := (sw - bw) / 2
		y := (sh - bh) / 2
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(g.boardImage, op)

	*/

}

func (g *Game) EventRegister() {

	csmsg.Processor.SetHandler(&csmsg.S2C_EnterSquare{}, func(list []interface{}) {
		/*		for _, it := range list {
				log.Debug("name [%v] type [%v]", reflect.TypeOf(it).Elem().Name(), reflect.TypeOf(it).Kind().String())
			}*/
		msg := list[0].(*csmsg.S2C_EnterSquare)
		//a := list[1].(*agent.Agent)
		//log.Debug("msg Count [%v]", msg.Count)
		g.Count = msg.TotalCount
	})

}
