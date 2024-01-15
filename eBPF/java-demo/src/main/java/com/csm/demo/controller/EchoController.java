package com.csm.demo.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class EchoController {

    private static final Logger logger = LoggerFactory.getLogger(EchoController.class);

	private static int counter = 100000;

	@RequestMapping(value = "/", method = RequestMethod.GET)
    public String root() {
		String hello = "helloworld";
        logger.info(hello);
		return hello;
    }

    @RequestMapping(value = "/hello", method = RequestMethod.GET)
    public String hello() {
		String hello = "helloworld";
        logger.info(hello);
		return hello;
    }

	@RequestMapping(value = "/ebpf/function/{message}", method = RequestMethod.GET)
    public String function(@PathVariable String message) {
		if (message.equals("count")) {
			functionCount();
			return "count";
		}else if (message.equals("exception")) {
			try {
				logger.info("we will trigger exception in function");
				int i = 1/0;
			} catch (Exception e) {
				String exception = e.getMessage();
				logger.error(exception);
			}
			return "exception";
		}else if (message.equals("latency")) {
			try {
				System.out.println("开始休眠");
				Thread.sleep(5000); // 休眠5秒
				System.out.println("休眠结束");
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			return "latency";
		}
		return "";
    }

	public void functionCount() {
		for (int i = 0; i < counter; ++i) {
			System.out.println(i);
		}
		return;
	}
}
