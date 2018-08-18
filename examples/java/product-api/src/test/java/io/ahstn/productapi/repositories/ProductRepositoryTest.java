package io.ahstn.productapi.repositories;

import io.ahstn.productapi.ProductApiApplication;
import io.ahstn.productapi.models.Product;
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
@SpringBootTest(classes = ProductApiApplication.class, webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
public class ProductRepositoryTest {
    private static final String PRODUCT_ENDPOINT = "http://localhost:8080/api/v1/products";

    @Autowired
    private ProductRepository productRepository;

    @Before
    public void setUp() {
        if (productRepository.findByName("pixel").isEmpty()) {
            Product user = new Product("pixel", 5);
            productRepository.save(user);
        }
    }

    @Test
    public void TestGetUserStatusOkay() {
        final Response res = RestAssured.get(PRODUCT_ENDPOINT + "/2");

        assertEquals(200, res.getStatusCode());
        assertTrue(res.asString().contains("name"));
        assertTrue(res.asString().contains("pixel"));
    }

    @Test
    public void TestGetAllUsersStatusOkay() {
        final Response res = RestAssured.get(PRODUCT_ENDPOINT);

        assertEquals(200, res.getStatusCode());
        assertTrue(res.asString().contains("name"));
        assertTrue(res.asString().contains("pixel"));
    }
}