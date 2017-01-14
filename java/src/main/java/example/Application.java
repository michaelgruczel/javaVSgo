package example;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@SpringBootApplication
@EnableSwagger2
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Bean
    public Docket newsApi() {
        return new Docket(DocumentationType.SWAGGER_2)
                .groupName("example")
                .select()
                .paths(PathSelectors.any())
                .build();
    }

    @Bean
    public CommandLineRunner init(MessageRepository repository) {
        return (args) -> {
            repository.save(new Message("Homer", "I don't believe in communication"));
            repository.save(new Message("Lisa", "Give it a chance"));
        };
    }
}
