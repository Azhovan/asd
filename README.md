### Developer Driven TODO App


Examples: 
```bash 
# Create a new context with name prefixed with init-*. 
# It also sets the current context to it

# this is useful to start fresh new context to write your notes 
# if this meant to be a temporary notes.
# this kind of notes won't be garbage collected unless created with --gc flag 
asd init 

# Export current active Note as Markdown file in current directory
asd get 

# Export a Note with provided identifer  as Markdown file in current directory
asd get --id [note_id]

# Update current active note's title
asd add title [t]

# Update the note's title with the provided identifier
asd add title [t] --id [note_id]

```