package render

import (
	"image"
	"math"

	"github.com/weqqr/panorama/game"
	"github.com/weqqr/panorama/world"
)

func overlayWithDepth(target *image.NRGBA, targetDepth *DepthBuffer, source *image.NRGBA, sourceDepth *DepthBuffer, origin image.Point, depthOffset float32) {
	width := source.Rect.Dx()
	height := source.Rect.Dy()

	for y := origin.Y; y < origin.Y+height; y++ {
		for x := origin.X; x < origin.X+width; x++ {
			targetZ := targetDepth.At(x, y)
			sourceZ := sourceDepth.At(x-origin.X, y-origin.Y) + depthOffset

			if sourceZ > targetZ {
				continue
			}

			targetDepth.Set(x, y, sourceZ)

			c := source.NRGBAAt(x-origin.X, y-origin.Y)
			if c.A == 0 {
				// TODO: support opacity
				continue
			}
			target.SetNRGBA(x, y, c)
		}
	}
}

func RenderBlock(nr *NodeRasterizer, block *world.MapBlock, game *game.Game) (*image.NRGBA, *DepthBuffer) {
	rect := image.Rect(0, 0, TileBlockWidth, TileBlockHeight)
	blockColor := image.NewNRGBA(rect)
	blockDepth := NewDepthBuffer(rect)

	// FIXME: nodes must define their origin points
	originX, originY := rect.Dx()/2-BaseResolution/2, rect.Dy()/2+BaseResolution/4+2

	for z := 0; z < world.MapBlockSize; z++ {
		for y := 0; y < world.MapBlockSize; y++ {
			for x := 0; x < world.MapBlockSize; x++ {
				node := block.GetNode(x, y, z)
				nodeName := block.ResolveName(node.ID)
				gameNode := game.Node(nodeName)

				nodeColor, nodeDepth := nr.Render(nodeName, &gameNode)

				tileOffsetX := originX + BaseResolution*(z-x)/2
				tileOffsetY := originY + BaseResolution/4*(z+x) - YOffsetCoef*y

				depthOffset := -float32(z+x)/math.Sqrt2 - 0.5*(float32(y))
				overlayWithDepth(blockColor, blockDepth, nodeColor, nodeDepth, image.Pt(tileOffsetX, tileOffsetY), depthOffset)
			}
		}
	}

	return blockColor, blockDepth
}
