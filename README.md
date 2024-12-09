# UY to USD converter

## Why

I made this project to learn go primarily, but I also wanted a converter that
I could have offline and easily accessible.

## How

The idea is that it fetches the conversion rate at startup, and displays a
window

### Fetcher

The main problem of this application is actually fetching the conversion rate.

While testing ways to fetch it without paying an API key, I was able to get it
programatically from doing an http request to a site containing this number,
searching in the body to find it, and then parse it from text to float.

There might be improvements in the fetch, but works well enough to stay.
Currently it works by searching for the parent of the number which has a unique
class, then searching when the parent closes, and getting whats between.

### GUI

#### Tech

For the GUI I used the http://www.github.com/lxn/win framework, specialized for
making GUI with winAPI for Windows.

Migration to other frameworks should be easy, but I figured this was a great
start.

#### Layout

The layout consist of a label on top displaying the conversion rate, a row of
two text boxes and a convert button at the bottom.

When the users inputs a number in the left text box and presses the convert
button, the input is taken and divided by the conversion rate, and displayed in
the right text box.

If the conversion fails "error" is displayed in the right text box.

## Limits

The fetcher is a bit sketchy and manual. Any update that changes the identifier
of the conversion rate parent will break the fetch. If this happens a fallback
convertion rate is used, which is 42.0.

Also my awesome frontend abilities don't help me a lot, and I couldn't figure
out an alternative of using text boxes from the framework that I use.
