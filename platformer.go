package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)



type ActionPlatformer int

const (
    Jump ActionPlatformer = iota
    Left
    Right
    Crouch
    // XBox style:
    ButtonA
    ButtonB
    ButtonX
    ButtonY
    Select
    Start
)

func MakePlatformerHandler() * Handler {
    h := Handler{make([]Mapping, 0, 24)}
    // classic
    h.AddMapping(Mapping{ebiten.KeyW, Jump})
    h.AddMapping(Mapping{ebiten.KeyA, Left})
    h.AddMapping(Mapping{ebiten.KeyD, Right})
    h.AddMapping(Mapping{ebiten.KeyS, Crouch})
    // inverted
    h.AddMapping(Mapping{ebiten.KeyUp, Jump})
    h.AddMapping(Mapping{ebiten.KeyLeft, Left})
    h.AddMapping(Mapping{ebiten.KeyRight, Right})
    h.AddMapping(Mapping{ebiten.KeyDown, Crouch})

    h.AddMapping(Mapping{ebiten.KeyJ, ButtonA})
    h.AddMapping(Mapping{ebiten.KeyK, ButtonB})
    h.AddMapping(Mapping{ebiten.KeySpace, ButtonX})
    h.AddMapping(Mapping{ebiten.KeyE, ButtonY})
    h.AddMapping(Mapping{ebiten.KeyEnter, Start})
    h.AddMapping(Mapping{ebiten.KeyBackspace, Select})

    return &h
}

func (h * Handler) DetectGamePads() {
    for _, id := range ebiten.GamepadIDs() {
        h.AddMapping(Mapping{GamepadButton{id, 0}, ButtonA})
        h.AddMapping(Mapping{GamepadButton{id, 1}, ButtonB})
        h.AddMapping(Mapping{GamepadButton{id, 2}, ButtonX})
        h.AddMapping(Mapping{GamepadButton{id, 3}, ButtonY})
        h.AddMapping(Mapping{GamepadButton{id, 6}, Select})
        h.AddMapping(Mapping{GamepadButton{id, 7}, Start})
        h.AddMapping(Mapping{GamepadAxis{id, 0, true, 0.3}, Right})
        h.AddMapping(Mapping{GamepadAxis{id, 0, false, 0.3}, Left})
        h.AddMapping(Mapping{GamepadAxis{id, 1, false, 0.3}, Jump})
        h.AddMapping(Mapping{GamepadAxis{id, 1, true, 0.3}, Crouch})
    }
}
