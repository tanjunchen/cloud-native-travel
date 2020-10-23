package tanjunchen.tcp.server;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * @Author tanjunchen
 * @Date 2020/10/23 19:13
 * @Version 1.0
 */
public class Server {
    public static void main(String[] args) {
        try {
            ServerSocket serverSocket = new ServerSocket(8888);
            Socket socket;
            System.out.println(" ......Starting Server...... ");
            int count = 0;
            while (true) {
                socket = serverSocket.accept();
                ServerThread serverThread = new ServerThread(socket);
                serverThread.start();
                count++;
                System.out.print("The Number of Client: " + count);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}
