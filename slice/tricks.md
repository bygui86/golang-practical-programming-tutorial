
# [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)

Since the introduction of the ` append ` built-in, most of the functionality of the ` container/vector ` package, which was removed in Go 1, can be replicated using ` append ` and ` copy `.

Here are the vector methods and their slice-manipulation analogues:

#### AppendVector
```go
a = append(a, b...)
```

#### Copy
```go
b = make([]T, len(a))
copy(b, a)
// or
b = append([]T(nil), a...)
// or
b = append(a[:0:0], a...)  // See https://github.com/go101/go101/wiki
```

#### Cut
```go
a = append(a[:i], a[j:]...)
```

#### Delete
```go
a = append(a[:i], a[i+1:]...)
// or
a = a[:i+copy(a[i:], a[i+1:])]
```

#### Delete without preserving order
```go
a[i] = a[len(a)-1] 
a = a[:len(a)-1]

```
**NOTE** If the type of the element is a _pointer_ or a struct with pointer fields, which need to be garbage collected, the above implementations of ` Cut ` and ` Delete ` have a potential _memory leak_ problem: some elements with values are still referenced by slice ` a ` and thus can not be collected. The following code can fix this problem:
> **Cut**
```go
copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k++ {
	a[k] = nil // or the zero value of T
}
a = a[:len(a)-j+i]
```

> **Delete**
```go
copy(a[i:], a[i+1:])
a[len(a)-1] = nil // or the zero value of T
a = a[:len(a)-1]
```

> **Delete without preserving order**
```go
a[i] = a[len(a)-1]
a[len(a)-1] = nil
a = a[:len(a)-1]
```

#### Expand
```go
a = append(a[:i], append(make([]T, j), a[i:]...)...)
```

#### Extend
```go
a = append(a, make([]T, j)...)
```

#### Filter (in place)

```go
n := 0
for _, x := range a {
	if keep(x) {
		a[n] = x
		n++
	}
}
a = a[:n]
```

#### Insert
```go
a = append(a[:i], append([]T{x}, a[i:]...)...)
```
**NOTE** The second ` append ` creates a new slice with its own underlying storage and  copies elements in ` a[i:] ` to that slice, and these elements are then copied back to slice ` a ` (by the first ` append `). The creation of the new slice (and thus memory garbage) and the second copy can be avoided by using an alternative way:
> **Insert**
```go
s = append(s, 0 /* use the zero value of the element type */)
copy(s[i+1:], s[i:])
s[i] = x
```

#### InsertVector
```go
a = append(a[:i], append(b, a[i:]...)...)
```

#### Push
```go
a = append(a, x)
```

#### Pop
```go
x, a = a[len(a)-1], a[:len(a)-1]
```

#### Push Front/Unshift
```go
a = append([]T{x}, a...)
```

#### Pop Front/Shift
```go
x, a = a[0], a[1:]
```

## Additional Tricks
### Filtering without allocating

This trick uses the fact that a slice shares the same backing array and capacity as the original, so the storage is reused for the filtered slice. Of course, the original contents are modified.

```go
b := a[:0]
for _, x := range a {
	if f(x) {
		b = append(b, x)
	}
}
```

For elements which must be garbage collected, the following code can be included afterwards:

```go
for i := len(b); i < len(a); i++ {
	a[i] = nil // or the zero value of T
}
```

### Reversing

To replace the contents of a slice with the same elements but in reverse order:
```go
for i := len(a)/2-1; i >= 0; i-- {
	opp := len(a)-1-i
	a[i], a[opp] = a[opp], a[i]
}
```
The same thing, except with two indices:
```go
for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
	a[left], a[right] = a[right], a[left]
}
```

### Shuffling

Fisher–Yates algorithm:

> Since go1.10, this is available at [math/rand.Shuffle](https://godoc.org/math/rand#Shuffle)

```go
for i := len(a) - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    a[i], a[j] = a[j], a[i]
}
```

### Batching with minimal allocation

Useful if you want to do batch processing on large slices.

```go
actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
batchSize := 3
batches := make([][]int, 0, (len(actions) + batchSize - 1) / batchSize)

for batchSize < len(actions) {
    actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)
```

Yields the following:
```go
[[0 1 2] [3 4 5] [6 7 8] [9]]
```

### In-place deduplicate (comparable) 

```go
import "sort"

in := []int{3,2,1,4,3,2,1,4,1} // any item can be sorted
sort.Ints(in)
j := 0
for i := 1; i < len(in); i++ {
	if in[j] == in[i] {
		continue
	}
	j++
	// preserve the original data
	// in[i], in[j] = in[j], in[i]
	// only set what is required
	in[j] = in[i]
}
result := in[:j+1]
fmt.Println(result) // [1 2 3 4]
```