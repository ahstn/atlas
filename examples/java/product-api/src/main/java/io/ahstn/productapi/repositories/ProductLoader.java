package io.ahstn.productapi.repositories;


import io.ahstn.productapi.models.Product;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
public class ProductLoader implements CommandLineRunner {

  private final ProductRepository repository;

  @Autowired
  public ProductLoader(ProductRepository repository) {
    this.repository = repository;
  }

  @Override
  public void run(String... strings) throws Exception {
    this.repository.save(new Product("Pixel XL", 2));
  }
}
