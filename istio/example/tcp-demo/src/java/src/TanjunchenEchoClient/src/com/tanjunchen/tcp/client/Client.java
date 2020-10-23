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

    private static String host;
    private static int port;
    private static int value;
    private static volatile int stop;

    public static String numbers() {
        String[] str = {"Client-1", "Client-2", "Client-3"};
        int length = str.length;
        Random random = new Random();
        int number = random.nextInt(length);
        return str[number];
    }

    public static void start() {
        try {
            Socket socket = new Socket(host, port);
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
                System.out.println("客户端：" + data);
            }

            is.close();
            ist.close();
            br.close();
            br.close();
            pw.close();
            os.close();
            socket.close();
        } catch (IOException e) {
            stop = 1;
            e.printStackTrace();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        // 传递的值为 1 则一直调用(休眠 1 秒)
        value = 0;
        if (args.length > 0) {
            value = Integer.parseInt(args[0]);
        }

        host = "localhost";
        if (args.length > 1) {
            host = args[1];
        }

        // 传递的值为 1,
        port = 8888;
        if (args.length > 2) {
            port = Integer.parseInt(args[2]);
        }

        System.out.println(host + " == " + port + " == " + value);
        if (value == 1) {
            while (stop == 0) {
                start();
                Thread.sleep(1000);
            }
        } else {
            start();
        }
    }
}
