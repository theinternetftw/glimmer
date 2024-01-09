package glimmer

import (
	"sync"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// WindowState contains what the window loop and program proper both need to touch
type WindowState struct {
	// the Width of the framebuffer
	RenderWidth int
	// the Height of the framebuffer
	RenderHeight int

	// a Mutex that must be held when reading or writing gfx in WindowState
	RenderMutex sync.Mutex

	// a Mutex that must be held when reading or writing input in WindowState
	InputMutex sync.Mutex

	// Pix is the raw RGBA bytes of the framebuffer. The RenderMutex must be held when touching it.
	Pix []byte

    keyBuf []ebiten.Key
	keyCodeArray [256]bool
	keyCodeMap   map[KeyCode]bool
	keyCharArray [256]bool
	keyCharMap   map[rune]bool

    updateCallback func(*WindowState)
    renderCallback func(*WindowState)
}

// CopyKeyCharArray writes the current ascii keystate to dest
func (s *WindowState) CopyKeyCharArray(dest []bool) {
	copy(dest, s.keyCharArray[:])
}

// CharIsDown returns the key state for that char
func (s *WindowState) CharIsDown(c rune) bool {
	if c >= 0 && c < 256 {
		return s.keyCharArray[byte(c)]
	}
	return s.keyCharMap[c]
}

// CodeIsDown returns the key state for that keyCode
func (s *WindowState) CodeIsDown(c KeyCode) bool {
	if c < 256 {
		return s.keyCodeArray[byte(c)]
	}
	return s.keyCodeMap[c]
}

func (s *WindowState) updateKey(key KeyCode, isPressed bool) {
    if key < 256 {
        s.keyCodeArray[byte(key)] = isPressed
    } else {
        s.keyCodeMap[key] = isPressed
    }

    sMap := KeyCodeToUnshiftedAsciiMap
    if s.CodeIsDown(KeyCodeShiftLeft) || s.CodeIsDown(KeyCodeShiftRight) {
        sMap = KeyCodeToShiftedAsciiMap
    }
    if v, ok := sMap[key]; ok {
        s.keyCharArray[v] = isPressed
    }
}
func (s *WindowState) updateKeyboardState() {
    s.keyBuf = s.keyBuf[:0]
    s.keyBuf = inpututil.AppendJustReleasedKeys(s.keyBuf)

    for _, key := range s.keyBuf {
        s.updateKey(KeyCode(key), false)
    }

    s.keyBuf = s.keyBuf[:0]
    s.keyBuf = inpututil.AppendJustPressedKeys(s.keyBuf)

    for _, key := range s.keyBuf {
        s.updateKey(KeyCode(key), true)
    }
}

type Game struct {
    window *WindowState
    renderWidth int
    renderHeight int
    windowWidth int
    windowHeight int
}
func (g *Game) Update() error {
    g.window.updateKeyboardState()
    if g.window.updateCallback != nil {
        g.window.updateCallback(g.window)
    }
    return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
    if g.window.renderCallback != nil {
        g.window.renderCallback(g.window)
    }
    g.window.RenderMutex.Lock()
    screen.WritePixels(g.window.Pix)
    g.window.RenderMutex.Unlock()
}
func (g *Game) Layout(newWindowWidth, newWindowHeight int) (int, int) {
    g.windowWidth = newWindowWidth
    g.windowHeight = newWindowHeight
    return g.renderWidth, g.renderHeight
}

type InitDisplayLoopOptions struct {
    WindowTitle string
    WindowWidth int
    WindowHeight int
    RenderWidth int
    RenderHeight int
    InitCallback func(*WindowState)
    UpdateCallback func(*WindowState)
    RenderCallback func(*WindowState)
}

// InitDisplayLoop creates a window and starts event loop
func InitDisplayLoop(opts InitDisplayLoopOptions) {

    windowState := WindowState{
        RenderWidth:      opts.RenderWidth,
        RenderHeight:     opts.RenderHeight,
        Pix:        make([]byte, 4*opts.RenderWidth*opts.RenderHeight),
        keyCodeMap: map[KeyCode]bool{},
        keyCharMap: map[rune]bool{},
        updateCallback: opts.UpdateCallback,
        renderCallback: opts.RenderCallback,
    }

    if opts.InitCallback != nil {
        go opts.InitCallback(&windowState)
    }

    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    ebiten.SetWindowSize(opts.WindowWidth, opts.WindowHeight)
    ebiten.SetWindowTitle(opts.WindowTitle)
    game := &Game{
        window: &windowState,
        windowWidth: opts.WindowWidth, 
        windowHeight: opts.WindowHeight, 
        renderWidth: opts.RenderWidth, 
        renderHeight: opts.RenderHeight, 
    }
    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
