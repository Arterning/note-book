@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface Authorize {

    AuthorizeType value() default AuthorizeType.PRIVILEGE;

}


public enum AuthorizeType {

    IGNORE,

    LOGIN,

    PRIVILEGE
}

