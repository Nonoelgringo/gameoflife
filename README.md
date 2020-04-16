# gameoflife
Personnal Go implementation of Conway's Game of Life.
For the rules I used the [Computerphile video](https://www.youtube.com/watch?v=E8kUJL04ELA) with John Conway himself.
  I haven't looked at other Game of Life programs (in any language) before doing this. Only allowed myself to look for help concerning pure Go problems.
  
## Usage

**Requirements** : A Unix environment (Linux/FreeBSD/MacOS)

### seeds

seed needs to be a **txt** file using '-' for dead cells and any string for a living cell.

*glider.txt*

```
- - - - - - - - - - - - -
- - - - - - - - - - - - -
- - - - - O O O - - - - -
- - - - - - - O - - - - -
- - - - - - O - - - - - -
- - - - - - - - - - - - -
- - - - - - - - - - - - -
```

### arguments

- first argument : seed file (**mandatory**)
- second argument: milliseconds between prints (*optionnal*- default is 150)

### usage example

```
$./gameoflife glider.txt 50
```

## Todo / upgrades

- Windows support
- more seeds included
- look at other Go implementations

