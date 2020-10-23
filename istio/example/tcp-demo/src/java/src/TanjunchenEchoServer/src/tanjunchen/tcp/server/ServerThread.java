package tanjunchen.tcp.server;

import java.io.*;
import java.net.Socket;
import java.util.Random;

/**
 * @Author tanjunchen
 * @Date 2020/10/23 16:32
 * @Version 1.0
 */
public class ServerThread extends Thread {

    Socket socket;

    public ServerThread(Socket socket) {
        this.socket = socket;
    }

    public String response() {
        String[] str = {"Server-A", "Server-B", "Server-C"};
        int length = str.length;
        Random random = new Random();
        int number = random.nextInt(length);
        return str[number];
    }

    @Override
    public void run() {
        InputStream is = null;
        InputStreamReader isr = null;
        BufferedReader br = null;
        OutputStream os = null;
        PrintWriter pw = null;
        try {
            is = socket.getInputStream();
            isr = new InputStreamReader(is);
            br = new BufferedReader(isr);

            String info;
            while ((info = br.readLine()) != null) {
                System.out.println(" 服务端收到: " + info);
            }
            socket.shutdownInput();

            os = socket.getOutputStream();
            pw = new PrintWriter(os);
            pw.write(" 服务端响应: " + response());
            pw.flush();
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                if (os != null) {
                    os.close();
                }
                if (br != null) {
                    br.close();
                }
                if (isr != null) {
                    isr.close();
                }
                if (is != null) {
                    is.close();
                }
                if (pw != null) {
                    pw.close();
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}
