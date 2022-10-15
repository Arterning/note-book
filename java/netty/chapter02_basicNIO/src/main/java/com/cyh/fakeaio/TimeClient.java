package com.cyh.fakeaio;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

public class TimeClient {


    public static void main(String[] args) {
        Socket socket;
        BufferedReader in = null;
        PrintWriter out = null;

        int port = 8080;
        try {

            socket = new Socket("localhost", port);
            in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
            out = new PrintWriter(socket.getOutputStream(), true);

            out.println("this is client, fuck you~~~~~");

        } catch (Exception e){

            try {
                if (in != null) {
                    in.close();
                }
                if (out != null) {
                    out.close();
                }
            } catch (IOException ex) {
                ex.printStackTrace();
            }
        }

    }
}
