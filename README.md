# Infinity- literally everything digital appears in the set of natural numbers


Anything that can be described as a sequence of bytes can be described as a single, very large integer.
Just take the bytes and stack them end to end. You can interpret these bytes as a single number.
That number appears in the set of [Natural Numbers](https://en.wikipedia.org/wiki/Natural_number).

So in a sense, everything that can be expressed digitally (software programs, books, music,
pictures, etc) **already exists**- we just have to find it on the number line.

This means that for example, the Bible appears in the set of numbers- all possible versions of it. 
Same for the Koran, and every other book that's ever been written (or will be written). 
Tomorrow's newspaper is in there. Your photographs and your favorite movie are all in there too.


![creation of adam](https://github.com/armhold/infinity/blob/master/samples/creation_of_adam.jpg "creation of adam")


### Encoding

We can describe a simple binary encoding for any file of bytes:

1. Each binary digit (a 0 or 1) represents an increasing power of 2. So the first bit
is the number of ones. The second bit, the numbers of twos. The third bit, the number of fours, etc. 
1. The last byte in the file holds the lowest-order bits, and the first byte the highest.
1. the rightmost byte (last one in the file) holds bits 0-7, which represent those powers of 
two that can represent the decimal numbers 0-255.
1. the next-to-last byte holds bits 8-15, which represent 256-65,535.
1. and so on.
1. Since each additional bit **doubles** the size, this quickly leads to some pretty big numbers.


This quickly leads to some pretty big numbers.

A kilobyte of data (1024 bytes- not at all
a big file) is 8,192 bits. To get a ballpark idea of where that data lies in the number line
we can assume it's all ones, and take the maximum value of 8,192 bits. That's 2^8192, or
**1.0907481356×10^2466**.

That's more than the number of atoms in the universe, but it's still a number.