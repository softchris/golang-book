# Use regular expressions to parse text

In this chapter, you'll learn how to use Regular Expressions to search and replace text.

## What is RegEx and what to use it for

Regular Expressions or RegEx for short, is used for searching and replacing text. Technically a RegEx is a sequence of characters that specifies a search pattern.

## Why use RegEx

RegEx primary usage area is for searching text, replacing it and also extracting text. There do exist string libraries that can do some of the functionality RegEx is capable of. Sometimes using those string libraries might even be the best thing to do. However, sometimes a RegEx pattern is better.

> Fair word of warning though, RegEx is hard to get right. You are encouraged to learn more of how they work cause they are quite powerful.

My hope is by you reading this chapter, that you will find RegEx less intimidating and see it as a valuable tool in your toolbox.

## Where is it used ?

RegEx shows up in many different contexts:

- **Text editors, any programs with a search**. In most text editors, for example Visual Studio Code, you can search for files and inside of files with a RegEx search pattern.
- **Code**, many programming languages and runtimes have libraries that helps you use RegEx.  

## Your first RegEx

Let's construct a simple RegEx to get a feel for it. Here it is:

```text
an
```

if you apply this search pattern `an` to the following text:

```text
highlands is a part of Scotland
```

It will match like so:

> highl**an**ds is a part of Scotl**an**d

For simpler cases, where you are looking to see if a specific word matches, in one or more places in a sentence, a pattern like the above is enough.

### RegEx in Go

To start using RegEx in Go, there's the regexp library. There are two approaches:

- `regexp` directly, here's an example:

   ```go
   matched, err := regexp.FindString("an", "highlands is a part of Scotland")
   ```

   Here, you get a boolean back that returns true if there's a match.

- **compiled**, in this way, you compile a regular expression and then calls a method on it, like so:

   ```go
   r, _ := regexp.Compile("an")
   matches := r.FindAllString("highlands is a part of Scotland", -1)
   ```

   The above returns a string array with all the matches, in this case `["an", "an"]`.

   In this version, you have more functions available.  

## Character classes

Character classes are able to distinguish between different types of characters. Different types can be newlines, digits, letters and so on.

Let's have a look at some common types you are likely to encounter:

|Type  |Description  |
|---------|---------|
|.     | This type matches any character except for a carriage return        |
|\     | This type escapes what's coming next        |
| \w | matches any character from the latin alphabet including underscore _ |
| \d | matches any digit |
| \D | this is the inverse of \d and matches any character that's not a digit |
| \s | matches a white space character like space tab, line feed etc. |

Lets show an example:

```go
matched, err := regexp.FindString("\d", "abc123")
```

There would be a match above due to 123. However, there would be no match against "abc" as there's no digits in it.

## Repetition

If you want to express repetition, there's two characters of interest:

