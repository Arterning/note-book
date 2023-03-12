/**
This is a situation when the relationship itself has an attribute.
Let's say we want to let students rate the courses.
 */

 /**
 
 01的实现非常简单直接 但是缺点是不能给relation添加自己的属性 
 比如学生给课程评分 这个评分需要写在relation table的字段中
 The implementation of a simple many-to-many relationship was rather straightforward. The only problem is that we cannot add a property to a relationship that way because we connected the entities directly. Therefore, we had no way to add a property to the relationship itself.

 */
@Embeddable
class CourseRatingKey implements Serializable {

    @Column(name = "student_id")
    Long studentId;

    @Column(name = "course_id")
    Long courseId;

    // standard constructors, getters, and setters
    // hashcode and equals implementation
}


@Entity
class CourseRating {

    @EmbeddedId
    CourseRatingKey id;

    @ManyToOne
    @MapsId("studentId")
    @JoinColumn(name = "student_id")
    Student student;

    @ManyToOne
    @MapsId("courseId")
    @JoinColumn(name = "course_id")
    Course course;

    /**
        给关系新增rating属性字段
     */
    int rating;
    
    // standard constructors, getters, and setters
}


class Student {

    // ...

    @OneToMany(mappedBy = "student")
    Set<CourseRating> ratings;

    // ...
}

class Course {

    // ...

    @OneToMany(mappedBy = "course")
    Set<CourseRating> ratings;

    // ...
}