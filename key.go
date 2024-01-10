package glimmer

import "github.com/hajimehoshi/ebiten/v2"

type KeyCode int

const (
	KeyCodeUnknown KeyCode = -1

	KeyCodeA            KeyCode = KeyCode(ebiten.KeyA)
	KeyCodeB            KeyCode = KeyCode(ebiten.KeyB)
	KeyCodeC            KeyCode = KeyCode(ebiten.KeyC)
	KeyCodeD            KeyCode = KeyCode(ebiten.KeyD)
	KeyCodeE            KeyCode = KeyCode(ebiten.KeyE)
	KeyCodeF            KeyCode = KeyCode(ebiten.KeyF)
	KeyCodeG            KeyCode = KeyCode(ebiten.KeyG)
	KeyCodeH            KeyCode = KeyCode(ebiten.KeyH)
	KeyCodeI            KeyCode = KeyCode(ebiten.KeyI)
	KeyCodeJ            KeyCode = KeyCode(ebiten.KeyJ)
	KeyCodeK            KeyCode = KeyCode(ebiten.KeyK)
	KeyCodeL            KeyCode = KeyCode(ebiten.KeyL)
	KeyCodeM            KeyCode = KeyCode(ebiten.KeyM)
	KeyCodeN            KeyCode = KeyCode(ebiten.KeyN)
	KeyCodeO            KeyCode = KeyCode(ebiten.KeyO)
	KeyCodeP            KeyCode = KeyCode(ebiten.KeyP)
	KeyCodeQ            KeyCode = KeyCode(ebiten.KeyQ)
	KeyCodeR            KeyCode = KeyCode(ebiten.KeyR)
	KeyCodeS            KeyCode = KeyCode(ebiten.KeyS)
	KeyCodeT            KeyCode = KeyCode(ebiten.KeyT)
	KeyCodeU            KeyCode = KeyCode(ebiten.KeyU)
	KeyCodeV            KeyCode = KeyCode(ebiten.KeyV)
	KeyCodeW            KeyCode = KeyCode(ebiten.KeyW)
	KeyCodeX            KeyCode = KeyCode(ebiten.KeyX)
	KeyCodeY            KeyCode = KeyCode(ebiten.KeyY)
	KeyCodeZ            KeyCode = KeyCode(ebiten.KeyZ)
	KeyCodeAltLeft      KeyCode = KeyCode(ebiten.KeyAltLeft)
	KeyCodeAltRight     KeyCode = KeyCode(ebiten.KeyAltRight)
	KeyCodeArrowDown    KeyCode = KeyCode(ebiten.KeyArrowDown)
	KeyCodeArrowLeft    KeyCode = KeyCode(ebiten.KeyArrowLeft)
	KeyCodeArrowRight   KeyCode = KeyCode(ebiten.KeyArrowRight)
	KeyCodeArrowUp      KeyCode = KeyCode(ebiten.KeyArrowUp)
	KeyCodeBackquote    KeyCode = KeyCode(ebiten.KeyBackquote)
	KeyCodeBackslash    KeyCode = KeyCode(ebiten.KeyBackslash)
	KeyCodeBackspace    KeyCode = KeyCode(ebiten.KeyBackspace)
	KeyCodeBracketLeft  KeyCode = KeyCode(ebiten.KeyBracketLeft)
	KeyCodeBracketRight KeyCode = KeyCode(ebiten.KeyBracketRight)
	KeyCodeCapsLock     KeyCode = KeyCode(ebiten.KeyCapsLock)
	KeyCodeComma        KeyCode = KeyCode(ebiten.KeyComma)
	KeyCodeContextMenu  KeyCode = KeyCode(ebiten.KeyContextMenu)
	KeyCodeControlLeft  KeyCode = KeyCode(ebiten.KeyControlLeft)
	KeyCodeControlRight KeyCode = KeyCode(ebiten.KeyControlRight)
	KeyCodeDelete       KeyCode = KeyCode(ebiten.KeyDelete)
	KeyCodeDigit0       KeyCode = KeyCode(ebiten.KeyDigit0)
	KeyCodeDigit1       KeyCode = KeyCode(ebiten.KeyDigit1)
	KeyCodeDigit2       KeyCode = KeyCode(ebiten.KeyDigit2)
	KeyCodeDigit3       KeyCode = KeyCode(ebiten.KeyDigit3)
	KeyCodeDigit4       KeyCode = KeyCode(ebiten.KeyDigit4)
	KeyCodeDigit5       KeyCode = KeyCode(ebiten.KeyDigit5)
	KeyCodeDigit6       KeyCode = KeyCode(ebiten.KeyDigit6)
	KeyCodeDigit7       KeyCode = KeyCode(ebiten.KeyDigit7)
	KeyCodeDigit8       KeyCode = KeyCode(ebiten.KeyDigit8)
	KeyCodeDigit9       KeyCode = KeyCode(ebiten.KeyDigit9)
	KeyCodeEnd          KeyCode = KeyCode(ebiten.KeyEnd)
	KeyCodeEnter        KeyCode = KeyCode(ebiten.KeyEnter)
	KeyCodeEqual        KeyCode = KeyCode(ebiten.KeyEqual)
	KeyCodeEscape       KeyCode = KeyCode(ebiten.KeyEscape)
	KeyCodeF1           KeyCode = KeyCode(ebiten.KeyF1)
	KeyCodeF2           KeyCode = KeyCode(ebiten.KeyF2)
	KeyCodeF3           KeyCode = KeyCode(ebiten.KeyF3)
	KeyCodeF4           KeyCode = KeyCode(ebiten.KeyF4)
	KeyCodeF5           KeyCode = KeyCode(ebiten.KeyF5)
	KeyCodeF6           KeyCode = KeyCode(ebiten.KeyF6)
	KeyCodeF7           KeyCode = KeyCode(ebiten.KeyF7)
	KeyCodeF8           KeyCode = KeyCode(ebiten.KeyF8)
	KeyCodeF9           KeyCode = KeyCode(ebiten.KeyF9)
	KeyCodeF10          KeyCode = KeyCode(ebiten.KeyF10)
	KeyCodeF11          KeyCode = KeyCode(ebiten.KeyF11)
	KeyCodeF12          KeyCode = KeyCode(ebiten.KeyF12)
	KeyCodeHome         KeyCode = KeyCode(ebiten.KeyHome)
	KeyCodeInsert       KeyCode = KeyCode(ebiten.KeyInsert)
	KeyCodeMetaLeft     KeyCode = KeyCode(ebiten.KeyMetaLeft)
	KeyCodeMetaRight    KeyCode = KeyCode(ebiten.KeyMetaRight)
	KeyCodeMinus        KeyCode = KeyCode(ebiten.KeyMinus)
	KeyCodePageDown     KeyCode = KeyCode(ebiten.KeyPageDown)
	KeyCodePageUp       KeyCode = KeyCode(ebiten.KeyPageUp)
	KeyCodePause        KeyCode = KeyCode(ebiten.KeyPause)
	KeyCodePeriod       KeyCode = KeyCode(ebiten.KeyPeriod)
	KeyCodePrintScreen  KeyCode = KeyCode(ebiten.KeyPrintScreen)
	KeyCodeQuote        KeyCode = KeyCode(ebiten.KeyQuote)
	KeyCodeScrollLock   KeyCode = KeyCode(ebiten.KeyScrollLock)
	KeyCodeSemicolon    KeyCode = KeyCode(ebiten.KeySemicolon)
	KeyCodeShiftLeft    KeyCode = KeyCode(ebiten.KeyShiftLeft)
	KeyCodeShiftRight   KeyCode = KeyCode(ebiten.KeyShiftRight)
	KeyCodeSlash        KeyCode = KeyCode(ebiten.KeySlash)
	KeyCodeSpace        KeyCode = KeyCode(ebiten.KeySpace)
	KeyCodeTab          KeyCode = KeyCode(ebiten.KeyTab)
	KeyCodeAlt          KeyCode = KeyCode(ebiten.KeyAlt)
	KeyCodeControl      KeyCode = KeyCode(ebiten.KeyControl)
	KeyCodeShift        KeyCode = KeyCode(ebiten.KeyShift)
	KeyCodeMeta         KeyCode = KeyCode(ebiten.KeyMeta)
)

