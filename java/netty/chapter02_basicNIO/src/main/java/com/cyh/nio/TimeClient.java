package com.cyh.nio;

/**
 * @author: CYH
 * @date: 2018/9/21 11:42
 */
public class TimeClient {

    public static void main(String[] args) {
        TimeClientHandle clientHandle = new TimeClientHandle(null, 15689);
        new Thread(clientHandle, "NIO-TimeClientHandler-001").start();
    }

}

