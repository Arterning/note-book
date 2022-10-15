package com.cyh.nio;

/**
 * @author: CYH
 * @date: 2018/9/21 11:18
 */
public class TimeServer {

    public static void main(String[] args) {
        MultiplexerTimeServer timeServer = new MultiplexerTimeServer(15689);
        new Thread(timeServer, "NIO-MultiplexerTimeServer-001").start();
    }

}

