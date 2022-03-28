# Go from the beginning - arrays and slices

An array is group of elements that are connected. You want to use an array when you have a group of something, like many orders, cars, rows in a file. The idea is to collect all that data in one structure. You also want to be able to iterate over it and possibly carry out an operation on it as a group.

## Declare an array

To declare an array, you need to specify the following properties:

- **capacity**, how many elements it holds.
- **type**, what type of elements it holds.
- **array content**, you can assign it elements at creation or do so later.

Here's the syntax:

```go
[<capacity>]<type>{...element}
```

It starts with the square brackets, `[]`. Within the square brackets you set the capacity, how many elements it can hold.

and here's a more real example:

```go
cities := [5]string{"NY", "LA"}
```

In the preceding code, an array of strings is declared. It has capacity for 5 elements and two of the places are filled with "NY" and "LA". Note also because we set the capacity to 5 and the number of elements it's assigned is 2, there's 3 spaces free.

### Capacity by inference

You don't have to set capacity to an explicit number, you can set it to `...`, in which case the capacity will be set to the number of elements you assign to it, like so:

```go
ids := [...]int{1, 2, 3, 4}
```

The preceding code has 4 elements, and that's also its capacity.

## Accessing elements

The way to access an element is by using it's index. Index is 0 based, meaning the first index is 0 and its last is the length -1. 

```go
ids := [...]int{1, 2, 3, 4}
ids[0] // 1
ids[3] // 4
```

## Length and capacity

Imagine we have the following array declared:

```go
cities := [5]string{"NY", "LA"}
```

- **length**. The length is defined as the number of elements in the array. You can use the `len()` method to find this out:

   ```go
   len(cities) // 2
   ```

- **capacity**. The capacity is how many elements the array can hold. `cap()` is the method you use to find the capacity:

   ```go
   cap(cities) // 5
   ```

## Slices

A slice is a part of an array. A slice is created when the slice operator is being used. Here's the syntax for the slice operator:

```go
s[i:p]
```

- `s`, the array
- `i`, the first index of the array to take elements from
- `p`, The variable p corresponds to the last element in the underlying array that can be used in the new slice. I.e cut right before this index.

```go
items := [5]int{1,2,3,4,5}
part = items[1:3] // 2,3
```

### Adding elements

A slice differs from an array, you can add items to it. The `append()` method lets you add elements to it. The syntax for `append()` is as follows:

```go
append(slice, element)
```

Here's how you can append to a slice:

```go
var numbers []int
numbers = append(numbers, 1)
numbers = append(numbers, 2) // 1,2
```

### Removing elements

Remove an element by constructing a new slice.

```go
letters := []string{"A", "B", "C", "D", "E"}
remove := 2 // remove index
// 0 - remove index, remove +1 to end   
letters = append(letters[:remove], letters[remove+1:]...)
```

### Create a slice with make

You can use the `make()` method to create a slice. Here's how:

```go
slice := make([]int, 5) // creates a slice with length 5 and capacity 5
```

You can set these to different values:

```go
slice2 := make([]int, 2, 5)
fmt.Println(slice2)
fmt.Println(len(slice2))
fmt.Println(cap(slice2))
```

Here, the slice has length 2 and capacity 5.

### Copy elements

```go
arr := [3]int{1, 2, 3}
dest := make([]int, 5)
copy(dest, arr[0:2]) // copies slice {1,2} into dest
fmt.Println(dest) // [1 2 0 0 0]
```

## Summary

In summary, you were taught how to use Arrays. Additionally, we covered working with slices. Slices allows you to work with an array like structure and change its size as you see fit.

### Learn more

<https://docs.microsoft.com/en-us/learn/modules/go-data-types/2-slices>
