package io.ahstn.usersapi.repositories;

import io.ahstn.usersapi.UsersApiApplication;
import io.ahstn.usersapi.models.User;
import io.restassured.RestAssured;
import io.restassured.response.Response;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = UsersApiApplication.class, webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
public class UserRepositoryTest {
    private static final String USER_ENDPOINT = "http://localhost:8080/api/v1/users";

    @Autowired
    private UserRepository userRepo;

    @Before
    public void setUp() {
        if (userRepo.findByEmail("adam@ahstn.io").isEmpty()) {
            User user = new User("Adam", "Houston", "adam@ahstn.io");
            userRepo.save(user);
        }
    }

    @Test
    public void TestGetUserStatusOkay() {
        final Response res = RestAssured.get(USER_ENDPOINT + "/2");

        assertEquals(200, res.getStatusCode());
        assertTrue(res.asString().contains("email"));
        assertTrue(res.asString().contains("adam@ahstn.io"));
    }

    @Test
    public void TestGetAllUsersStatusOkay() {
        final Response res = RestAssured.get(USER_ENDPOINT);

        assertEquals(200, res.getStatusCode());
        assertTrue(res.asString().contains("email"));
        assertTrue(res.asString().contains("adam@ahstn.io"));
    }
}