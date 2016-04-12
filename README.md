# Infinity

Literally everything digital appears in the set of natural numbers.


### Let's find a famous painting in the set of Natural Numbers

![creation of adam](https://github.com/armhold/infinity/blob/master/samples/creation_of_adam.jpg "creation of adam")

Can we find a famous painting hidden among the integers in the set of Natural Numbers?
It turns out we can- Michelangelo's [Creation of Adam](https://en.wikipedia.org/wiki/The_Creation_of_Adam)
happens to reside at exactly this number:

    25519637684623283551604138549800901322256540519555396335820617899064263322699621
    95191687493347723876911325592757091527584399836110327600568809057692273948571948
    08265352774295543927463930227303989499812656204789274349264352246002620646195049
    80743234655174227362545608569251825909819314291568940243800729262655548989815338
    54946158978193172722003989771610404779019733491375454052733009346842568422009696
    74942993143476565213005563390876806416489284574517449719256327217688989948909567
    [ ... ]
    66075392419937785208072076024225421796258819260095866487851661074765988817906435
    95488615051785336967485705800344694295344605450895377295574423161887490912578460
    55236362647031265463447554894035823105547224316844121735586980345834721507867209
    39588739309981549612372297190758970535136257561299604029885537909193046058541165
    07892131487938814207003552600673252186494827588209195808322307452037885128631910
    361


NB: I've elided most of the lines for brevity; you can download the actual file [here](https://github.com/armhold/infinity/blob/master/samples/creation_of_adam-number.txt).)

That is a number with 160,165 digits. It's a simple (though long) integer on the number line,
and it happens to contain Michelangelo's Creation of Adam. It's always been there, even
before Michelangelo created his famous fresco in 1511.



### Background

Anything that can be described as a sequence of bytes can also be described by a single, 
very large integer. Just take the bytes and stack them end to end. You can interpret these 
bytes as a single integer. By definition, that integer appears in the set of
[Natural Numbers](https://en.wikipedia.org/wiki/Natural_number).

So in a sense, everything that can be expressed digitally (software programs, books, music,
pictures, etc) **already exists**- we just have to find it on the number line.

This means that, for example, the Bible appears in the set of numbers- all possible versions of it. 
Same for the Koran, and every other book that's ever been written (or will be written). 
Tomorrow's newspaper is in there. **Your photographs** and **your favorite movie** are all in there too.

This program explores how we can discover the numbers associated with digital files.


### Representation


1. Any file contains a sequence of bytes- we string these bytes together to create a single number.
1. Each byte has eight bits of storage.
1. Every bit doubles the potential "range" for the number.
1. Since a byte has eight bits, each time you add a byte, the range for the number
is doubled eight times.

This quickly leads to some pretty big numbers.

A kilobyte of data (1024 bytes- not at all
a big file) is 8,192 bits. To get a ballpark idea of where that data lies on the number line
we can assume it's all ones, and take the maximum value of 8,192 bits. That's 2^8192, or
**1.0907481356×10^2466**.

That's more than the number of atoms in the universe, but it's still a number.

### Demonstration

Consider the string "ABC". This can be written as three bytes. In decimal, the bytes are (via ASCII/Unicode encoding)
65, 66, and 67. In binary, this gives us:

    A = 65 = 01000001
    B = 66 = 01000010
    C = 67 = 01000011

If we concatenate the binary, we get:

    01000001 01000010 01000011  
     
Or 4,276,803 in decimal. So the number for the string "ABC" is 4,276,803. Therefore "ABC" exists
on the number line near the 4.2 millionth number.

This project includes a program to print the number line "coordinate" for any file. Run it
like this:

    $ go install github.com/armhold/infinity/cmd/find_number
    $ find_number samples/creation_of_adam.jpg
    
Or to just search for a string without using a file:

    $ find_number -t "meaning of life"
      568016600252838567935837243227858533



### Notes

The image "Creation of Adam" by Michelangelo is in the [Public Domain](https://commons.wikimedia.org/w/index.php?curid=15461165).
