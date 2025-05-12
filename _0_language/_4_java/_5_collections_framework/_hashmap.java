// https://www.programiz.com/java-programming/library/hashmap/put

import java.util.HashMap;
import java.util.Map.Entry;

public class _hashmap {
    public static void main(String[] args) {

        // HashMap class implements the Map interface.
        // HashMap<K, V> numbers = new HashMap<>();

        // create a hashmap
        HashMap<String, Integer> languages = new HashMap<>();

        // add elements to hashmap
        System.out.println("Initial HashMap: " + languages);
        languages.put("Java", 8);
        languages.put("JavaScript", 1);
        languages.put("Python", 3);
        System.out.println("HashMap: " + languages);

        // Access HashMap Elements
        int value = languages.get("Python");
        System.out.println("Value at index 1: " + value);
        System.out.println("Keys: " + languages.keySet()); // return set view of keys
        System.out.println("Values: " + languages.values()); // return set view of values
        System.out.println("Key/Value mappings: " + languages.entrySet()); // return set view of key/value pairs

        // Change HashMap Value
        languages.replace("Java", 0);
        System.out.println("HashMap using replace(): " + languages);

        // Remove HashMap Elements
        int val = languages.remove("Python");
        System.out.println("Removed value: " + val);
        System.out.println("Updated HashMap: " + languages);

        // Iterate through a HashMap
        System.out.print("Keys: ");
        for (String key : languages.keySet()) System.out.print(key+" ");  // iterate through keys only
        System.out.print("\nValues: ");
        for (Integer i : languages.values()) System.out.print(i + " ");  // iterate through values only
        System.out.print("\nEntries: ");
        for (Entry<String, Integer> entry : languages.entrySet())  System.out.print(entry + " ");  // iterate through key/value entries
    }
}
