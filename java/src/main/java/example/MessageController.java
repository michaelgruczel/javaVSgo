package example;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

import static org.springframework.http.MediaType.APPLICATION_JSON_VALUE;

@RestController
public class MessageController {

    @Autowired
    MessageRepository messageRepository;

    @RequestMapping(method = RequestMethod.GET, value = "/messages", produces = APPLICATION_JSON_VALUE)
    public List<Message> messages() {
        return returnAllMessages();
    }
    
    @RequestMapping(method = RequestMethod.POST, value = "/message", produces = APPLICATION_JSON_VALUE)
    public List<Message> message(@RequestBody Message message) {
        messageRepository.save(message);
        return returnAllMessages();
    }

    private List<Message> returnAllMessages() {
        List<Message> result = new ArrayList<Message>();
        messageRepository.findAll().iterator().forEachRemaining(result::add);
        return result;
    }
}
