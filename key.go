package glimmer

import "golang.org/x/mobile/event/key"

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found below:

// Copyright (c) 2009 The Go Authors. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
// 
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// NOTE: this makes it so no x/mobile deps are needed downstream

type KeyCode uint32

const (
	CodeUnknown KeyCode = KeyCode(key.CodeUnknown)

	CodeA KeyCode = KeyCode(key.CodeA)
	CodeB KeyCode = KeyCode(key.CodeB)
	CodeC KeyCode = KeyCode(key.CodeC)
	CodeD KeyCode = KeyCode(key.CodeD)
	CodeE KeyCode = KeyCode(key.CodeE)
	CodeF KeyCode = KeyCode(key.CodeF)
	CodeG KeyCode = KeyCode(key.CodeG)
	CodeH KeyCode = KeyCode(key.CodeH)
	CodeI KeyCode = KeyCode(key.CodeI)
	CodeJ KeyCode = KeyCode(key.CodeJ)
	CodeK KeyCode = KeyCode(key.CodeK)
	CodeL KeyCode = KeyCode(key.CodeL)
	CodeM KeyCode = KeyCode(key.CodeM)
	CodeN KeyCode = KeyCode(key.CodeN)
	CodeO KeyCode = KeyCode(key.CodeO)
	CodeP KeyCode = KeyCode(key.CodeP)
	CodeQ KeyCode = KeyCode(key.CodeQ)
	CodeR KeyCode = KeyCode(key.CodeR)
	CodeS KeyCode = KeyCode(key.CodeS)
	CodeT KeyCode = KeyCode(key.CodeT)
	CodeU KeyCode = KeyCode(key.CodeU)
	CodeV KeyCode = KeyCode(key.CodeV)
	CodeW KeyCode = KeyCode(key.CodeW)
	CodeX KeyCode = KeyCode(key.CodeX)
	CodeY KeyCode = KeyCode(key.CodeY)
	CodeZ KeyCode = KeyCode(key.CodeZ)
   
	Code1 KeyCode = KeyCode(key.Code1)
	Code2 KeyCode = KeyCode(key.Code2)
	Code3 KeyCode = KeyCode(key.Code3)
	Code4 KeyCode = KeyCode(key.Code4)
	Code5 KeyCode = KeyCode(key.Code5)
	Code6 KeyCode = KeyCode(key.Code6)
	Code7 KeyCode = KeyCode(key.Code7)
	Code8 KeyCode = KeyCode(key.Code8)
	Code9 KeyCode = KeyCode(key.Code9)
	Code0 KeyCode = KeyCode(key.Code0)

	CodeReturnEnter        KeyCode = KeyCode(key.CodeReturnEnter)
	CodeEscape             KeyCode = KeyCode(key.CodeEscape)
	CodeDeleteBackspace    KeyCode = KeyCode(key.CodeDeleteBackspace)
	CodeTab                KeyCode = KeyCode(key.CodeTab)
	CodeSpacebar           KeyCode = KeyCode(key.CodeSpacebar)
	CodeHyphenMinus        KeyCode = KeyCode(key.CodeHyphenMinus)
	CodeEqualSign          KeyCode = KeyCode(key.CodeEqualSign)
	CodeLeftSquareBracket  KeyCode = KeyCode(key.CodeLeftSquareBracket)
	CodeRightSquareBracket KeyCode = KeyCode(key.CodeRightSquareBracket)
	CodeBackslash          KeyCode = KeyCode(key.CodeBackslash)
	CodeSemicolon          KeyCode = KeyCode(key.CodeSemicolon)
	CodeApostrophe         KeyCode = KeyCode(key.CodeApostrophe)
	CodeGraveAccent        KeyCode = KeyCode(key.CodeGraveAccent)
	CodeComma              KeyCode = KeyCode(key.CodeComma)
	CodeFullStop           KeyCode = KeyCode(key.CodeFullStop)
	CodeSlash              KeyCode = KeyCode(key.CodeSlash)
	CodeCapsLock           KeyCode = KeyCode(key.CodeCapsLock)

	CodeF1  KeyCode = KeyCode(key.CodeF1)
	CodeF2  KeyCode = KeyCode(key.CodeF2)
	CodeF3  KeyCode = KeyCode(key.CodeF3)
	CodeF4  KeyCode = KeyCode(key.CodeF4)
	CodeF5  KeyCode = KeyCode(key.CodeF5)
	CodeF6  KeyCode = KeyCode(key.CodeF6)
	CodeF7  KeyCode = KeyCode(key.CodeF7)
	CodeF8  KeyCode = KeyCode(key.CodeF8)
	CodeF9  KeyCode = KeyCode(key.CodeF9)
	CodeF10 KeyCode = KeyCode(key.CodeF10)
	CodeF11 KeyCode = KeyCode(key.CodeF11)
	CodeF12 KeyCode = KeyCode(key.CodeF12)

	CodePause         KeyCode = KeyCode(key.CodePause)
	CodeInsert        KeyCode = KeyCode(key.CodeInsert)
	CodeHome          KeyCode = KeyCode(key.CodeHome)
	CodePageUp        KeyCode = KeyCode(key.CodePageUp)
	CodeDeleteForward KeyCode = KeyCode(key.CodeDeleteForward)
	CodeEnd           KeyCode = KeyCode(key.CodeEnd)
	CodePageDown      KeyCode = KeyCode(key.CodePageDown)

	CodeRightArrow KeyCode = KeyCode(key.CodeRightArrow)
	CodeLeftArrow  KeyCode = KeyCode(key.CodeLeftArrow)
	CodeDownArrow  KeyCode = KeyCode(key.CodeDownArrow)
	CodeUpArrow    KeyCode = KeyCode(key.CodeUpArrow)

	CodeKeypadNumLock     KeyCode = KeyCode(key.CodeKeypadNumLock)
	CodeKeypadSlash       KeyCode = KeyCode(key.CodeKeypadSlash)
	CodeKeypadAsterisk    KeyCode = KeyCode(key.CodeKeypadAsterisk)
	CodeKeypadHyphenMinus KeyCode = KeyCode(key.CodeKeypadHyphenMinus)
	CodeKeypadPlusSign    KeyCode = KeyCode(key.CodeKeypadPlusSign)
	CodeKeypadEnter       KeyCode = KeyCode(key.CodeKeypadEnter)
	CodeKeypad1           KeyCode = KeyCode(key.CodeKeypad1)
	CodeKeypad2           KeyCode = KeyCode(key.CodeKeypad2)
	CodeKeypad3           KeyCode = KeyCode(key.CodeKeypad3)
	CodeKeypad4           KeyCode = KeyCode(key.CodeKeypad4)
	CodeKeypad5           KeyCode = KeyCode(key.CodeKeypad5)
	CodeKeypad6           KeyCode = KeyCode(key.CodeKeypad6)
	CodeKeypad7           KeyCode = KeyCode(key.CodeKeypad7)
	CodeKeypad8           KeyCode = KeyCode(key.CodeKeypad8)
	CodeKeypad9           KeyCode = KeyCode(key.CodeKeypad9)
	CodeKeypad0           KeyCode = KeyCode(key.CodeKeypad0)
	CodeKeypadFullStop    KeyCode = KeyCode(key.CodeKeypadFullStop)
	CodeKeypadEqualSign   KeyCode = KeyCode(key.CodeKeypadEqualSign)

	CodeF13 KeyCode = KeyCode(key.CodeF13)
	CodeF14 KeyCode = KeyCode(key.CodeF14)
	CodeF15 KeyCode = KeyCode(key.CodeF15)
	CodeF16 KeyCode = KeyCode(key.CodeF16)
	CodeF17 KeyCode = KeyCode(key.CodeF17)
	CodeF18 KeyCode = KeyCode(key.CodeF18)
	CodeF19 KeyCode = KeyCode(key.CodeF19)
	CodeF20 KeyCode = KeyCode(key.CodeF20)
	CodeF21 KeyCode = KeyCode(key.CodeF21)
	CodeF22 KeyCode = KeyCode(key.CodeF22)
	CodeF23 KeyCode = KeyCode(key.CodeF23)
	CodeF24 KeyCode = KeyCode(key.CodeF24)

	CodeHelp KeyCode = KeyCode(key.CodeHelp)

	CodeMute       KeyCode = KeyCode(key.CodeMute)
	CodeVolumeUp   KeyCode = KeyCode(key.CodeVolumeUp)
	CodeVolumeDown KeyCode = KeyCode(key.CodeVolumeDown)

	CodeLeftControl  KeyCode = KeyCode(key.CodeLeftControl)
	CodeLeftShift    KeyCode = KeyCode(key.CodeLeftShift)
	CodeLeftAlt      KeyCode = KeyCode(key.CodeLeftAlt)
	CodeLeftGUI      KeyCode = KeyCode(key.CodeLeftGUI)
	CodeRightControl KeyCode = KeyCode(key.CodeRightControl)
	CodeRightShift   KeyCode = KeyCode(key.CodeRightShift)
	CodeRightAlt     KeyCode = KeyCode(key.CodeRightAlt)
	CodeRightGUI     KeyCode = KeyCode(key.CodeRightGUI)

	CodeCompose KeyCode = KeyCode(key.CodeCompose)
)
