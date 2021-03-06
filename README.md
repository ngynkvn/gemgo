# GemGo
a Gemini protocol compatible client (WIP)

Gemgo is a primitive client for the [Gemini protocol](https://gemini.circumlunar.space/), 
a simple request-response transaction protocol somewhat similar to HTTP or Gopher

Below is a screencast of the website linked above as viewed from Gemini.

[![asciicast](https://asciinema.org/a/334265.svg)](https://asciinema.org/a/334265)
## Example
#### To run
```
go run ./cmd/client/app.go/ [url]
```
Try with `go run ./cmd/client/app.go/ gemini://gemini.circumlunar.space`
```
PS C:\Users\Kev\Desktop\geminimi> go run .\cmd\client\app.go gemini.circumlunar.space
2020/05/26 20:30:06 Scheme was not given, assuming scheme gemini
2020/05/26 20:30:06 Port was not given, assuming port 1965
2020/05/26 20:30:06 Visiting gemini://gemini.circumlunar.space:1965/
header [20] text/gemini
0       | # Project Gemini
1       |
2       | ## Overview
3       |
4       | Gemini is a new internet protocol which:
5       |
6       | * Is heavier than gopher
7       | * Is lighter than the web
8       | * Will not replace either
9       | * Strives for maximum power to weight ratio
10      | * Takes user privacy very seriously
11      |
12      | ## Resources
13      |
14      | => docs/      Gemini documentation
15      | => software/  Gemini software
16      | => servers/   Known Gemini servers
17      | => https://lists.orbitalfox.eu/listinfo/gemini        Gemini mailing list
18      | => gemini://gemini.conman.org/test/torture/   Gemini client torture test
19      |
20      | ## Web proxies
21      |
22      | => https://portal.mozz.us/?url=gemini%3A%2F%2Fgemini.circumlunar.space%2F&fmt=fixed   Gemini-to-web proxy service
23      | => https://proxy.vulpes.one/gemini/gemini.circumlunar.space   Another Gemini-to-web proxy service
24      |
25      | ## Search engines
26      |
27      | => gemini://gus.guru/ Gemini Universal Search engine
28      | => gemini://houston.coder.town        Houston search engine
29      |
30      | ## Geminispace aggregators (experimental!)
31      |
32      | => capcom/    CAPCOM
33      | => gemini://rawtext.club:1965/~sloum/spacewalk.gmi    Spacewalk
34      |
35      | ## Free Gemini hosting
36      |
37      | => users/     Users with Gemini content on this server
```

## Spec

https://gemini.circumlunar.space/docs/spec-spec.txt

## Why?

I wanted to see the "hidden" layer of the Project Gemini website, and the spec was easy to understand. 

(Downloading a binary probably would have been easier)
