Initially, the template method pattern is a design that defines a template method function in abstract class and concrete implementation with concreate class.
However, since Go does not have an abstract class, we separate the responsibility of defining the template method and the responsibility to describe concrete processing.

* display
template method A structure for the function.
DisplayStrategy intarerface is embedded.

* DisplayStrategy(interface)  
It defines behavior related to display.

* CharDisplay  
An implementation of DisplayStrategy.

* StringDisplay  
An implementation of DisplayStrategy.
