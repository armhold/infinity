# Infinity- literally everything digital appears in the set of natural numbers


Anything that can be described as a sequence of bytes, can be described as a single, very large integer.
Just take the bytes and stack them end to end. You can interpret these bytes as a
single, very large integer. That integer appears in the set of Natural Numbers.

So everything that **can be** expressed digitally (software programs, books, music,
pictures, etc) **already exists**- we just have to find it on the number line.

This means that for example, the Bible appears in the set of numbers- all
possible versions of it. Same for the Koran, and every other book that's ever
been written (or will be written). Tomorrow's newspaper is in there. Your photographs,
and all your home movies are in there too.



### Encoding

There are multiple ways to handle the encoding. The simplest
is probably just to interpret them the way computers already do:

1. the leftmost byte has bits 0-7, which represent those powers of two that can represent 0-255
1. the next byte after that has bits 8-15, which represent 256-65,535.
1. and so on.


This quickly leads to some pretty big numbers.

A kilobyte of data (1024 bytes- not at all
a big file) is 8,192 bits. To get a ballpark idea of where that data lies in the number line
we can assume it's all ones, and take the maximum value of 8,192 bits. That's 2^8192, or
**1.0907481356×10^2466**.

That's more than the number of atoms in the universe, but it's a valid number.