- `+`, matches 1 to many characters.

   ```text
   \w+
   ```

   Given the string "aaaa bab" it would match:  

   **aaaaab** **bab** as the above describes matching characters but not the white space.

   ```go
   r, _ := regexp.Compile("\\w+")
   matches := r.FindAllString("aaaa bab", -1)
   ```

   Note the extra `\`, we need that because of the way we construct our Regex.

- `*`, matches 0 to many characters. Lets say you want to match a postal address that starts with "PA" and may contain 0 or many numbers. It should then match strings:

   ```text
   PA
   PA111 
   ```

   We can use a `*` to construct this looking like so:

   ```go
   regexp.MatchString("PA*", "PA")
   regexp.MatchString("PA*", "PA111")
   ```

- `?`, also known as a greedy or optional quantifier. It looks backwards and makes it optional and takes it, if it can. Consider this case:

   ```text
   http
   https
   ```  

   If you want to match them both, you can type:

   ```text
   https?
   ```  

   Another example is:

   ```text
   r, _ := regexp.Compile("an.")
   matches := r.FindAllString("and ant an", -1)
   ```

   The above will only match **and** and **ant** but not *an*. If we modify the regex to `an.?` it will match **and ant an**.

## Anchors and boundaries

There are different anchors you can use like for example:

- `^`, beginning of the string. The following states that the string needs to begin with the following string "INV" to signify the start of an invoice row:

   ```text
   ^INV
   ```

- `$`, end of the string. An example could be matching a string ends with a certain domain ".com":

   ```text
   \.com$
   ```

## Groups

Groups are way to capture part of a string and have that returned. It's very useful for parsing out the info you need. Consider this example parsing out the info from a CSV row:

```text
Name: myarticle, Price: 114, Quantity: 3 
```

To get the data you need, you want everything after the colon, :. You can construct a RegEx like so:

```go
\w+:\s?(\w+)
```

what we are doing is defining we want to capture a group using parenthesis `()` but that group should happen after:

- a number of letters, `\w+`
- followed by a colon, `:`
- followed by 0 or 1 space `\s?`
- then our group `(\w+)`, one ore more letters

All this ends up capturing myarticle, 114 and 3.

## Named groups

A named group is a group you want to capture where the groups have names. Why would you want that? Well, say that you want to break down a URL in pieces and wants to know what's what. Given a URL "http://myapi.com/products?page=1", you have:

- `http`, the protocol.
- `myapi.com`, the domain.
- `/products`, is the route.
- `?page=1`, is the query parameters.

So how can we break it apart and give it a name?

Well, to break it apart, we will use something called named groups, it will allow us to look at our matches and know what's what. So instead of getting:

```text
http
```

We will get a key and value that says:

```text
protocol: http
```

Syntax wise, we need to use `?<name of our group>` within our parenthesis ().

You use the following syntax:

```text
(?<mygroup>\w+)
```

In Go, we need a `P` right after the question mark, so the code for this would be:

```go
r, err := regexp.Compile(`(?P<mygroup>\w+):`)
```

### Extract the data from a URL

Let's approach this problem then given the string "http://myapi.com/products?page=1":

- matching the protocol:

    ```text
    ^(?<protocol>\w+):
    ```

- domain, to match the domain as well, we're looking to capture everything after http:// and until the next /:

   ```text
   ^(?<protocol>\w+):\/\/(?<domain>\w+\.\w+)\/?
   ```

- route, ok so we've matched up "http://mydomain.com" so far, now lets match the route, i.e what happens after the / but before any questions marks, ?
- query params

Here's what our Go code would look like:

```go
r, err := regexp.Compile(`^(?P<protocol>\w+):\/\/(?P<domain>\w+\.\w+)\/(?P<route>\w+)\/?`)
```

Ok, so we have the pattern, what about printing the parsed parts?

To pair the named groups with their values, we need to combine values from both the Regex and the response. First, we call `FindStringSubmatch()`, that will give us the values.

```go
m := r.FindStringSubmatch("http://myapi.com/products")
```

Then, we need to match the names with these values. We will need to call `r.SubexpNames()` and iterate over the response.

```go
result := make(map[string]string)
 for i, name := range r.SubexpNames() {
  if i != 0 && name != "" {
   result[name] = m[i]
  }
 }
```

Note this line where each name is assigned a value:

```go
result[name] = m[i]
```

Finally, to get the values, we can print them out as they are now in a map structure:

```go
fmt.Println(result["protocol"]) // http
fmt.Println(result["domain"]) // myapi.com
fmt.Println(result["route"]) // products
```

## Assignment - create a Go program that parses a URL

From the above use case on named groups, write a Go program that takes a URL and analyzes it. It should work like so:

```output
Type URL: http://myapi.com/products
The URL consist of:
protocol: http
domain: myapi.com
route: products
```

### Solution

```go
package main

import (
 "fmt"
 "log"
 "regexp"
)

func main() {
 var url string
 fmt.Println("Type URL: ")
 fmt.Scan(&url)

 r, err := regexp.Compile(`^(?P<protocol>\w+):\/\/(?P<domain>\w+\.\w+)\/(?P<route>\w+)\/?`)
 if err != nil {
  log.Fatal("Error compiling: ", err)
 }
 m := r.FindStringSubmatch(url)
 if m == nil {
  panic("mo match")
 }
 result := make(map[string]string)
 for i, name := range r.SubexpNames() {
  if i != 0 && name != "" {
   result[name] = m[i]
  }
 }
 fmt.Println("The URL consist of:")
 fmt.Println(result["protocol"])
 fmt.Println(result["domain"])
 fmt.Println(result["route"])
}

