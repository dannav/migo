# Migo
An opinionated helper library for rendering go `html/template` files.

## NOTE: In Progress
This project is currently in progress.

TODO
- write to any writer interface{} (not just http.ResponseWriter)
- memoize template parsing

## Use
Migo allows you to render go `html/template` files with some syntactic sugar. Migo is opinionated with folder structure but allows you to define a base path where your template files are stored.

Case: a `templates` folder with the following files/directories:

- account
  - signin.tmpl
  - signup.tmpl
  - shared
    - layout.tmpl
- home
  - index.tmpl
- shared
  - layout.tmpl


Instantiate Migo
```go
  r := migo.New("templates")
```

Using migo you can render the signin template under account by calling `r.Render(rw, "account/signin")` where `rw` is an `http.ResponseWriter`.
