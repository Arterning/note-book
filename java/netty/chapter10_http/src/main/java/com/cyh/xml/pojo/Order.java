package com.cyh.xml.pojo;

import lombok.Data;

/**
 * Order information.
 */
@Data
public class Order {

    private long orderNumber;

    private Customer customer;

    /** Billing address information. */
    private Address billTo;

    private Shipping shipping;

    /**
     * Shipping address information. If missing, the billing address is also
     * used as the shipping address.
     */
    private Address shipTo;

    private Float total;


}
