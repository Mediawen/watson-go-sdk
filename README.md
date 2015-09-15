# watson-go-sdk

Simple Watson SDK for the [Go programming language](http://golang.org/).

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  * [Installation](#installation)
  * [Usage](#usage)
  * [Demos](#demos)
  * [Contributors](#contributors)
  * [License](#license)

---------------------------------------
## Features

  * Lightweight
  * Native Go implementation. No C-bindings, just pure Go

So far only [Speech To Text](http://www.ibm.com/smarterplanet/us/en/ibmwatson/developercloud/speech-to-text.html) and [Text To Speech](http://www.ibm.com/smarterplanet/us/en/ibmwatson/developercloud/text-to-speech.html) functionalities have been implemented. 

---------------------------------------
## Requirements

  * Go 1.2 or higher

---------------------------------------
## Installation

- `watson-go-sdk` is available as a normal Go package with this Github branch. Just include it in your dependencies (imports) on your code. Then type:

go get "github.com/mediawen/watson-go-sdk"

---------------------------------------
## Usage

- Coming Soon

---------------------------------------
## Roadmap

- Fix issues reported on this repository
- Add other Watson functionalities:

---------------------------------------
## Demos

- Text To Speech API [demo](https://speech-to-text-demo.mybluemix.net/).
- Speech To Text API [demo](https://text-to-speech-demo.mybluemix.net/).
- At [MediaWen International](http://mediawen.com), we use these technologies to enhance our platform of closed captioning, subtitling, and automatic dubbing: [STVHub](http://stvhub.com). By exemple, we generated the voice over (or Automatic Dubbing) on following video of the French minister of foreing affairs. Just watch it and listen him in [Spanish](https://www.youtube.com/watch?v=tF852LsSwoo) or in [English](https://www.youtube.com/watch?v=8sWZMea-q2I).

![Automatic Dubbing image](doc/img/fabius.jpg)

---------------------------------------
## Contributors

- [Philippe Anel](https://github.com/xigh)

---------------------------------------
## License

The file cookie.go has been picked from go net/http package. It's under the following BSD License:

    Copyright (c) 2012 The Go Authors. All rights reserved.
    
    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are
    met:
    
       * Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
       * Redistributions in binary form must reproduce the above
    copyright notice, this list of conditions and the following disclaimer
    in the documentation and/or other materials provided with the
    distribution.
       * Neither the name of Google Inc. nor the names of its
    contributors may be used to endorse or promote products derived from
    this software without specific prior written permission.
    
    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

All the rest of the code is under the the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
