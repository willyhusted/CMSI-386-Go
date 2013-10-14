func powers(base int, limit int, callback func(int)) {
    result := 1
        for result <= limit {
            callback(result)
            result = result * base
            }
}
