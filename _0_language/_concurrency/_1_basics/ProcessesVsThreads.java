public class ProcessesVsThreads {

    public static void main(String[] args) {
        // Example of Thread
        Runnable task = () -> {
            System.out.println("This is running in a thread.");
        };
        
        Thread thread = new Thread(task);
        thread.start();  // Start a new thread
        
        // Example of Process (Using ProcessBuilder to launch a new process)
        try {
            ProcessBuilder processBuilder = new ProcessBuilder("java", "-version");
            Process process = processBuilder.start();
            process.waitFor();  // Wait for the process to complete
            System.out.println("Process has finished.");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
