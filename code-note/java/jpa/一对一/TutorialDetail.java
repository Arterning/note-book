package com.bezkoder.spring.jpa.onetoone.model;

import java.util.Date;

import javax.persistence.*;
/**
https://www.bezkoder.com/jpa-one-to-one/

the Tutorial entity and @MapsId annotation that makes the id field serve as both Primary Key and Foreign Key (shared primary key).
MapsId 让这个字段既作为主键 也作为外键
 */
@Entity
@Table(name = "tutorial_details")
public class TutorialDetails {
  @Id
  private Long id;

  @Column
  private Date createdOn;

  @Column
  private String createdBy;

  /**
  关联关系和级联关系我们统一在一边定义
  另外一边只需要使用mappedBy就可以了
  @OneToOne(mappedBy = "tutorial")
  private TutorialDetail detail;
   */
  @OneToOne(fetch = FetchType.LAZY)
  @MapsId
  @JoinColumn(name = "tutorial_id")
  private Tutorial tutorial;


  public TutorialDetails() {
  }

  public TutorialDetails(String createdBy) {
    this.createdOn = new Date();
    this.createdBy = createdBy;
  }

  // getters and setters
}