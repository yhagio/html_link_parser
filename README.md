# HTML link parser

Get all links (`<a>` tags) and print out the href link and its text
from a provided url.

### How to run

```bash
go run main.go
```

Then it asks you to type URL to parse.
(in this example, I use `https://golang.org/`)

Output

```
$ go run main.go

=====👽HTML Link Parser👽=====

Type URL (i.e. https://golang.org):
http://golang.org

Reading http://golang.org ...
Parsing ...

<<< RESULT (Total: 18) >>>

(1)
 Link:/
 Text:The Go Programming Language

(2)
 Link:/
 Text:Go

(3)
 Link:#
 Text:▽

(4)
 Link:/doc/
 Text:Documents

(5)
 Link:/pkg/
 Text:Packages

(6)
 Link:/project/
 Text:The Project

(7)
 Link:/help/
 Text:Help

(8)
 Link:/blog/
 Text:Blog

(9)
 Link:http://play.golang.org/
 Text:Play

(10)
 Link:#
 Text:Run

(11)
 Link:#
 Text:Share

(12)
 Link://tour.golang.org/
 Text:Tour

(13)
 Link:/dl/
 Text:Download Go Binary distributions available for Linux, Mac OS X, Windows, and more.

(14)
 Link://blog.golang.org/
 Text:Read more

(15)
 Link:https://developers.google.com/site-policies#restrictions
 Text:noted

(16)
 Link:/LICENSE
 Text:BSD license

(17)
 Link:/doc/tos.html
 Text:Terms of Service

(18)
 Link:http://www.google.com/intl/en/policies/privacy/
 Text:Privacy Policy
```
