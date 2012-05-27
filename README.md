djw_learning_golang
===================

Code used to evaluate and learn Go and to solve ProjectEuler.net problems.

```
/**
 * @author David Wortham <djwortham@gmail.com>
 */
```

About
=====

[ProjectEuler.net](http://projecteuler.net/) is a good way to test your basic programming knowledge against a new language.  The problems get progressively harder (as a general trend) and some reuse knowledge/code from previous problems so you might put on your refactoring hat to save yourself some work. DRY!

[Go](http://golang.org/) is a relatively new language from Google that is fast, compiles down to machine code, and removes much of the cruft found in other higher-level languages.  I decided to give it a go and learn it for myself.

What I've Learned so far
========================

I found myself pleasantly surprised that unit testing is built into the language tools.  Benchmarkeing is also available from the same command.

Go 1.0
------

Go is now in 1.0 trim as of March 2012. Sadly there were lots of API changes in Go since the beta days so a lot of the code on the web is outdated.  Also, "go" is a common English word so the SNR on Google and StackOverflow searches is low (although "Golang" is a good substitute).


Code examples lacking
---------------------

I'm a little disheartened that Go has so few code examples in the official documentation.  Hopefully the Go+GitHub community will be able to rectify that.

TextMate Bundle
---------------

I learned that Alan Quatermain built a "bundle" for my favorite IDE (TextMate 1.x for Mac).  It was, sadly, built for a beta version of Go.  GitHubbers have, I'm happy to discover, forked and fixed many of these issues for the Go 1.0 public release.  The most authoritative fork I've found yet is here:
[Go TextMate Bundle](https://github.com/sdefresne/Go.tmbundle).
It supposedly has fixes to work with the upcoming TextMate 2.0, but I wouldn't know.


Best uses
---------

Go makes is very easy to spool up a server (HTTP, SMTP, DNS, etc.) and is designed for high concurrency, but I find it lacking in my forté: Web Applications development.  I'll be evaluating its value as a Web Application framework flatform in the near future.



License
=====

This entire git repository is licensed under the MIT License

```
Copyright (C) 2012 David Wortham

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

Please drop me a line if you are using any code licensed from this repository for any commercial and/or non-profit activitied; I'm just interested in how it's used.
