package com.cyh.xml.pojo;

import java.util.List;

import lombok.Data;

/**
 * Customer information.
 */
@Data
public class Customer {

    private long customerNumber;

    /** Personal name. */
    private String firstName;

    /** Family name. */
    private String lastName;

    /** Middle name(s), if any. */
    private List<String> middleNames;


}
