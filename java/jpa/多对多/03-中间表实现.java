/**
Many-to-Many With a New Entity
 */

/**

Let's say we want to let students register for courses. Also, we need to store the point when a student registered for a specific course. On top of that, we want to store what grade she received in the course.

In an ideal world, we could solve this with the previous solution, where we had an entity with a composite key. However, the world is far from ideal, and students don't always accomplish a course on the first try.

In this case, there are multiple connections between the same student-course pairs, or multiple rows, with the same student_id-course_id pairs. We can't model it using any of the previous solutions because all primary keys must be unique. So, we need to use a separate primary key.

Therefore, we can introduce an entity, which will hold the attributes of the registration

新的需求来了 我们需要给学生在这门课程中获得的分数记录下来 
现实中的问题是，学生总是可能发生补考的，学生-课程这个关系可能存在多条记录 也就是补考很多次 每次考试都是不同的分数
所以这个relation是存在多份的 每份的考试成绩是不同的 
而之前使用的联合主键方式无法支持多个 相同的student-cores的relation
所以我们需要引入这种方式
 */
@Entity
class CourseRegistration {

    @Id
    Long id;

    @ManyToOne
    @JoinColumn(name = "student_id")
    Student student;

    @ManyToOne
    @JoinColumn(name = "course_id")
    Course course;

    LocalDateTime registeredAt;

    int grade;
    
    // additional properties
    // standard constructors, getters, and setters
}


class Student {

    // ...

    @OneToMany(mappedBy = "student")
    Set<CourseRegistration> ratings;

    // ...
}

class Course {

    // ...

    @OneToMany(mappedBy = "course")
    Set<CourseRegistration> ratings;

    // ...
}