# Did I Wear It Already?

This is a side project created to get familiar with [BoltDB](https://github.com/boltdb/bolt) and use it for some simple data manipulation.

## Background

_Do you have anything to wash? I'm doing the laundry!_

Does this question terrify you as well? In a split second you need to figure out which of those _not-so-much-dirty_ clothes should actually end in in the washing machine. It's easy for underpants and socks (please! one day max!), but T-shirs or pants? Unless you put a stain on them, you're likely to want to wear it again, right? This is where `diwia` comes in!

## Installation

In order to fetch and build `diwia` from Go sources, simply type:

    $ go get -u github.com/slomek/diwia
    ...

    $ Usage:
    diwia [command]

    Available Commands:
    add         Adds piece of clothes
    help        Help about any command
    list        Lists all clothes
    wash        Clears clothes' wear number
    wear        Adds to clothes' wear number

    Flags:
    -h, --help   help for diwia

    Use "diwia [command] --help" for more information about a command.

## Usage

### Define your wardrobe

In order to add a new piece of clothing, use `add` command:

    # Add an item with its name
    $ diwia add pants.maroon
    A new item has been inserted

    # Add an item with its name and description
    $ diwia add pants.jeans_old "Worn-out jeans"
    A new item has been inserted

In order to list your clothes, use `list` command:

    $ diwia list
    Found 2 clothes:
    - pants.jeans_old (Worn-out jeans, wears=0)
    - pants.maroon (wears=0)

### Marking clothes as worn

In order to mark that you've had a particular piece of clothing on, use `wear` command:

    $ diwia wear pants.maroon
    No previous wears found for 'pants.maroon'
    A wear-number for item has been updated

    $ diwia wear pants.maroon
    A wear-number for item has been updated

    $ diwia wear pants.jeans_old
    No previous wears found for 'pants.jeans_old'
    A wear-number for item has been updated

    $ diwia list
    Found 2 clothes:
    - pants.jeans_old (Worn-out jeans, wears=1)
    - pants.maroon (wears=2)

### Marking clothes as clean

In order to select which clothes were washed (and their wear number should be reset), use `wash` command:

    $ diwia wash pants.jeans_old
    A wear counts for item has been reset for [pants.jeans_old]
    
    $ diwia wash pants.jeans_old pants.maroon
    A wear counts for item has been reset for [pants.jeans_old pants.maroon]
    
    $ diwia list
    Found 2 clothes:
    - pants.jeans_old (Worn-out jeans, wears=0)
    - pants.maroon (wears=0)
