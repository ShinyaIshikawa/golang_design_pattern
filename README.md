<H1>Design Pattern in Go</H1>

<H2>Overview</H2>
Design pattern impletation of Go programming.
Design pattern is based on [Introduction to design patterns learned in Java language](https://www.amazon.co.jp/exec/obidos/ASIN/4797327030/hyuki-22/) written by Hiroshi Yuki.

## Description
Go does not have the concept of inheritance, so we can not use the design pattern with inheritance.
When asking design pattern using inheritance, template method pattern may come to mind.
In Go, the template method pattern is expressed using duck typing by interface and structure embedding.

Abstract classes are frequently defined in object-oriented languages ​​such as Java, but Go thinks that it is best practice to keep interface usage to a minimum at go.


<H3>Itarator　Pattern (to count one by one) </H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/iterator>

<H3>Adapter Pattern (To cover the skin)</H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/adapter>

<H3>Template Method Pattern(Leave concrete processing to subclass)</H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/template>
* Realize template method pattern in object composition.

## Requirement

## License
MIT
## Author

[Shinya Ishikawa](https://github.com/ShinyaIshikawa)
