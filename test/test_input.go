package main

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"

    ei "github.com/Lattay/ebiten-actions"
)

var not_inited = true

type Game struct {
    actionStream <-chan ei.ActionContainer
    feedbackStream chan<- int
    handler * ei.Handler
}

func (g *Game) Update() error {
    if not_inited {
        g.handler.DetectGamePads()
        not_inited = false
    }
    return nil
}

func (g * Game) Draw(screen * ebiten.Image) {
    for act := range g.actionStream {
        if act.End {
            break
        } else {
            switch act.Data {
            case ei.Jump:
                ebitenutil.DebugPrintAt(screen, "Jump", 10, 10)
            case ei.Crouch:
                ebitenutil.DebugPrintAt(screen, "Crouch", 10, 60)
            case ei.Left:
                ebitenutil.DebugPrintAt(screen, "Left", 10, 110)
            case ei.Right:
                ebitenutil.DebugPrintAt(screen, "Right", 10, 160)
            case ei.ButtonA:
                ebitenutil.DebugPrintAt(screen, "ButtonA", 10, 210)
            case ei.ButtonB:
                ebitenutil.DebugPrintAt(screen, "ButtonB", 10, 260)
            case ei.ButtonX:
                ebitenutil.DebugPrintAt(screen, "ButtonX", 10, 310)
            case ei.ButtonY:
                ebitenutil.DebugPrintAt(screen, "ButtonY", 10, 360)
            case ei.Select:
                ebitenutil.DebugPrintAt(screen, "Select", 10, 410)
            case ei.Start:
                ebitenutil.DebugPrintAt(screen, "Start", 10, 460)
            default:
                ebitenutil.DebugPrintAt(screen, "?????", 300, 240)
            }
        }
    }
    g.feedbackStream <- 0
}

func (g * Game) Layout(outsideWidth, outsideHeight int) (screeWidth, screenHeight int) {
    return 640, 480
}

func main() {
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Test Input")

    actionStream := make(chan ei.ActionContainer, 32)
    feedbackStream := make(chan int, 2)
    h := ei.MakePlatformerHandler()
    game := &Game {
        actionStream,
        feedbackStream,
        h,
    }

    go h.HandleEvents(actionStream, feedbackStream)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
