@Test
public void corePermissionCheck() {
    ArrayList<Object> userPermissions = Lists.newArrayList(23423423, 111111111188L, 23423424);

    ArrayList<Object> urlPermissions = Lists.newArrayList(23423, 2322, 45756, 111111111188L);

    boolean disjoint = Collections.disjoint(userPermissions, urlPermissions);

    System.out.println(disjoint);

}