### Developer Driven TODO App


Examples: 
```bash 
# Create a new context with name prefixed with init-*. 
# It also sets the current context to it

# this is useful to start fresh new context to write your notes 
# if this meant to be a temporary notes.
# this kind of notes won't be garbage collected unless created with --gc flag 
asd init 

[IDEA]
# A shorthand for asd init command 
asd -i 

# Create a context with a custom name
asd -i --name catgeory1

# A shorthand for asd init --name command  
asd -i -n category1 

# [IDEA]
# Deletes notes after 10 hour
asd -i --gc=10h 

# Print current context in standard output
asd context

# Set the current context to category1, an error is returned if category1 
# hasn't been defined already. 
asd --context=category1

# A shorthand for --context command 
asd -c category1
```