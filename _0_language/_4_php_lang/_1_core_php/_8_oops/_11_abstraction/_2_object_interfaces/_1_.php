<?php

/**
 * Object interfaces allow you to create code which specifies which methods a class must implement, 
 * without having to define how these methods are implemented.
 * 
 *  Interfaces share a namespace with classes and traits, 
 * so they may not use the same name.
 * 
 * All methods declared in an interface must be public; 
 * this is the nature of an interface. 
 * 
 * In practice, interfaces serve two complementary purposes: 
 * 1. To allow developers to create objects of different classes that may be used interchangeably 
 * because they implement the same interface or interfaces
 *  A common example is java List,  multiple database access services, multiple payment gateways, or different caching strategies
 *  Different implementations may be swapped out without requiring any changes to the code that uses them. 
 * 
 * 2. To allow a function or method to accept and operate on a parameter that conforms to an interface, 
 * while not caring what else the object may do or how it is implemented.
 * These interfaces are often named like Iterable, Cacheable, Renderable, 
 * or so on to describe the significance of the behavior
 * 
 * Interfaces may define magic methods to require implementing classes to implement those methods. 
 *  Classes may implement more than one interface if desired by separating each interface with a comma
 * 
 * An interface, together with type declarations, provides a good way to make sure that a particular object contains particular methods
 */

//  interface example
interface Template
{
    public function setVariable($name, $var);
    public function getHtml($template);
}

// Implement the interface
// This will work
class WorkingTemplate implements Template
{
    private $vars = [];

    public function setVariable($name, $var)
    {
        $this->vars[$name] = $var;
    }

    public function getHtml($template)
    {
        foreach ($this->vars as $name => $value) {
            $template = str_replace('{' . $name . '}', $value, $template);
        }

        return $template;
    }
}

// This will not work
// Fatal error: Class BadTemplate contains 1 abstract methods
// and must therefore be declared abstract (Template::getHtml)
class BadTemplate implements Template
{
    private $vars = [];

    public function setVariable($name, $var)
    {
        $this->vars[$name] = $var;
    }
}
