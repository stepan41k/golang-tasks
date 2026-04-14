if err == io.EOF { ... }

// Способ 2
if errors.Is(err, io.EOF) { ... }