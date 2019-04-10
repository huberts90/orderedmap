# orderedmap
Ordered map with JSON utility

# Usage
```Go
dict := OrderedMap{
    {"firstName", "Julius"},
    {"lastName", "Caesar"},
    {"age", 50},
}
json, err := json.Marshal(dict)

if err != nil {
    fmt.Printf("JSON for order map error: %s", err)
} else {
    fmt.Printf("JSON %s", string(json))
}
```
