# Render markdown from blackfriday/v2

This is an attempt to parse and render markdown using blackfriday/v2.

It struggles to parse list items that don't contain spacing before or
after the list. As the intent is to parse poor markdown and fix the
rendering, using the package doesn't work nicely for more complex
documents. Noted issues were:

- prefixed text before a list item with a trailing colon does not get parsed correctly
- suffixed text to the last list item will be included within the list item text

It's sufficient for simple markdown files, however, given this is undesired
behaviour, a different library needs to be used to implement a wider scope.

The main problem with the library seems to be that it's design was intended
to render html. This means we would have to resort to `<ul>` and `<ol>` elements
whenever a list is ordered. To re-render markdown from this AST, several
input pieces of information are omitted, like indentation or counts. It's not
possible to restore the input markdown from the resulting AST.

The renderer is incomplete/unfinished/abandoned. It's missing support for
tables at the very least, and likely several other items.

Try not to use it.