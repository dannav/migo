# Migo
An opinionated helper library for rendering go `html/template` files.

## In Progress
This project is currently in progress.

Todo
- write to any writer interface{} (not just http.ResponseWriter)
- memoize template parsing
- allow any file extension (currently only searches for `.tmpl`)
- allow layout file name to be defined (default is layout)

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

Using migo you can render the signin template under the account directory by calling `r.Render(rw, "account/signin", data)` where `rw` is an `http.ResponseWriter` and data satisfies type `interface{}`.

## Opinionated
If you'd like a directory of template files to share a common layout, create a shared folder and file `layout.tmpl` which will be the layout for that directory. All directories that don't have `shared/layout.tmpl` will inherit the layout defined in `{template_base_path}/shared/layout.tmpl`.

Currently a base layout must be defined in `{template_base_path}/shared/layout.tmpl`.
