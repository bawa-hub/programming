import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Problems {
    public static void main(String[] args) {

        // convert list to array and vice versa
        ArrayList languages = new ArrayList<>();

        // Add elements in the list
        languages.add("Java");
        languages.add("Python");
        languages.add("JavaScript");
        System.out.println("ArrayList: " + languages);

        // Create a new array of String type
        String[] arr = new String[languages.size()];

        // Convert ArrayList into the string array
        languages.toArray(arr);
        System.out.print("Array: ");
        for (String item : arr) {
            System.out.println(item + ", ");
        }

        // convert java array into list
        String[] array = { "Java", "Python", "C" };
        System.out.println("Array: " + Arrays.toString(array));

        List lang = new ArrayList<>(Arrays.asList(array));
        System.out.println("List: " + lang);

    }
}
