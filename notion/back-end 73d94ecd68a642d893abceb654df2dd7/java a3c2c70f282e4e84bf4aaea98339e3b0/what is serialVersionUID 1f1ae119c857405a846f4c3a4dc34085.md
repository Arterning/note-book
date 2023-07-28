# what is serialVersionUID

**Simply put, we use the *serialVersionUID* attribute to remember versions of a *Serializable* class to verify that a loaded class and the serialized object are compatible.**

If *serialVersionUID* is not provided in a *Serializable* class, the JVM will generate one automatically. However, **it is good practice to provide the *serialVersionUID* value and update it after changes to the class so that we can have control over the serialization/deserialization process**. We'll take a closer look at it in a later section.