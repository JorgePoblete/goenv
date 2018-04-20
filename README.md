# goenv
Golang config loader from environment variables

The main goal of this is to make you worry about your configs, your conf structs and that kind of things and not in how to read each new variable that you add to your config.

This `load` function receives a reflect.Value of a struct and set its fields according to the tags of the fields, so each new conf variable or substruct that you need, you only need to add it and its corresponding tags, and thats it! the confs will be loaded from the environment by pure reflect magic.

Example
```
$ export SOME_STRING=someStringDefinition
$ ./goenv
Config for SOME_OTHER missing
Config for SUB_BOOL missing
Config for SUB_SOME_OTHER missing

Readed conf: {Some:{VarString:someStringDefinition VarInt:1313 VarFile:} SubSome:{ImBool:false Some:{VarString:ClassicStringIsClassic VarInt:1313 VarFile:}} Other:{VarString: VarInt:0}}
unset SOME_STRING
```
If the enviroment var is not set, try to use a FILE but if this is not defined use default, 
and  finally if default was not set only print a message

To use the file configuration, you need create a file like config_file
```
$ export SOME_OTHER_FILE=config_file
$ ./goenv
Config for SUB_BOOL missing
Config for SUB_SOME_OTHER missing

Readed conf: {Some:{VarString:ClassicStringIsClassic VarInt:1313 VarFile:dummy_conf} SubSome:{ImBool:false Some:{VarString:ClassicStringIsClassic VarInt:1313 VarFile:}} Other:{VarString: VarInt:0}}
```
