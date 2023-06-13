# JPA 一对多 和  多对多

Created time: May 12, 2023 10:57 AM

```jsx
@ManyToOne(fetch = FetchType.LAZY, optional = false)
@JoinColumn(name = "tutorial_id", nullable = false)
private Tutorial tutorial;

@ManyToMany(fetch = FetchType.LAZY,
      cascade = {
          CascadeType.PERSIST,
          CascadeType.MERGE
      })
 @JoinTable(name = "tutorial_tags",
        joinColumns = { @JoinColumn(name = "tutorial_id") },
        inverseJoinColumns = { @JoinColumn(name = "tag_id") })
private Set<Tag> tags = new HashSet<>();
```

```jsx
@OneToMany(mappedBy = "tutorial")
private Set<TutorialComment> comments = new HashSet<>();
```

```jsx
@ManyToMany(fetch = FetchType.LAZY,
      cascade = {
          CascadeType.PERSIST,
          CascadeType.MERGE
      },
      mappedBy = "tags")
  @JsonIgnore
  private Set<Tutorial> tutorials = new HashSet<>();
```