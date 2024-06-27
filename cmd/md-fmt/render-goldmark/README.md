# Render markdown to markdown using goldmark

The package is being used by Hugo to provide HTML rendering.

Some reading:

- https://github.com/yuin/goldmark/
- https://github.com/shurcooL/markdownfmt/issues/56
- https://github.com/Kunde21/markdownfmt

CLI Installation:

```
go install github.com/Kunde21/markdownfmt/v3/cmd/markdownfmt@latest
```

Unfortunately, the cli tool doesn't support front-matter (hugo). 
The readme lists two other options:

## bwplotka/mdox

URL: https://github.com/bwplotka/mdox/

This seems to be a very configurable docs toolset.

## moorereason/mdfmt

URL: https://github.com/moorereason/mdfmt

Mdfmt uses shurcooL/markdownfmt to provide formatting.
This has two unfortunate issues:

- lists are not indented in a configurable way (8 spaces)
- uses blackfriday to implement a markdown renderer.

Front matter support was solved by using a dependency on `gohugoio/hugo/parser`.
It essentially reads in the markdown file with hugo to get the front-matter,
and only sends the markdown contents into markdownfmt.