var KeyCodeToUnshiftedAsciiMap = map[KeyCode]rune{
	KeyCodeA:            'a',
	KeyCodeB:            'b',
	KeyCodeC:            'c',
	KeyCodeD:            'd',
	KeyCodeE:            'e',
	KeyCodeF:            'f',
	KeyCodeG:            'g',
	KeyCodeH:            'h',
	KeyCodeI:            'i',
	KeyCodeJ:            'j',
	KeyCodeK:            'k',
	KeyCodeL:            'l',
	KeyCodeM:            'm',
	KeyCodeN:            'n',
	KeyCodeO:            'o',
	KeyCodeP:            'p',
	KeyCodeQ:            'q',
	KeyCodeR:            'r',
	KeyCodeS:            's',
	KeyCodeT:            't',
	KeyCodeU:            'u',
	KeyCodeV:            'v',
	KeyCodeW:            'w',
	KeyCodeX:            'x',
	KeyCodeY:            'y',
	KeyCodeZ:            'z',
	KeyCodeBackquote:    '`',
	KeyCodeBackslash:    '\\',
	KeyCodeBackspace:    '\b',
	KeyCodeBracketLeft:  '[',
	KeyCodeBracketRight: ']',
	KeyCodeComma:        ',',
	KeyCodeDelete:       '\x7f',
	KeyCodeDigit0:       '0',
	KeyCodeDigit1:       '1',
	KeyCodeDigit2:       '2',
	KeyCodeDigit3:       '3',
	KeyCodeDigit4:       '4',
	KeyCodeDigit5:       '5',
	KeyCodeDigit6:       '6',
	KeyCodeDigit7:       '7',
	KeyCodeDigit8:       '8',
	KeyCodeDigit9:       '9',
	KeyCodeEnter:        '\n',
	KeyCodeEqual:        '=',
	KeyCodeEscape:       '\x1b',
	KeyCodeMinus:        '-',
	KeyCodePeriod:       '.',
	KeyCodeQuote:        '\'',
	KeyCodeSemicolon:    ';',
	KeyCodeSlash:        '/',
	KeyCodeSpace:        ' ',
	KeyCodeTab:          '\t',
}

