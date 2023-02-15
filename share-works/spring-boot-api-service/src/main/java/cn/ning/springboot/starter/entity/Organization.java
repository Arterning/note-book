package cn.ning.springboot.starter.entity;


import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.*;

@Entity
@Table(name = "organization")
@Setter
@Getter
public class Organization {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;



    @OneToMany(mappedBy = "organization")
    private Set<Person> persons = new HashSet<>();
}
