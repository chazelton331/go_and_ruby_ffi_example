### Go + Ruby via FFI
##### (or Go by itself via CLI)

Run the script `build_go_library.sh`, which will produce the library code (`libbongo.so` and `libbongo.h`) and an executable (`bon`).

Run the executable via the command line, like so:

```
bon 1 2 3 4 5 6
# => 21
```

Or run the code through the Ruby program `sum.rb`

```
ruby sum.rb
# => 300
```
