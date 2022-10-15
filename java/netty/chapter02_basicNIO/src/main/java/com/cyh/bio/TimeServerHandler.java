package com.cyh.bio;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

public class TimeServerHandler implements Runnable{

    private Socket socket;

    public TimeServerHandler(Socket socket) {
        this.socket = socket;
    }

    @Override
    public void run() {

        BufferedReader in = null;
        PrintWriter out = null;

        try {
            in = new BufferedReader(new InputStreamReader(this.socket.getInputStream()));
            out = new PrintWriter(this.socket.getOutputStream(), true);
            String body = null;
            while (true) {

                body = in.readLine();
                if (body == null) {
                    break;
                }
                System.out.println("the time server receive order:" + body);
            }
        } catch (Exception e) {
            try {
                in.close();
                out.close();
            } catch (IOException ex) {
                ex.printStackTrace();
            }

        }

    }
}
