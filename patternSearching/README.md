# patternSearching

Pattern searching program written in Golang

## Usage

```bash
$ go run main.go abcdefghijklmn def
Text:     abcdefghijklmn
Pattern:  def
Match at: 3

$ go run main.go アメンボ赤いなあいうえお 赤い
Text:     アメンボ赤いなあいうえお
Pattern:  赤い
Match at: 4
```

```bash
$ ./speedtest.sh
naive

real    0m0.326s
user    0m0.259s
sys     0m0.173s

improved

real    0m0.288s
user    0m0.236s
sys     0m0.149s
```


## License

### MIT License

Copyright (c) 2019 Masaharu TASHIRO

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
