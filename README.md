## Monmind - obfuscate multiple strings & hide text from binary searching 
Obfuscation strings in golang code

<p align="center">
	<img src="monmind.png" alt="kill"/>
</p>



## INSTALL

You can install monmind by running:

```
go get  github.com/CRYBOII/monmind
```
 


## USAGE

Once [monmind](https://github.com/CRYBOII/monmind) is installed 
go to workspacd folder and create .ob file, 

and slap your strings in there on different lines.

now you can run this command 

```
monmind
```

this should be generate the file `obfascated.go`
all encoded strings will be contians in that file
 
### Example 

in .ob file 
```
this is secret
this is super secret
```
when we run `monmind` obfuscated file is generated
```
package main

import (
	"unsafe"
)

func eax() uint8 {
	return uint8(unsafe.Sizeof(true))
}

//  [ YLM01E5A6 ] ====>  this is secret

func YLM01E5A6() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			eax() << eax() << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			eax() << eax() << eax() << eax() << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ YLMEECE42 ] ====>  this is super secret

func YLMEECE42() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			eax() << eax() << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			eax() << eax() << eax() << eax() << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
			eax() << eax() << eax() << eax() << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

```

now you could use a encoded functions string like this

`fmt.Println(YLMEECE42())`

or with different package 


`fmt.Println(bar.YLMEECE42())`

#### this will print 
```
this is super secret
```
#### without obfuscated

<p align="center">
	<img src="capture-20211103-160212.png" alt="kill"/>
</p>

#### with obfuscated

<p align="center">
	<img src="capture-20211103-160509.png" alt="kill"/>
</p>

## Acknowledgments

* Thanks [jeromer](https://github.com/jeromer) 
