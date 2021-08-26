| address   |      value   |
|----------|:-------------:|
|0001|person{firstName: "alex"..}|

Turn `address` into `value` with `*address`

Turn `value` into `address` with `&value`

Data types

|Value Types|Reference Types|
|----------|:-------------:|
|int|slices|
|float|maps|
|string|channels|
|bool|pointers|
|structs|functions|

Value types need to be changed using pointers in a functions.