# shakerSortByDoublyLinkedList

Shaker sort using doubly linked list written in Golang

## Usage

```bash
$ go run main.go 20
N: 20
Unsorted list
19,13,11,15,15,20,31,4,20,31,30,11,1,20,29,30,22,15,11,25,
----------------------------------------
13,11,15,15,19,20,4,20,31,30,11,1,20,29,30,22,15,11,25,31,
1,13,11,15,15,19,20,4,20,31,30,11,11,20,29,30,22,15,25,31,
1,11,13,15,15,19,4,20,20,30,11,11,20,29,30,22,15,25,31,31,
1,4,11,13,15,15,19,11,20,20,30,11,15,20,29,30,22,25,31,31,
1,4,11,13,15,15,11,19,20,20,11,15,20,29,30,22,25,30,31,31,
1,4,11,11,13,15,15,11,19,20,20,15,20,22,29,30,25,30,31,31,
1,4,11,11,13,15,11,15,19,20,15,20,20,22,29,25,30,30,31,31,
1,4,11,11,11,13,15,15,15,19,20,20,20,22,25,29,30,30,31,31,
1,4,11,11,11,13,15,15,15,19,20,20,20,22,25,29,30,30,31,31,
----------------------------------------
Sorted list
1,4,11,11,11,13,15,15,15,19,20,20,20,22,25,29,30,30,31,31,
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
