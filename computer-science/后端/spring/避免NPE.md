在Java 8中，你可以通过使用流（Stream）和 `forEach` 方法来避免在集合为空时导致空指针错误。流提供了一种优雅的方式来处理集合，可以确保在集合为空时不会引发异常。

例如，如果你要遍历一个集合并执行某些操作，可以使用流的 `forEach` 方法。以下是一个示例：

```java
List<String> myList = // 从某个地方获取集合数据

myList.stream()
      .forEach(item -> {
          // 执行操作
          System.out.println(item);
      });

```

使用流的 `forEach` 方法可以安全地处理空集合，不会导致空指针异常。如果集合为空，它将不执行任何操作，而不会引发异常。

如果你要在集合为空时采取特定的操作，你可以使用 `ifPresent` 方法，如下所示：

```java
List<String> myList = // 从某个地方获取集合数据

myList.stream()
      .findFirst()
      .ifPresent(item -> {
          // 集合非空时执行操作
          System.out.println(item);
      });

```

在上述示例中，我们使用 `findFirst` 方法获取集合中的第一个元素，然后使用 `ifPresent` 方法检查是否存在元素，如果存在，则执行操作。

这些方法可以帮助你更加优雅地处理集合为空的情况，而不必担心空指针异常。

你说得对，我之前的回答有误。如果 `myList` 为空，调用 `stream()` 方法确实会引发 `NullPointerException`。我为之前的错误回答表示歉意。

要避免在集合为空时引发异常，你可以在调用 `stream()` 方法之前进行空检查，例如：

```java
List<String> myList = // 从某个地方获取集合数据

if (myList != null) {
    myList.stream()
          .forEach(item -> {
              // 执行操作
              System.out.println(item);
          });
}

```

或者，如果你使用 Java 8+，你可以使用 `Optional` 来处理空集合，如下所示：

```java
List<String> myList = // 从某个地方获取集合数据

Optional.ofNullable(myList)
        .orElse(Collections.emptyList())
        .stream()
        .forEach(item -> {
            // 执行操作
            System.out.println(item);
        });

```

在上述示例中，我们首先使用 `Optional.ofNullable(myList)` 将 `myList` 转换为 `Optional`，然后使用 `orElse(Collections.emptyList())` 来提供一个默认的空集合，这样就可以确保在 `myList` 为空时不会引发异常。

这样可以安全地处理空集合的情况，而不必担心空指针异常。

