package input
import (
    "math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Handler struct {
    mappings []Mapping
}

type Device int

const (
    Keyboard Device = iota
    MouseButton
    MouseWheel
    MousePointer
    Gamepad
)

type Action interface {}
type KeyInput interface {}

type MouseWheelInput struct {
    pos bool
    x bool
}

type GamepadButton struct {
    id ebiten.GamepadID
    button ebiten.GamepadButton
}

type GamepadAxis struct {
    id ebiten.GamepadID
    axis int
    pos bool
    trigger float64
}

type Mapping struct {
    key KeyInput
    action Action
}

type ActionContainer struct {
    End bool
    Data Action
}

func (h * Handler) AddMapping(m Mapping) {
    h.mappings = append(h.mappings, m)
}

func (h * Handler) ClearAction(a Action) {
    new_m := make([]Mapping, 0, len(h.mappings))
    for _, m := range h.mappings {
        if m.action != a {
            new_m = append(new_m, m)
        }
    }
    h.mappings = new_m
}

func (h * Handler) NumMappings() int {
    return len(h.mappings)
}

func (h * Handler) HandleEvents(c chan<- ActionContainer, user <-chan int) {
    c <- ActionContainer{true, nil}
    for {
        select {
        case <-user:
            c <- ActionContainer{true, nil}
            continue
        default:
        }
        for _, m := range h.mappings {
            active := false
            switch k := m.key.(type) {
            case ebiten.MouseButton:
                active = ebiten.IsMouseButtonPressed(k);
            case MouseWheelInput:
                xoff, yoff := ebiten.Wheel()
                active = (k.x && ((k.pos && xoff > 0) || (!k.pos && xoff < 0))) ||
                        (!k.x && ((k.pos && yoff > 0) || (!k.pos && yoff < 0)))
            case GamepadAxis:
                delta := ebiten.GamepadAxis(k.id, k.axis)
                active = (k.pos && delta > math.Abs(k.trigger)) || (!k.pos && delta < -math.Abs(k.trigger))
            case GamepadButton:
                active = ebiten.IsGamepadButtonPressed(k.id, k.button)

            case ebiten.Key:
                active = ebiten.IsKeyPressed(k)

            default: panic(0)
            }
            if active {
                c <- ActionContainer{false, m.action}
            }
        }
    }
}
