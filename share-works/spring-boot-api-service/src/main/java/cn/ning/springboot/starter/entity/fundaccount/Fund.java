package cn.ning.springboot.starter.entity.fundaccount;


import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "fund")
@Setter
@Getter
public class Fund {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;



    @OneToMany(mappedBy = "fund")
    private Set<Account> accounts = new HashSet<>();
}
