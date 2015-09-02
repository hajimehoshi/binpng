// Copyright 2014 Hajime Hoshi
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

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	palette := color.Palette([]color.Color{
		color.Transparent, color.Opaque,
	})
	p := image.NewPaletted(img.Bounds(), palette)
	draw.Draw(p, p.Bounds(), img, image.ZP, draw.Src)
	e := png.Encoder{CompressionLevel: png.BestCompression}
	if err := e.Encode(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}
