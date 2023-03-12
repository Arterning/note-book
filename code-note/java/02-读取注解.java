public class Test {

    public static void main(String[] args) {
        List<String> allUrls = new ArrayList<>();
            RequestMapping classAnnotation = method.getDeclaringClass().getAnnotation(RequestMapping.class);
            ArrayList<String> classUrls = Optional.ofNullable(classAnnotation)
                    .map(RequestMapping::value)
                    .map(array -> new ArrayList<>(Arrays.asList(array)))
                    .orElse(new ArrayList<>());
    
            GetMapping getAnnotation = method.getAnnotation(GetMapping.class);
            PostMapping postAnnotation = method.getAnnotation(PostMapping.class);
            PutMapping putAnnotation = method.getAnnotation(PutMapping.class);
            PatchMapping patchAnnotation = method.getAnnotation(PatchMapping.class);
            RequestMapping requestAnnotation = method.getAnnotation(RequestMapping.class);
            if (getAnnotation != null) {
                ArrayList<String> subUrls = new ArrayList<>(Arrays.asList(getAnnotation.value()));
                subUrls.forEach(subUrl -> {
                    if (classUrls.isEmpty()) {
                                 allUrls.add(subUrl);
                    } else {
                        classUrls.forEach(classUrl -> {
                            allUrls.add(classUrl + subUrl);
                        });
                    }
                });
    
        }
    
    
        
    }
    

    private String getFullMethodName(Method method) {
        String className = method.getDeclaringClass().getName();
        return className + "." + method.getName();
    }
}