```

## Replacing

A common use case for Regex is when it's used to replace something with something else.

There's more than one method in Go you could be using but one you could use is `ReplaceAllString()` that sits on the compiled RegEx object:

```go
r := regexp.MustCompile(`aa`)
 s := r.ReplaceAllString("aabbcc", "cc") // s = ccbbcc
```

The above replaces all occurrences of `aa` with `cc` on the string `aabbcc`.

You can also use capture groups and replace a captured group with a string. Here's an example:

```go
r := regexp.MustCompile(`(\d)`)
 s := r.ReplaceAllString("productid:114", "0${1}") // s = productid:0114
```

in the above case, we replace 114 with itself but we also prepend it with a 0.

### Use case, replace XML Nodes

Imagine you are working with XML for example and want to rename all nodes with a certain name.

Here's your XML

```xml
<books>
    <book>
      <author>Shakespeare</author>
      <title>Romeo and Juliet</title>
      <pages>400</pages>
      <type>paperback</type>
      <cost>17</cost>
    </book>
    <book>
      <author>Shakespeare</author>
      <title>Hamlet</title>
      <pages>270</pages>
      <type>paperback</type>
      <cost>15</cost>
    </book>
</books>
```

Imagine `title` should be replaced by `name`, how do we do that?

Well, it would be straight forward to replace title by name. Let's say we have this file content though:

```xml
<books>
    <book>
      <author>Shakespeare</author>
      <title>The title is Romeo and Juliet</title>
      <pages>400</pages>
      <type>paperback</type>
      <cost>17</cost>
    </book>

</books>
```

Then we would not only rename the element `title` to `name` but also the content would be replaced o "The title is Romeo and Juliet", that's NOT what we want.

We need to restrict the replace operation to only target element, like so:

```text
\<\/?(title)\>
```

The above would match for example `<title>` and `</title>`. If we try this however on this XML, we almost get what we want:

```xml
<author>Shakespeare</author>
```

becomes

```xml
nameShakespearename
```

What happened, why did we loose `<>` ? We need a way to express keeping what was there before AND replace the name. A way to do that is to express capture groups on `<>` and the element name, like so:

```text
(\<\/?)(title)(\>)
```

Now we have three groups, we need to fit the result together, and this is something we can express like so:

```text
${1}name${3}
```

- `${1}` corresponds to capture group matching `<` or `</`
- `name` is the string we replace `title` with.
- `{3}` corresponds to capture group matching `>`.

## Assignment - replace content

Take the file *books.xml* containing:

```xml
<books>
    <book>
      <author>Shakespeare</author>
      <title>Romeo and Juliet</title>
      <pages>400</pages>
      <type>paperback</type>
      <cost>17</cost>
    </book>
    <book>
      <author>Shakespeare</author>
      <title>Hamlet</title>
      <pages>270</pages>
      <type>paperback</type>
      <cost>15</cost>
    </book>
</books>
```

and replace:

- author with name
- cost with price

TIP: you might need to apply the replace twice.

## Solution II

```go
package main

import (
 "fmt"
 "regexp"
)

func main() {
 file := `<books>
    <book>
      <author>Shakespeare</author>
      <title>Romeo and Juliet</title>
      <pages>400</pages>
      <type>paperback</type>
      <cost>17</cost>
    </book>
    <book>
      <author>Shakespeare</author>
      <title>Hamlet</title>
      <pages>270</pages>
      <type>paperback</type>
      <cost>15</cost>
    </book>
</books>`

 r := regexp.MustCompile(`(\<\/?)(title)(\>)`)
 s := r.ReplaceAllString(file, "${1}name${3}")
 fmt.Println(s)

 r = regexp.MustCompile(`(\<\/?)(cost)(\>)`)
 s = r.ReplaceAllString(s, "${1}price${3}")
 fmt.Println(s)
}
```
