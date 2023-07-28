# java stream group by

```java
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Student {
    private String name;
    private String gender;
    private String studentId;
    private String className;

    // 构造函数和getter/setter方法省略

    public static void main(String[] args) {
        List<Student> students = Arrays.asList(
                new Student("Alice", "Female", "001", "Class A"),
                new Student("Bob", "Male", "002", "Class B"),
                new Student("Cindy", "Female", "003", "Class A"),
                new Student("David", "Male", "004", "Class B"),
                new Student("Emily", "Female", "005", "Class A")
        );

        // 使用Stream按班级分组
        Map<String, List<Student>> groupedStudents = students.stream()
                .collect(Collectors.groupingBy(Student::getClassName));

        // 打印分组结果
        for (Map.Entry<String, List<Student>> entry : groupedStudents.entrySet()) {
            String className = entry.getKey();
            List<Student> studentsInClass = entry.getValue();

            System.out.println("班级：" + className);
            System.out.println("学生列表：");
            for (Student student : studentsInClass) {
                System.out.println(student.getName() + " - " + student.getStudentId());
            }
            System.out.println();
        }
    }
}
```