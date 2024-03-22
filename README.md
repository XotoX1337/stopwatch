
[![Go Reference](https://pkg.go.dev/badge/github.com/XotoX1337/stopwatch.svg)](https://pkg.go.dev/github.com/XotoX1337/stopwatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/XotoX1337/stopwatch)](https://goreportcard.com/report/github.com/XotoX1337/stopwatch)

# stopwatch
stopwatch is a small package to easily profile code in go.

## Download
```
go get github.com/XotoX1337/stopwatch
```

## Examples

### Single Section
```go
import "github.com/XotoX1337/stopwatch"

watch := stopwatch.New()
section, err := watch.Start("foo")
if err != nil {
    fmt.Println(err)
}
//code to profile
section, err = section.Stop("foo")
if err != nil {
    fmt.Println(err)
}
fmt.Printf("execution took %s", section.Duration())
```

### Time Laps within Sections
```go
import "github.com/XotoX1337/stopwatch"

watch := stopwatch.New()
section, err := watch.Start("foo")
if err != nil {
    fmt.Println(err)
}
//code to profile
section.Lap()
//code to profile inside section "foo"
section.Lap()
//additional code to profile inside section "foo"

section, err = section.Stop("foo")
if err != nil {
    fmt.Println(err)
}
fmt.Printf("execution took %s", section.Duration())
for i, lap := range section.Laps() {
    fmt.Printf("%d. Lap took %s", i+1, lap.Duration())
}
```

