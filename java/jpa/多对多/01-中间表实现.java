/**
using @JoinTable or even @JoinColumn isn't required. JPA will generate the table and column names for us. 

However, the strategy JPA uses won't always match the naming conventions we use. 

So, we need the possibility to configure table and column names.

 */
@Entity
class Student {

    @Id
    Long id;

    @ManyToMany
    @JoinTable(
    name = "course_like", 
    joinColumns = @JoinColumn(name = "student_id"), 
    inverseJoinColumns = @JoinColumn(name = "course_id"))
    Set<Course> likedCourses;

    // additional properties
    // standard constructors, getters, and setters
}

@Entity
class Course {

    @Id
    Long id;

    @ManyToMany(mappedBy = "likedCourses")
    Set<Student> likes;

    // additional properties
    // standard constructors, getters, and setters
}