var KeyCodeToShiftedAsciiMap = map[KeyCode]rune{
	KeyCodeA:            'A',
	KeyCodeB:            'B',
	KeyCodeC:            'C',
	KeyCodeD:            'D',
	KeyCodeE:            'E',
	KeyCodeF:            'F',
	KeyCodeG:            'G',
	KeyCodeH:            'H',
	KeyCodeI:            'I',
	KeyCodeJ:            'J',
	KeyCodeK:            'K',
	KeyCodeL:            'L',
	KeyCodeM:            'M',
	KeyCodeN:            'N',
	KeyCodeO:            'O',
	KeyCodeP:            'P',
	KeyCodeQ:            'Q',
	KeyCodeR:            'R',
	KeyCodeS:            'S',
	KeyCodeT:            'T',
	KeyCodeU:            'U',
	KeyCodeV:            'V',
	KeyCodeW:            'W',
	KeyCodeX:            'X',
	KeyCodeY:            'Y',
	KeyCodeZ:            'Z',
	KeyCodeBackquote:    '~',
	KeyCodeBackslash:    '|',
	KeyCodeBackspace:    '\b',
	KeyCodeBracketLeft:  '{',
	KeyCodeBracketRight: '}',
	KeyCodeComma:        '<',
	KeyCodeDelete:       '\x7f',
	KeyCodeDigit0:       ')',
	KeyCodeDigit1:       '!',
	KeyCodeDigit2:       '@',
	KeyCodeDigit3:       '#',
	KeyCodeDigit4:       '$',
	KeyCodeDigit5:       '%',
	KeyCodeDigit6:       '^',
	KeyCodeDigit7:       '&',
	KeyCodeDigit8:       '*',
	KeyCodeDigit9:       '(',
	KeyCodeEnter:        '\n',
	KeyCodeEqual:        '+',
	KeyCodeEscape:       '\x1b',
	KeyCodeMinus:        '_',
	KeyCodePeriod:       '>',
	KeyCodeQuote:        '"',
	KeyCodeSemicolon:    ':',
	KeyCodeSlash:        '?',
	KeyCodeSpace:        ' ',
	KeyCodeTab:          '\t',
}

