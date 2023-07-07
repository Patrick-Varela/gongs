# gongs

A Go package for the ngspice shared library.

# About

This is a very experimental package that uses cgo to communicate and dynamically link to the ngspice shared library written in C.
It provides a rudimentary way to create multiple instances of the library,
making it possible to be used with a web application where each user can have its own ngspice space.

# Usage

First you need to place the ngspice shared library (i.e. libngspice.so or ngspice.dll) in the same folder as your project's root folder.
You can also use spinit to load models, like you would do for ngspice, just put it inside `share/ngspice/scripts`.
To use the package in your go project, follow these steps:

1. Create a callback struct and define its attributes.
2. This struct is used to send the callbacks that receive data from the ngspice library in a synchronous way.

```Go
cbs := new(gongs.NgCallbacks)
```

Here's an example of a callback:

```Go
cbs.SendChar = func(i string, b1 int, r gongs.NgIReturn) int {
  fmt.Printf("NGSPICE: %s\n", i)
  return 2 + 2
}
```

Note that `gongs.NgIReturn` is an interface used to identify the callback's caller instance. 2. Initialize the gongs instance with the previously created `cbs` and clear the cache folder:

```go
ng := gongs.Init(cbs)
gongs.ClearCache()
```

3. Use the package as you would use the C ngspice shared library:

```go
ng.Circ([]string{
  ".include 32nm_HP.pm",
  "Vvdd vdd gnd 1",
  "Va a gnd PWL (0n 0 10n 0 10.01n 1 20n 1)",
  "Mp1 vdd a inv vdd PMOS w=140n l=32n",
  "Mn1 inv a gnd gnd NMOS w=70n l=32n",
  ".end"})
ng.Command("tran 1ns 20ns 0ns")
```

That's it! Now the callback function will return the information you need. You can also use asynchronous functions like `func (ngspice *NgSpice) AllPlots() []string`,
which work exactly like the official library.

# Known Issue

There is a bug that I haven't been able to pinpoint yet. Sometimes, after using an instance,
the garbage collector moves something we are using, causing the program to crash.
I'll look futher into it later

Also, at the moment, `32nm_HP.pm` is always included.

# License

This project is licensed under the MIT License. See the LICENSE file for details.
