package com.cyh.fakeaio;

import com.cyh.bio.TimeServerHandler;

import java.net.ServerSocket;
import java.net.Socket;

public class TimeServer {
    public static void main(String[] args) throws  Exception {
        int port = 8080;

        if( args != null && args.length > 0) {
            port = Integer.valueOf(args[0]);
        }

        ServerSocket server = null;

        //create thread pool
        TimeServerHandlerExecutePool timeServerHandlerExecutePool = new TimeServerHandlerExecutePool(10, 10);
        try {

            server = new ServerSocket(port);
            System.out.println("the time server is start in port" + port);
            Socket socket = null;
            while (true) {
                socket = server.accept();

                //new Thread(new TimeServerHandler(socket)).start();
                timeServerHandlerExecutePool.execute(new TimeServerHandler(socket));
            }
        } finally {
            if (server != null) {
                System.out.println("the time server close");
                server.close();
                server = null;
            }
        }
    }
}
