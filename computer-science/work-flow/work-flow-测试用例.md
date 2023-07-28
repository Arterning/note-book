# spring boot 如何编写测试用例

在 Spring Boot 中编写测试用例是一项很重要的任务，可以确保代码的正确性和可靠性。

Spring Boot 提供了几种不同的方法来编写测试用例：

单元测试：使用 JUnit 或 TestNG 等测试框架，对单个类或方法进行测试。

集成测试：使用 Spring Test 或 TestRestTemplate 等测试工具，在真实的运行环境中进行测试。

端到端测试：使用 Spring Boot Test 项目，在完整的应用程序环境中进行测试。

以下是一个使用 JUnit 和 Mockito 编写的简单的单元测试示例：

```java
import static org.junit.Assert.assertEquals;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

@RunWith(MockitoJUnitRunner.class)
public class MyServiceTest {

  @Mock
  private MyDependency myDependency;

  @InjectMocks
  private MyService myService;

  @Test
  public void testGetData() {
    when(myDependency.getData()).thenReturn("mocked data");
    String data = myService.getData();
    assertEquals("mocked data", data);
  }
}

```

@RunWith(MockitoJUnitRunner.class)是 JUnit 的一个注解，表示使用 Mockito 的 JUnit 运行器来运行测试用例。

Mockito 是一个 Java 库，用于模拟对象以实现单元测试。MockitoJUnitRunner 类是 Mockito 的一个 JUnit 运行器，它可以在测试用例运行之前，帮助您初始化和配置所需的模拟对象。

在这个测试用例中，使用了 @Mock 注解标记 MyDependency 类，以表示该类是需要被模拟的。此外，使用了 @InjectMocks 注解标记 MyService 类，以表示该类是需要被注入模拟对象的。

如果没有使用 MockitoJUnitRunner，则需要手动使用 MockitoAnnotations.initMocks(this) 方法初始化模拟对象，而使用 MockitoJUnitRunner 后，Mockito 将自动完成这一步。

因此，添加 @RunWith(MockitoJUnitRunner.class) 注解是为了方便地使用 Mockito 进行测试，并不强制要求使用该注解。

# spring repository 测试

```java
@Entity
public class Person {
    @Id
    @GeneratedValue
    private Long id;

    private String firstName;

    private String lastName;

    // Getters and setters omitted for brevity
}

public interface PersonRepository extends JpaRepository<Person, Long> {
    Optional<Person> findByLastName(String lastName);
}


// 在这个测试中，你可以使用 TestEntityManager 插入一个 Person 实体并且查询它，以验证 PersonRepository 的数据访问功能是否正确。
@DataJpaTest
public class PersonRepositoryTest {

    @Autowired
    private TestEntityManager entityManager;

    @Autowired
    private PersonRepository personRepository;

    @Test
    public void findByLastName_shouldReturnPerson() {
        // given
        Person peterPan = new Person("Peter", "Pan");
        entityManager.persist(peterPan);
        entityManager.flush();

        // when
        Optional<Person> result = personRepository.findByLastName("Pan");

        // then
        assertThat(result).isNotEmpty().get().isEqualToComparingFieldByField(peterPan);
    }
}

```
