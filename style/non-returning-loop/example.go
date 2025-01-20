package example

// ruleid: detect-potential-infinite-loops
for {
    fmt.Println("Infinite loop")
}

for {
    if condition {
        break
    }
}

for {
    return
}
