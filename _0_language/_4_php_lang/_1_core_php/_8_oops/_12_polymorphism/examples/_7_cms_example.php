<?php

// Let's explore another example involving a content management system (CMS) where different types of content elements need to be displayed on a webpage.
//  We'll consider polymorphism to handle various content types, such as articles, images, and videos

// ContentElement Interface
interface ContentElement
{
    public function render();
}

// Article class implementing the interface
class Article implements ContentElement
{
    private $title;
    private $content;

    public function __construct($title, $content)
    {
        $this->title = $title;
        $this->content = $content;
    }

    public function render()
    {
        // Rendering logic for displaying an article
        echo "<h2>{$this->title}</h2>";
        echo "<p>{$this->content}</p>";
    }
}

// Image class implementing the interface
class Image implements ContentElement
{
    private $src;
    private $alt;

    public function __construct($src, $alt)
    {
        $this->src = $src;
        $this->alt = $alt;
    }

    public function render()
    {
        // Rendering logic for displaying an image
        echo "<img src='{$this->src}' alt='{$this->alt}' />";
    }
}

// Video class implementing the interface
class Video implements ContentElement
{
    private $url;

    public function __construct($url)
    {
        $this->url = $url;
    }

    public function render()
    {
        // Rendering logic for displaying a video
        echo "<iframe width='560' height='315' src='{$this->url}' frameborder='0' allowfullscreen></iframe>";
    }
}

// Page class using polymorphism for rendering content
class Page
{
    private $contentElements = [];

    public function addContentElement(ContentElement $element)
    {
        $this->contentElements[] = $element;
    }

    public function renderContent()
    {
        foreach ($this->contentElements as $element) {
            $element->render();
            echo "<hr>"; // Add a horizontal line between content elements for separation
        }
    }
}

// Client code using polymorphism to add and render different content elements
$page = new Page();

// Add content elements
$page->addContentElement(new Article("Introduction to Polymorphism", "Polymorphism is a key concept in OOP."));
$page->addContentElement(new Image("image.jpg", "An example image"));
$page->addContentElement(new Video("https://www.youtube.com/embed/example"));

// Render content
$page->renderContent();


// The ContentElement interface defines a common method render that all content elements must implement.
// Concrete classes (Article, Image, Video) implement this interface with their specific rendering logic.
// The Page class uses polymorphism by allowing the dynamic addition and rendering of different content elements.

// This example illustrates how polymorphism allows you to treat different types of content elements uniformly, making it easy to extend and maintain the rendering logic as new content types are added to the system. The Page class doesn't need to know the specific type of each content element; it can render them using a common interface.