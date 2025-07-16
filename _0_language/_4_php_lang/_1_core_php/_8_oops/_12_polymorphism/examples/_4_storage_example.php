<?php

// Storage Interface
interface StorageInterface
{
    public function upload($file, $destination);
    public function download($file);
    public function delete($file);
}

// Local Storage class implementing the interface
class LocalStorage implements StorageInterface
{
    public function upload($file, $destination)
    {
        echo "Uploaded $file to local destination: $destination\n";
        // Local storage-specific upload logic
    }

    public function download($file)
    {
        echo "Downloaded $file from local storage\n";
        // Local storage-specific download logic
    }

    public function delete($file)
    {
        echo "Deleted $file from local storage\n";
        // Local storage-specific deletion logic
    }
}

// Cloud Storage class implementing the interface
class CloudStorage implements StorageInterface
{
    public function upload($file, $destination)
    {
        echo "Uploaded $file to cloud destination: $destination\n";
        // Cloud storage-specific upload logic
    }

    public function download($file)
    {
        echo "Downloaded $file from cloud storage\n";
        // Cloud storage-specific download logic
    }

    public function delete($file)
    {
        echo "Deleted $file from cloud storage\n";
        // Cloud storage-specific deletion logic
    }
}

// NAS (Network-Attached Storage) class implementing the interface
class NASStorage implements StorageInterface
{
    public function upload($file, $destination)
    {
        echo "Uploaded $file to NAS destination: $destination\n";
        // NAS storage-specific upload logic
    }

    public function download($file)
    {
        echo "Downloaded $file from NAS storage\n";
        // NAS storage-specific download logic
    }

    public function delete($file)
    {
        echo "Deleted $file from NAS storage\n";
        // NAS storage-specific deletion logic
    }
}

// Client code using polymorphism to switch between storage systems
function performStorageOperations(StorageInterface $storage, $file, $destination)
{
    $storage->upload($file, $destination);
    $storage->download($file);
    $storage->delete($file);
}

// Example usage
$localStorage = new LocalStorage();
performStorageOperations($localStorage, "file.txt", "/local/path");

$cloudStorage = new CloudStorage();
performStorageOperations($cloudStorage, "document.pdf", "/cloud/storage");

$nasStorage = new NASStorage();
performStorageOperations($nasStorage, "image.jpg", "/nas/drive");


// The StorageInterface defines common methods for uploading, downloading, and deleting files.
// The LocalStorage, CloudStorage, and NASStorage classes implement the interface with their specific logic for each operation.
// The performStorageOperations function takes a StorageInterface as a parameter, allowing it to work with any class that implements the interface.

// This example demonstrates how you can use polymorphism to handle different storage systems without modifying the client code that interacts with these storage classes. It makes it easy to switch between local storage, cloud storage, and network-attached storage without affecting the overall structure of the code.