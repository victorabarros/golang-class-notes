# Chapter 3 - Basic Data Typer

https://learning.oreilly.com/library/view/the-go-programming/9780134190570/ebook_split_027.html

## 3.1 Integers

> There are four distinct sizes ofsigned integers—8, 16, 32, and 64 bits—represented by the types int8, int16, int32, and int64, and corresponding unsigned versions uint8, uint16, uint32, and uint64.
> The type rune is a synonym for int32 and conventionally indicates that a value is a Unicode code point.
> Similarly, the type byte is an synonym for uint8.
> The range ofvalues of an n-bit number is from `−2**(n−1)` to `2**(n−1)−1`.
> The behavior of `/` depends on whether its operands are integers,so 5.0/4.0 is 1.25, but 5/4 is 1 because integer division truncates the result toward zero.

## 3.2 Floating-Point Numbers