var UnshiftedAsciiToKeyCodeMap = map[rune]KeyCode{
	'a':    KeyCodeA,
	'b':    KeyCodeB,
	'c':    KeyCodeC,
	'd':    KeyCodeD,
	'e':    KeyCodeE,
	'f':    KeyCodeF,
	'g':    KeyCodeG,
	'h':    KeyCodeH,
	'i':    KeyCodeI,
	'j':    KeyCodeJ,
	'k':    KeyCodeK,
	'l':    KeyCodeL,
	'm':    KeyCodeM,
	'n':    KeyCodeN,
	'o':    KeyCodeO,
	'p':    KeyCodeP,
	'q':    KeyCodeQ,
	'r':    KeyCodeR,
	's':    KeyCodeS,
	't':    KeyCodeT,
	'u':    KeyCodeU,
	'v':    KeyCodeV,
	'w':    KeyCodeW,
	'x':    KeyCodeX,
	'y':    KeyCodeY,
	'z':    KeyCodeZ,
	'`':    KeyCodeBackquote,
	'\\':   KeyCodeBackslash,
	'\b':   KeyCodeBackspace,
	'[':    KeyCodeBracketLeft,
	']':    KeyCodeBracketRight,
	',':    KeyCodeComma,
	'\x7f': KeyCodeDelete,
	'0':    KeyCodeDigit0,
	'1':    KeyCodeDigit1,
	'2':    KeyCodeDigit2,
	'3':    KeyCodeDigit3,
	'4':    KeyCodeDigit4,
	'5':    KeyCodeDigit5,
	'6':    KeyCodeDigit6,
	'7':    KeyCodeDigit7,
	'8':    KeyCodeDigit8,
	'9':    KeyCodeDigit9,
	'\n':   KeyCodeEnter,
	'=':    KeyCodeEqual,
	'\x1b': KeyCodeEscape,
	'-':    KeyCodeMinus,
	'.':    KeyCodePeriod,
	'\'':   KeyCodeQuote,
	';':    KeyCodeSemicolon,
	'/':    KeyCodeSlash,
	' ':    KeyCodeSpace,
	'\t':   KeyCodeTab,
}

var ShiftedAsciiToKeyCodeMap = map[rune]KeyCode{
	'A':    KeyCodeA,
	'B':    KeyCodeB,
	'C':    KeyCodeC,
	'D':    KeyCodeD,
	'E':    KeyCodeE,
	'F':    KeyCodeF,
	'G':    KeyCodeG,
	'H':    KeyCodeH,
	'I':    KeyCodeI,
	'J':    KeyCodeJ,
	'K':    KeyCodeK,
	'L':    KeyCodeL,
	'M':    KeyCodeM,
	'N':    KeyCodeN,
	'O':    KeyCodeO,
	'P':    KeyCodeP,
	'Q':    KeyCodeQ,
	'R':    KeyCodeR,
	'S':    KeyCodeS,
	'T':    KeyCodeT,
	'U':    KeyCodeU,
	'V':    KeyCodeV,
	'W':    KeyCodeW,
	'X':    KeyCodeX,
	'Y':    KeyCodeY,
	'Z':    KeyCodeZ,
	'~':    KeyCodeBackquote,
	'|':    KeyCodeBackslash,
	'\b':   KeyCodeBackspace,
	'{':    KeyCodeBracketLeft,
	'}':    KeyCodeBracketRight,
	'<':    KeyCodeComma,
	'\x7f': KeyCodeDelete,
	')':    KeyCodeDigit0,
	'!':    KeyCodeDigit1,
	'@':    KeyCodeDigit2,
	'#':    KeyCodeDigit3,
	'$':    KeyCodeDigit4,
	'%':    KeyCodeDigit5,
	'^':    KeyCodeDigit6,
	'&':    KeyCodeDigit7,
	'*':    KeyCodeDigit8,
	'(':    KeyCodeDigit9,
	'\n':   KeyCodeEnter,
	'+':    KeyCodeEqual,
	'\x1b': KeyCodeEscape,
	'_':    KeyCodeMinus,
	'>':    KeyCodePeriod,
	'"':    KeyCodeQuote,
	':':    KeyCodeSemicolon,
	'?':    KeyCodeSlash,
	' ':    KeyCodeSpace,
	'\t':   KeyCodeTab,
}
