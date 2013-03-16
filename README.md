This repository was inspired by mattyw's blog post [Using go to unmarshal json lists with multiple types](http://mattyjwilliams.blogspot.co.uk/2013/01/using-go-to-unmarshal-json-lists-with.html)

To summarize the article, we are given a JSON string of the form:

```json
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        }
    ]
}
```

And our goal is to unmarshal it into a Go data structure. The article goes into more detail, and two solutions were proposed. A commenter came up with a third solution, and another commenter [dustin](https://github.com/dustin/) proposed using his library called [jsonpointer](https://github.com/dustin/go-jsonpointer), which operates on the raw byte array of the json string, instead of unmarshalling first and then traversing the data structure.

I used Dustin's library, and to great avail, the only gotcha was that json strings were returned with the double quotes in them and some trailing spaces, but I made a little function that returned a slice of the original bytes: 

```go
func trimJsonBytes(toTrim []byte) []byte {
    // implementation found in solution.go
}
```

Here is the algorithm:

```
loop i:=0,1,2,...∞
  if  /things/i/name  is empty
    make a person
    append it to the person array

  else if /things/i/city is empty
    make a place
    append it to the places array

  else 
    end of array, break out of loop

end loop
```

