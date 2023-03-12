package com.bezkoder.spring.hibernate.onetomany.model;

import javax.persistence.*;

import org.hibernate.annotations.OnDelete;
import org.hibernate.annotations.OnDeleteAction;

import com.fasterxml.jackson.annotation.JsonIgnore;

/**
– @Entity annotation indicates that the class is a persistent Java class.
– @Table annotation provides the table that maps this entity.

https://www.bezkoder.com/jpa-one-to-many/

We also use the @JoinColumn annotation to specify the foreign key column (tutorial_id). If you don’t provide the JoinColumn name, the name will be set automatically.

 */
@Entity
@Table(name = "comments")
public class Comment {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "comment_generator")
  private Long id;

  @Lob
  private String content;

  /**
  many端声明外键 使用JoinColumn
   */
  @ManyToOne(fetch = FetchType.LAZY, optional = false)
  @JoinColumn(name = "tutorial_id", nullable = false)
  @OnDelete(action = OnDeleteAction.CASCADE)
  @JsonIgnore
  private Tutorial tutorial;

  // getters and setters
}