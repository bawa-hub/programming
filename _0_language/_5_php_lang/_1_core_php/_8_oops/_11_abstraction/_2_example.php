<?php

interface Content
{
    public function getTitle();
    public function getContent();
    public function display();
}

class Article implements Content
{
    private $title;
    private $body;

    public function __construct($title, $body)
    {
        $this->title = $title;
        $this->body = $body;
    }

    public function getTitle()
    {
        return $this->title;
    }

    public function getContent()
    {
        return $this->body;
    }

    public function display()
    {
        echo "<h2>{$this->title}</h2>";
        echo "<p>{$this->body}</p>";
    }
}

class Image implements Content
{
    private $url;
    private $description;

    public function __construct($url, $description)
    {
        $this->url = $url;
        $this->description = $description;
    }

    public function getTitle()
    {
        return "Image";
    }

    public function getContent()
    {
        return "<img src='{$this->url}' alt='{$this->description}' />";
    }

    public function display()
    {
        echo $this->getContent();
        echo "<p>{$this->description}</p>";
    }
}

class ContentManager
{
    private $contents = [];

    public function addContent(Content $content)
    {
        $this->contents[] = $content;
    }

    public function displayAll()
    {
        foreach ($this->contents as $content) {
            $content->display();
            echo "<hr>"; // Add a horizontal line between content items for separation
        }
    }
}

// Using abstraction in the Content Management System
$contentManager = new ContentManager();

// Add content items
$article = new Article("Understanding Abstraction", "Abstraction is a key concept in OOP.");
$image = new Image("image.jpg", "An example image");

$contentManager->addContent($article);
$contentManager->addContent($image);

// Display all content items
$contentManager->displayAll();

// The Content interface defines a common set of methods (getTitle, getContent, display) for different types of content.
// Concrete classes (Article, Image) implement this interface with specific implementations for title, content, and display.
// The ContentManager class uses polymorphism and abstraction by working with objects that adhere to the common Content interface. It can add and display different types of content without knowing their specific implementations.

// This abstraction allows the content management system to handle various content types uniformly, making it easy to extend and maintain as new content types are added.