import java.util.concurrent.TimeUnit;

class SingleThread {
    public static void main(String[] args) {
        while (true) {
            System.out.println("Hello, World!");

            try {
                TimeUnit.SECONDS.sleep(1);
            } catch(InterruptedException ex) {
                Thread.currentThread().interrupt();
            }
        } 
    }
}
