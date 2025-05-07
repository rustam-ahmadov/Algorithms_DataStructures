public class Main {

    private static class School {
        public void Marks(){
            System.out.println('h');
        }
    }
    static class Student extends School {
//        public void Marks(){
////            System.out.println("student");
//        }
    }
    public static String reverse(StringBuilder word) {
        StringBuilder result = new StringBuilder();
        for (int i = word.length() - 1; i >= 0; i--) {
            result.append(word.charAt(i));
        }
        return result.toString();
    }

    public static void executeCommand(StringBuilder word, String command) {
        switch (command) {
            case "add":
                word.append("x");
                break;
            case "delete":
                word.deleteCharAt(word.length() - 1);
                break;
        }
    }

    public static String minDistance(int N, String word, String[] commands) {
        StringBuilder sb = new StringBuilder(word);
        for (int i = 0; i < N; i++) {
             executeCommand(sb, commands[i]);
        }
        return reverse(sb);
    }

    public static void main(String[] args) {

//        final int N = 5;
//        String word = "hello";
//        String[] commands = {"add", "delete", "add", "delete", "add"};
//
//
//
//        // "hellox" -> "xolleh"
//
//        System.out.println(minDistance(N, word, commands));


        Double d1 = new Double(134.132442);
        Double d2 = new Double(134.132442);
        d1.compareTo(d2);

        System.out.println('h');
    }
}
