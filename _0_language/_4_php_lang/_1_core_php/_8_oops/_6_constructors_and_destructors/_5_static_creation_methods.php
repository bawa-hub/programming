<?php

// PHP only supports a single constructor per class. In some cases, however, it may be desirable to allow an object to be constructed in different ways with different inputs. The recommended way to do so is by using static methods as constructor wrappers. 

class Product
{

    private ?int $id;
    private ?string $name;

    private function __construct(?int $id = null, ?string $name = null)
    {
        $this->id = $id;
        $this->name = $name;
    }

    public static function fromBasicData(int $id, string $name): static
    {
        $new = new static($id, $name);
        return $new;
    }

    public static function fromJson(string $json): static
    {
        $data = json_decode($json);
        return new static($data['id'], $data['name']);
    }

    public static function fromXml(string $xml): static
    {
        // Custom logic here.
        $data = convert_xml_to_array($xml);
        $new = new static();
        $new->id = $data['id'];
        $new->name = $data['name'];
        return $new;
    }
}


$p1 = Product::fromBasicData(5, 'Widget');
$p2 = Product::fromJson($some_json_string);
$p3 = Product::fromXml($some_xml_string);
