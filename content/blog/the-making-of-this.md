+++
date = "2017-01-28T15:15:02-08:00"
description = "Behind the Scenes"
title = "Who, What, Why, and How"

+++

## A New Hope

I'm terrible at blogging, there is no doubt about it. I have struggled for many years with it, but now I have hit a point in my career where having a personal blog is recommended, and it's best for me to suck it up and maintain one.

With that in mind, I figured I would go into some details on how this blog is built. Who, I guess is irrelevant (just read my about page!), so I'll talk about What I used, Why I used it, and How I use it. When I titled this, I was going to introduce myself, but I ended up moving that to the About page, as well my industry's standard disclaimer of `this is mine, not representative of my employer, blah blah blah`.

## The Empire Strikes Back

In regards to what I used, I actually used a static website generator. I have tried using a fully-featured CMS engine, like [WordPress](https://wordpress.org/), [Medium](https://medium.com/), or [Ghost](https://ghost.org/). While I thought all of them were pretty spiffy, I felt it was entirely too much for me to cope with. I wanted control, I wanted to host it myself, and I wanted to be able to treat my blog like code. Blogging as Code(<sup>TM</sup>). I have tried [Jekyll](https://jekyllrb.com/) in the past, and it was alright; it had one major flaw, though: *IT IS WRITTEN IN RUBY*.

For those of you that don't know me, I have a deep set disgust for most things Ruby. While Ruby is fine, and it is interpreted similar to Python, I can't get over the syntax. I feel like I'm programming like a child with it's simplistic, expressive language. It's slow, and any form of parallelism or concurrency is just...challenging.

So Jekyll was out the door, time to search for something I could contribute to, if I ever had the desire. I'm a huge fan of Google's [Go](https://golang.org), so I started looking for something like Jekyll, but in Go. I didn't have to look far, I stumbled across [Hugo](https://gohugo.io/). Their tag line: `Make the Web Fun Again`. Perfect, I like fun. Especially as I'm not a UI wizard...those mystical powers reside with others. Looking into Hugo, I found it was produced by [spf13](http://spf13.com/), who writes one of my favourite command line client interfaces, [cobra](https://github.com/spf13/cobra). I was exceedingly pleased, and I knew he experience would be a fun one, both to learn and to maintain.

> Yes We Can.
> 
> --<cite>Barack Obama</cite>

## Return of the Jedi

I searched long and hard for the right [theme](http://themes.gohugo.io/). I'm picky about presentation, and I feel first impressions are the most important. I tried so many different themes, for awhile. After lots of *close, but no cigar* tries, I finally settled on this one, [Cactus Plus](http://themes.gohugo.io/hugo-theme-cactus-plus/). It's simplistic, fast (all of Hugo is fast, lol), the text is okay, and it's overall, a decent theme. I like the [Highlight.js](https://highlightjs.org/) markdown renderer, it produces nice contrast to the white style. This is the Why.

Python:
```python
from sys import stdout; stdout.write("See? Looks good.")
```
Go:
```go
package main

import "fmt"

func main() {
    fmt.Println("See? Still looks good.")
}
```
C:
```c
#include <stdio.h>
int main() {
    printf("Still going strong.");
    return 0;
}
```
C#:
```csharp
using System;

class Program
{
    static void Main()
    {
        string works = "Man, this highlighter is super awesome.";
        Console.WriteLine(works);
    }
}
```

I really like it, it's clean. Now, I have highlight.js configured to use the [Darcula](https://github.com/bulenkov/Darcula) theme from Konstantin Bulenkov. I use various [Jetbrains](https://www.jetbrains.com/) editors, and they all use the Darcula theme. I feel it's easier on the eyes. I'll go over how I integrated Hugo with WebStorm in a later post.


## The Force Awakens

How I made this is pretty simple, I mostly just followed the directions and Read The F***ing Manual. One of my favourite sayings:

> Real men don't need instructions.
>
> --<cite>Tim "The Toolman" Taylor</cite>

Is so unbelievably inaccurate it's laughable, lol. I read manuals all the time. API, vehicle, toaster (oh yes),  and many other manuals. Without manuals, I would be helpless. I would highly recommend reading the [Hugo manual](https://gohugo.io/overview/introduction/) if you want to create your own.

That's about it, for now. Thanks!