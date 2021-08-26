maps in go
1. keys and values are exact same types (all values must be the same type, and all keys must be the same type)

Difference between maps and structs

|Maps|Structs|
|----------|-------------|
|All keys must be the same type||
|Use to represent a collection of related properties|Use to represent a "thing" with a lot of different properties|
|All values must be the same type|Values can be of different type|
|Keys are indexed - we can iterate over them|Keys don't support indexing|
|Don't need to know all the keys at compile time|You need to know all the different fields at the compile time|
|Reference Type!|Value Type!|

