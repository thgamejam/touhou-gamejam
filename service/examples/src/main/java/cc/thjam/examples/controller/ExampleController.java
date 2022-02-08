package cc.thjam.examples.controller;

import cc.thjam.examples.api.Example.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value="/api/v1/example", produces = "application/json;charset=UTF-8")
public class ExampleController {

    private static final Logger logger = LoggerFactory.getLogger(ExampleController.class);

    private static Integer COUNT = 0;

    @PostMapping(value="")
    public ExampleReply addUser(@RequestBody ExampleRequest request) {
        logger.info("样例请求：{}" , request.toString());

        ExampleReply.Builder reply = ExampleReply.newBuilder();
        reply.setUsername(request.getUsername());
        reply.setCount(++COUNT);
        return reply.build();
    }

}
