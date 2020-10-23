package com.tanjunchen.tcp.client;

import java.io.*;
import java.net.Socket;
import java.util.Random;

/**
 * @Author tanjunchen
 * @Date 2020/10/23 19:15
 * @Version 1.0
 */
public class Client {

    public static String numbers() {
        String[] str = {"1", "2", "3"};
        int length = str.length;
        Random random = new Random();
        int number = random.nextInt(length);
        return str[number];
    }

    public static void main(String[] args) {
        try {
            Socket socket = new Socket("localhost", 8888);
            OutputStream os = socket.getOutputStream();
            PrintWriter pw = new PrintWriter(os);
            pw.write(numbers());
            pw.flush();
            socket.shutdownOutput();

            InputStream is = socket.getInputStream();
            InputStreamReader ist = new InputStreamReader(is);
            BufferedReader br = new BufferedReader(ist);

            String data;
            while ((data = br.readLine()) != null) {
                System.out.println("com.tanjunchen.tcp.client.Client ===> " + data);
            }

            is.close();
            ist.close();
            br.close();
            br.close();
            pw.close();
            os.close();
            socket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
