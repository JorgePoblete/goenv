# goenv
Golang config loader from environment variables

The main goal of this is to make you worry about your configs, your conf structs and that kind of things and not in how to read each new variable that you add to your config.

This `load` function receives a reflect.Value of a struct and set its fields according to the tags of the fields, so each new conf variable or substruct that you need, you only need to add it and its corresponding tags, and thats it! the confs will be loaded from the environment by pure reflect magic.
